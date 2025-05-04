FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o server main.go

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/static ./static
COPY --from=builder /app/data ./data

EXPOSE 8080

CMD ["./server"]
