FROM golang:1.17

# Set necessary environment variables needed for your application here
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY .. .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Command to run the executable
CMD ["./main"]