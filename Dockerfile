# FROM golang:1.21 as builder
# WORKDIR /app
# COPY . .

# # Install "air" for hot reload
# RUN go install github.com/cosmtrek/air@v1.49.0

# RUN ls -l /go/bin
# RUN go mod tidy
# RUN go mod vendor

# # Command to run the application using "air" for hot reload during development
# CMD ["air", "-c", ".air.toml"]

# Stage 1: Build stage
FROM golang:1.22 AS build

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY . .
# RUN go mod tidy

# Copy the source code
#COPY . .
RUN ls -l

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/http/main.go


# Stage 2: Final stage
FROM alpine:edge

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/myapp .
# COPY --from=build /app/variables ./variables
# COPY --from=build /app/logs ./logs

# Set the timezone and install CA certificates
# RUN apk --no-cache add ca-certificates tzdata

# Set the entrypoint command
ENTRYPOINT ["/app/myapp"]