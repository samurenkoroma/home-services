# syntax=docker/dockerfile:1

FROM golang:1.25.5-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# RUN GOARCH=arm64 go build -o application cmd/main.go
RUN go build -o application cmd/main.go

FROM alpine

COPY --from=build /app/application /app/application

CMD ["/app/application"]