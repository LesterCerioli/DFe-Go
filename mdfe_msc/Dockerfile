# Use the official Go image for building the Go app
FROM golang:1.18-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download the dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o mdfe_msc

# Use a minimal base image to run the built app
FROM alpine:latest

# Set the working directory inside the runtime container
WORKDIR /root/

# Copy the built Go app from the previous stage
COPY --from=build /app/mdfe_msc .

# Expose the port that the Fiber app will run on
EXPOSE 3000

# Command to run the Go app
CMD ["./mdfe_msc"]
