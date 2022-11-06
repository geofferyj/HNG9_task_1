## Build
FROM golang:1.16-alpine AS build

WORKDIR /app

COPY main.go ./
COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN go build -o /hng9-t1

## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build /hng9-t1 /hng9-t1

EXPOSE 80

ENTRYPOINT ["/hng9-t1"]