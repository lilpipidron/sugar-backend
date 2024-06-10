FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sugar-backend ./cmd/sugar-backend/main.go

EXPOSE 8080

CMD ["./main.go"]
