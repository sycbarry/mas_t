# Use the official Golang image as base
FROM golang:latest as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules and dependencies file
COPY go.mod .
COPY go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory in the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/app .

# Expose port 8080 to the outside world
EXPOSE 8080
EXPOSE 80 
EXPOSE 443

# Command to run the executable
# CMD ["./app"]

