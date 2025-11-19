# Base Image
FROM golang:1.25-alpine

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

# Install migrate tool for database migrations
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
#RUN go build -o main .

# Expose ports to the outside world
EXPOSE 3000
EXPOSE 8080
EXPOSE 8081
EXPOSE 9090
EXPOSE 9091
EXPOSE 50051
EXPOSE 50052
EXPOSE 8889
EXPOSE 3001
EXPOSE 5672

# Entrypoint script
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]