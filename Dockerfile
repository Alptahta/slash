# syntax=docker/dockerfile:1

FROM golang:1.20 AS build-stage

# Set destination for COPY
WORKDIR /app

# Copy the source code.
# TODO: Do not forget to replace with 'COPY go.mod go.sum ./' once you add a Third-party library
COPY . ./

# Download Go modules
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /slash

# Run
CMD ["/slash"]