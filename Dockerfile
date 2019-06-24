FROM golang:alpine AS builder
ADD ./ /
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /wsbackend /cmd

FROM scratch
COPY --from=builder /wsbackend /wsbackend
ENTRYPOINT ["/wsbackend"]