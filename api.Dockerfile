# Use an official Golang runtime as a parent image
FROM golang:1.21.1-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application inside the container
RUN go build -o api .

# Expose port 8080 for the application
EXPOSE 8080

# Define the command to run your application
CMD ["./api", "server"]
