# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.18-buster as builder

# Create and change to the app directory.
WORKDIR /app

# Copy the Go module files.
COPY go.mod go.sum ./

# Download dependencies.
RUN go mod download

# Copy the rest of the application source code.
COPY . .

# Build the application.
RUN go build -o Openapi-todo-app ./cmd/

# Use a multi-stage build to create a lean production image.
FROM debian:buster-slim

# Install required libraries.
RUN apt-get update && apt-get install -y --no-install-recommends \
	ca-certificates \
	&& rm -rf /var/lib/apt/lists/*

# Expose port 3000.
EXPOSE 3000

# Copy the binary from the builder stage.
COPY --from=builder /app/go-starter-server /app/go-starter-server

# Self signed certificate
COPY cert.pem /etc/cert.pem
COPY key.pem /etc/key.pem

# Run the binary.
CMD ["/app/go-starter-server"]

