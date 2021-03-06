FROM golang:1.16.0-alpine3.11 as builder

RUN apk update && apk upgrade && \
    apk --update add git gcc make

WORKDIR /app

COPY . .

RUN make go-build

# Distribution
#FROM alpine:latest
#
#RUN apk update && apk upgrade && \
#    apk --update --no-cache add tzdata && \
#    mkdir /app
#
#WORKDIR /app
#
#EXPOSE 9090
#
#COPY --from=builder /app/goscription /app
#
#CMD /app/goscription http