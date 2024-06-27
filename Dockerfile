FROM golang:latest AS builder

WORKDIR /app

COPY . /app


RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/sugar-backend/main.go

EXPOSE 8080

CMD ["./main"]

