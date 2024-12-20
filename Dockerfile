# Use the official Golang image
FROM golang:1.23

# Set the working directory
WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Command to run the app
CMD ["air", "-c", ".air.toml"]
