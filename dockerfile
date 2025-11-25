# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o /calculator ./cmd/calculator

# Run stage (scratch = minimal)
FROM scratch
COPY --from=builder /calculator /calculator
ENTRYPOINT ["/calculator"]
