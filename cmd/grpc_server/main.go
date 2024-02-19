package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/sarastee/auth/internal/config"
	envcfg "github.com/sarastee/auth/internal/config/env"
	desc "github.com/sarastee/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "./config/prod.env", "path to config file")
}

type server struct {
	desc.UnimplementedUserV1Server
	pool *pgxpool.Pool
	sq   squirrel.StatementBuilderType
}

// User variable struct
type User struct {
	ID        int64
	Name      string
	Email     string
	Role      string
	CreateAt  time.Time
	UpdatedAt time.Time
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Printf(os.Getwd())
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := envcfg.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := envcfg.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	//if err != nil {
	//	log.Fatalf("failed to connect to database: %v", err)
	//}
	//defer pool.Close()

	pool, err := NewClient(ctx, 5, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{
		pool: pool,
	})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("user create request")

	builderInsert := s.sq.Insert("users").
		PlaceholderFormat(squirrel.Dollar).
		Columns("name", "email", "password", "password_confirm", "role").
		Values(req.GetName(), req.GetEmail(), req.GetPassword(), req.GetPasswordConfirm(), req.GetRole()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var userID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}

	log.Printf("inserted user with id: %d", userID)

	return &desc.CreateResponse{
		Id: userID,
	}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("user get request")

	builderSelect := s.sq.Select("id", "name", "email", "role", "created_at", "updated_at").
		From("users").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": req.GetId()}).
		Limit(1)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Fatalf("failed to build query to get user: %v", err)
	}

	q, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to execute query to get user: %v", err)
	}

	var user User
	for q.Next() {
		err = q.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreateAt, &user.UpdatedAt)
		if err != nil {
			log.Fatalf("failed to read user data from database: %v", err)
		}
	}
	userRole := desc.Role_value[user.Role]

	userData := desc.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(userRole),
		CreatedAt: timestamppb.New(user.CreateAt),
		UpdateAt:  timestamppb.New(user.UpdatedAt),
	}

	log.Printf("user info: %v", user)

	return &userData, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("user update request")

	builderUpdate := s.sq.Update("users").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": req.GetId()})

	if req.GetName() != "" {
		builderUpdate = builderUpdate.Set("name", req.GetName())
	}
	if req.GetEmail() != "" {
		builderUpdate = builderUpdate.Set("email", req.GetEmail())
	}
	if req.GetRole() != desc.Role_UNKNOWN {
		builderUpdate = builderUpdate.Set("role", req.GetRole().String())
	}

	builderUpdate = builderUpdate.Set("updated_at", timestamppb.Now().AsTime().Format("2006-01-02 15:04:05"))

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Fatalf("failed build query to update user: %v", err)

	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to execute query to update user: %v", err)
	}

	log.Printf("user was updated")

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("user delete request")

	builderDelete := s.sq.Delete("users").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": req.GetId()})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed build query to delete user: %v", err)
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to execute query to delete user: %v", err)
	}

	log.Printf("user was deleted")
	return &emptypb.Empty{}, nil
}

func NewClient(ctx context.Context, maxAttempts int, dsn string) (*pgxpool.Pool, error) {
	for ; maxAttempts != 0; maxAttempts-- {
		select {
		case <-ctx.Done():
		default:
			if pool, err := connect(ctx, dsn); err != nil {
				log.Printf("remaining attempts %d", maxAttempts)
				continue
			} else {
				log.Printf("Connection to database is successful")
				return pool, nil
			}
		}
	}

	return nil, errors.New("connection has not been established")
}

func connect(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeoutCause(ctx, 3*time.Second, errors.New("timeout 3s exceeded"))
	defer cancel()

	if pool, err := pgxpool.Connect(ctx, dsn); err != nil {
		return nil, errors.Wrap(err, "pgxpool.Connect")
	} else if err := pool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "unable to ping to database connect")
	} else {
		return pool, nil
	}
}

//type PgConfig struct {
//	DBName     string
//	DBUser     string
//	DBPassword string
//	DBHost     string
//}

//func (c PgConfig) DSN() string {
//	return fmt.Sprintf(
//		"postgres://%s:%s@%s/%s", // ?pool_max_conns=100
//		//"postgres://username:password@localhost:5432/database_name",
//		c.DBUser, url.QueryEscape(c.DBPassword), c.DBHost, c.DBName,
//	)
//}
