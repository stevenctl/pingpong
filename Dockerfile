FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Build the application
RUN go mod download
RUN go build -o client ./cmd/client
RUN go build -o server ./cmd/server

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/client .
RUN cp /build/server .

FROM alpine

COPY --from=builder /dist/client /
COPY --from=builder /dist/server /

# Command to run
CMD ["/server"]