# Start from golang base image
FROM golang:1.15-alpine as builder

# Set the current working directory inside the container
WORKDIR /build

# Copy go.mod, go.sum files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy sources to the working directory
COPY . .

# Build the Go app
ARG project
#RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server $project
RUN env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o server $project

# Start a new stage from busybox
FROM busybox:latest

WORKDIR /dist

# Copy the build artifacts from the previous stage
COPY --from=builder /build/server .
# Give executable permissions
RUN chmod +x ./server
# Run the executable
CMD ["./server"]