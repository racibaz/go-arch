# Base Image
FROM golang:1.24-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Install dependencies
RUN apk add make # install make
RUN apk add curl # install curl for testing


# Add Maintainer Info
LABEL maintainer="Recai Cansız <r.c67@hotmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD air
