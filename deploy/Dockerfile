FROM golang:1.16.1-alpine3.13 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o go_starter app/main.go

FROM alpine
RUN apk update && apk --no-cache add tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone && \
    mkdir -p /app/config
WORKDIR /app
COPY --from=builder /app/go_starter /app
COPY --from=builder /app/config/config.yml /app/config
EXPOSE 8080
CMD ["/app/go_starter"]