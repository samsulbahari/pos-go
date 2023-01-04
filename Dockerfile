# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:alpine AS build

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/clean-arsitecture main.go

##
## Deploy
##
FROM alpine

WORKDIR /app


COPY --from=build /app/clean-arsitecture /app/clean-arsitecture


ENTRYPOINT /app/clean-arsitecture