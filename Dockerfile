# syntax=docker/dockerfile:1
FROM golang:1.17

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o backend-service ./app/

EXPOSE 8080

CMD ["./backend-service"]