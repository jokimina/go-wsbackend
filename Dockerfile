FROM golang:alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add --no-cache git

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

WORKDIR /go/src/wsbackend
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /wsbackend /go/src/wsbackend/cmd

FROM scratch
COPY --from=builder /go/src/wsbackend/data /data
COPY --from=builder /wsbackend /app/wsbackend
ENTRYPOINT ["/app/wsbackend"]
