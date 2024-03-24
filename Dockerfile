FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sugar-backend ./cmd/sugar-backend/sugar-backend.go

EXPOSE 8080

CMD ["./sugar-backend"]
