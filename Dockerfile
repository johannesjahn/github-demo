# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.25.6-alpine AS build

WORKDIR /src

COPY go.mod ./

COPY . .

RUN CGO_ENABLED=0 go build -o /out/demo main.go

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=build /out/demo .

EXPOSE 3000

CMD ["./demo"]
