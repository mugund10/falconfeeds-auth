FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o falconfeeds-auth ./main.go

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/falconfeeds-auth .

EXPOSE 8080
CMD ["./falconfeeds-auth"]
