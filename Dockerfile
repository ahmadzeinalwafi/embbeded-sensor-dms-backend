# Step 1: Build the Go binary
FROM golang:1.22-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire project
COPY . .

# Build the Go binary (assuming the entry point is cmd/web/main.go)
RUN go build -o main ./cmd/web/main.go

# Step 2: Create the final minimal image to run the app
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/
COPY config.yml /config.yml

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your app runs on
EXPOSE 8888

# Command to run the app
CMD ["./main"]
