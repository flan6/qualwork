# Step 1: Build the Go binary
FROM golang:latest AS builder

WORKDIR /app

ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# Step 2: Build the final image
FROM alpine:latest

RUN apk add --no-cache bash

COPY --from=builder /app/main /bin/main

EXPOSE 8080

RUN chmod +x /bin/main

CMD ["/bin/main"]
