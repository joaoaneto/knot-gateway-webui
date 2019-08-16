# Start from the latest golang base image
FROM golang:latest

# Install glide dependency manager
RUN curl https://glide.sh/get | sh

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/CESARBR/knot-gateway-webui/backend

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download development tools
RUN make tools

# Download application dependencies
RUN make deps

# Build the Go App
RUN make bin

# Expose port 8080 to the outside world
EXPOSE 8080



# Command to run the executable
CMD make run