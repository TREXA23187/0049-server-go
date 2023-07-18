# Use an official Golang runtime as a parent image
FROM golang:1.19

# Set the working directory in the container to /app
WORKDIR /app

ADD . /app

# Download all dependencies. If the go.mod and go.sum files
# are not changed, Docker will reuse the cached layer.
RUN go mod download

# Build the Go app
RUN go build -o main .

# Run the Go app
CMD ["./main"]

# This container exposes port 8080 to the outside world
EXPOSE 8080
