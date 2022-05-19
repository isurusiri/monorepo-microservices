FROM golang:1.16.5 AS builder
ARG SERVICE_NAME
WORKDIR $GOPATH/src/github.com/isurusiri/monorepo-microservices
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o app main/$SERVICE_NAME/main.go

FROM alpine:latest
WORKDIR /root/
RUN mkdir -p ./main/app
COPY --from=builder /go/src/github.com/isurusiri/monorepo-microservices/app .
COPY --from=builder /go/src/github.com/isurusiri/monorepo-microservices/config/config.yaml ./config/
CMD ["./app"]

EXPOSE 8080