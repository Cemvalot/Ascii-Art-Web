# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

LABEL maintainer="gpatoula, pkalliag, cemvalot"
LABEL version="1.0"
LABEL description="An ASCII Art Web Application Dockerize"

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application with verbose output
RUN go build -v -o ascii-art-web-dockerize .

# Stage 2: Create a minimal image with the binary
FROM alpine:3.18

# Set the working directory
WORKDIR /root/

# Copy the binary and necessary directories from the builder stage
COPY --from=builder /app/ascii-art-web-dockerize .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/banners ./banners
COPY --from=builder /app/static ./static

# Install bash and core utilities (optional)
RUN apk add --no-cache bash

# Expose the port your Go server listens on
EXPOSE 8080

# Command to run the binary
CMD ["./ascii-art-web-dockerize"]
