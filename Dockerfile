FROM golang:1.21.7-alpine AS builder

COPY . /
WORKDIR /

RUN go mod download
RUN go build -o ./bin/auth_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /bin/auth_server .

CMD ["./auth_server"]
