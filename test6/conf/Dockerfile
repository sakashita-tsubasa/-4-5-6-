FROM golang:1.20.0-alpine as dev

WORKDIR /go/src

RUN apk update && apk add --no-cache git zip
RUN apk update \
    && apk add --no-cache git \
    && apk add --no-cache gcc musl-dev
