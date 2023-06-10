# syntax=docker/dockerfile:1

FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
# TODO: Do not forget to replace with 'COPY go.mod go.sum ./' once you add a Third-party library
COPY go.mod ./
RUN go mod download

# Copy the source code.
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /slash

# Run
CMD ["/slash"]