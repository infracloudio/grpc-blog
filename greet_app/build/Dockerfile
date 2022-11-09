FROM golang:1.18-alpine AS builder

# Install `git` binary
RUN apk update
RUN apk add git build-base

## Install go tools
RUN go install github.com/go-task/task/v3/cmd/task@latest

WORKDIR /service

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN task compile

FROM alpine:3.14.2

COPY --from=builder ./service/build/out/greet_app greet_app
EXPOSE 50059

ENTRYPOINT ["./greet_app"]