FROM golang:1.19.1-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY *.go ./
COPY large.json ./

RUN go build -o start-server

FROM alpine:3.16.2

WORKDIR /

COPY --from=build /app/start-server /start-server
COPY --from=build /app/large.json /large.json

ENTRYPOINT ["/start-server"]
