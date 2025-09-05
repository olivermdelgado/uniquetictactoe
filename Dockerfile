# -------- build --------
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .

RUN go build -o app main.go

# -------- run --------
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/app .

# Run the program
CMD ["./app"]