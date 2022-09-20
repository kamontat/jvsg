FROM golang:1.19.1-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY *.go ./

RUN go build -o start-client

FROM alpine:3.16.2

WORKDIR /

COPY --from=build /app/start-client /start-client

ENTRYPOINT ["/start-client"]
