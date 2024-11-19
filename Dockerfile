FROM golang:1.23.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bank

FROM scratch
COPY --from=builder /app/bank /app/bank

EXPOSE 8080

ENTRYPOINT ["/app/bank"]