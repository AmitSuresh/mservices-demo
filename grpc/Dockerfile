# Use the official Go image as the base
FROM golang:latest

# Install any required packages, including ca-certificates
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Create a non-root user and switch to it
RUN addgroup --system app && \
    adduser --system --ingroup app app

# Set the working directory
WORKDIR /app

# Copy the application source code and necessary files
COPY . .

# Install grpc-health-probe
RUN go install github.com/grpc-ecosystem/grpc-health-probe@latest

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -o server ./main.go

# Change ownership of the application files to the non-root user
RUN chown -R app:app /app

# Switch to the non-root user
USER app

# Expose the application port
EXPOSE 9082

# Set the entrypoint for the application
ENTRYPOINT ["/app/server"]

