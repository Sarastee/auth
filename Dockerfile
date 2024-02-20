FROM golang:1.21.7-alpine AS builder

COPY . /github.com/sarastee/auth/source/
WORKDIR /github.com/sarastee/auth/source/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/sarastee/auth/bin/auth_server .

CMD ["./auth_server"]
