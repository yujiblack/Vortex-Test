# Step 1: Build the Go binary
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Copy the module files and source code
COPY go.mod ./
COPY *.go ./

# Build the executable (This will FAIL until the code is fixed!)
RUN go build -o vox-test .

# Step 2: Run the binary in a tiny container
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/vox-test .

# Run the app
CMD ["./vox-test"]