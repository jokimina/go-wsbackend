FROM golang:alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add --no-cache git tzdata

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /go/src/wsbackend
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /wsbackend /go/src/wsbackend/cmd

FROM scratch
# curl -o cacert.pem https://curl.haxx.se/ca/cacert.pem
ADD data/cacert.pem /etc/ssl/certs/

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/src/wsbackend/data /data
COPY --from=builder /wsbackend /app/wsbackend
ENTRYPOINT ["/app/wsbackend"]
