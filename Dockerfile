FROM golang:alpine

MAINTAINER Maintainer


WORKDIR /app

ADD . / /app/

RUN apk update
RUN apk add --no-cache git
RUN apk add --no-cache sqlite-libs sqlite-dev
RUN apk add --no-cache build-base
RUN go get ./...

RUN go build

EXPOSE 80
EXPOSE 443

ENTRYPOINT ["./shop"]
