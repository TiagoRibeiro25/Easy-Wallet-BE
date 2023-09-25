# Use the official Golang image as a builder stage
FROM golang:1.21.1 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the rest of the source code
COPY ./src ./src

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./dist/easywalletapi ./src

# Expose port 5000 to the outside world
EXPOSE 5000

# Command to run the executable from the ./dist/ directory
CMD ["./dist/easywalletapi"]
