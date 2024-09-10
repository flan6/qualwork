# Step 1: Build the Go binary
FROM golang:latest AS builder

WORKDIR /build

ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# Step 2: Build the final image
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache bash

COPY --from=builder /build/main main
COPY --from=builder /build/static static/
COPY --from=builder /build/assets assets/

EXPOSE 8080

RUN chmod +x /app/main

CMD ["/app/main"]
