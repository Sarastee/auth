FROM golang:1.21.7-alpine AS builder

COPY . /github.com/sarastee/auth/source/
WORKDIR /github.com/sarastee/auth/source/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/sarastee/auth/source/bin/auth_server .
COPY --from=builder /github.com/sarastee/auth/source/config/prod.env .

CMD ["systemctl start postgresql.service",  "./auth_server"]
