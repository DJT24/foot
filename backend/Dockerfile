# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server/main.go

# Final stage
FROM alpine:3.19
RUN adduser -D app
WORKDIR /app
COPY --from=builder /server .
USER app
EXPOSE 8080
CMD ["/server"]
