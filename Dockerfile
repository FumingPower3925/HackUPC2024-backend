# Stage 1: Build the application
FROM golang:1.22.2

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the application
RUN go build cmd/main.go

ENTRYPOINT [ "./main" ]
