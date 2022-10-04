# syntax=docker/dockerfile:1

##
## Build stage
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /chat-server /app/cmd/server/main.go

##
## Deploy stage
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /chat-server /chat-server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/chat-server"]
