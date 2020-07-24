######################## Build stage 1 ########################
FROM golang:1.14.1-alpine3.11

WORKDIR /go/src/dummyhttps
ADD main.go main.go
ADD ./.ssl-certs /root/.ssl-certs

# RUN apk --no-cache add ca-certificates gcc musl-dev git bash openssl curl curl-dev tzdata unzip wget
RUN apk add ca-certificates openssl curl

RUN CGO_ENABLED=0 GOOS=linux go build \
  -o /root/runtime.x64 \
  -a \
  -tags netgo \
  -ldflags "-d -s -w" \
  -installsuffix cgo \
  ./main.go

RUN chmod +x /root/runtime.x64

ENTRYPOINT ["/root/runtime.x64"]