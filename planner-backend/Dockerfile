### Builder stage ###
FROM golang:1.21-alpine as builder

WORKDIR /usr/local/go/src/planner-backend

# Copy go mod and sum files
COPY go.* .

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main ./cmd

### Final stage ###
FROM alpine:latest

# Create appuser
RUN addgroup -g 10001 -S appuser && \
    adduser -u 10001 -S appuser -G appuser

# Use an unprivileged user.
USER appuser

# Copy the binary file from the builder stage
COPY --from=builder /app/main /app/main

# Expose port 8080
EXPOSE 8080

# Set the GIN_MODE environment variable
ENV GIN_MODE=release

# Run the binary
ENTRYPOINT [ "/app/main", "--prod" ]