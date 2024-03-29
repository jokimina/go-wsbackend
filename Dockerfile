FROM golang:alpine AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
   && apk update && apk add --no-cache git

WORKDIR /go/src/wsbackend
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags="-w -s" -o /wsbackend /go/src/wsbackend/cmd

FROM alpine:3.7
RUN apk add tzdata ca-certificates && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata && rm -rf /var/cache/apk/*
COPY --from=builder /wsbackend /app/wsbackend
ENTRYPOINT ["/app/wsbackend"]
