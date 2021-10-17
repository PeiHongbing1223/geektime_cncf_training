FROM golang:1.17.2 AS build

COPY httpserver.go /go/src/httpserver/

WORKDIR /go/src/httpserver

ENV GOPROXY=https://goproxy.cn

RUN go mod init \
 && go mod tidy \
 && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o httpserver

FROM alpine:3.12

MAINTAINER NDQ

COPY --from=build /go/src/httpserver/httpserver /bin/httpserver

ENV VERSION=httpserver:v2.0

EXPOSE 80

ENTRYPOINT ["/bin/httpserver"]
