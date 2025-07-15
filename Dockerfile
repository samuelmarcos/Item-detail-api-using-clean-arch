# Build stage
FROM golang:latest as builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY src/cmd ./src/cmd
COPY src/internal ./src/internal
COPY src/docs ./src/docs

# Build the application
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./src/cmd/main.go

# Run stage
FROM scratch
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"] 