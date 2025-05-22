# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app .

# Final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
COPY --from=builder /app/views ./views
COPY --from=builder /app/src ./src
COPY --from=builder /app/assets ./assets
CMD ["./app"]
