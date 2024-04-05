# Start from a smaller base image like alpine
FROM golang:1.21.5 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Start a new stage from a smaller base image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary from the previous stage
COPY --from=builder /main /

# Expose port 3000 to the outside world
EXPOSE 3000
RUN chmod +x /main
# Command to run the executable
CMD ["/main"]
