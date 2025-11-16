# ---- Build Stage ----
# Use the official Golang image to build the application.
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies.
# This is done in a separate step to leverage Docker's layer caching.
# Dependencies are only re-downloaded if go.mod or go.sum changes.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application, creating a statically linked binary.
# CGO_ENABLED=0 is important for creating a pure Go binary.
# -o /server names the output file 'server'.
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# ---- Final Stage ----
# Use a minimal, non-root base image for the final container.
# 'scratch' is the most minimal, but 'alpine' is also very small
# and includes shell access, which can be useful for debugging.
FROM alpine:latest

# Set the working directory
WORKDIR /

# Copy ONLY the compiled binary from the 'builder' stage.
# Nothing else from the build environment (source code, tools) is included.
COPY --from=builder /server /server

# Tell Docker that the container listens on this port at runtime.
# Replace '8080' with the port your app uses.
EXPOSE 4000

# The command to run when the container starts.
CMD ["/server"]