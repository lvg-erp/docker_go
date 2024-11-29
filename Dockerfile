# Build stage
FROM golang:1.23-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o server ./cmd/server

# Final stage
FROM alpine:latest

# Set the working directory inside the final container
WORKDIR /app

# Copy the binary built in the previous stage
COPY --from=build /app/server .

# Expose the port your application will listen on (adjust as needed)
EXPOSE 8080

# Run your Go application
CMD ["./server"]