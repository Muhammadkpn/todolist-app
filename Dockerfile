FROM golang:1.20 as builder
WORKDIR /app
COPY . .

# Install "air" for hot reload
RUN go install github.com/cosmtrek/air@latest

RUN ls -l /go/bin
RUN go mod tidy
RUN go mod vendor

# Command to run the application using "air" for hot reload during development
CMD ["air", "-c", ".air.toml"]