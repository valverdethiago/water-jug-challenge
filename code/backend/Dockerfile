# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:1.18 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container.
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Build the Go app.
RUN go build -o backend .

# This is the second stage, where the compiled application is run.
# Use a smaller image to reduce the final size.
FROM gcr.io/distroless/base-debian10

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/backend .

# This container exposes port 8080 to the outside world.
EXPOSE 8080

# Run the binary program produced by 'go install'.
CMD ["./backend"]
