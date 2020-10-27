FROM golang:1.15-alpine as builder

WORKDIR /src/

COPY . .

RUN go build main.go \
    && chmod +x main

RUN apk add bash