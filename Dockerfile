FROM golang:1.20.1 AS builder
WORKDIR /go/src/uow
COPY ./go.mod .
COPY ./go.sum .
COPY ./vendor ./vendor

