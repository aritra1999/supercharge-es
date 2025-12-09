# Stage 2: Build binary
FROM golang:1.25-alpine AS builder

RUN adduser -D builderuser

WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Stage 2: Run
FROM scratch

WORKDIR /app

COPY --from=builder /app/server .

USER 1000
EXPOSE 8000

ENTRYPOINT ["./server"]
