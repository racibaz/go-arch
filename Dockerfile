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

# Install mockery tool for generating mocks
RUN go install github.com/vektra/mockery/v3@latest

# Install swag tool for API documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest
ENV PATH="/go/bin:${PATH}"

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Generate Swagger documentation
RUN swag init -o api/openapi-spec

# Build the Go app
#RUN go build -o main .

# Expose ports to the outside world
# SERVER_PORT
EXPOSE 3000
# SWAGGER_PORT
EXPOSE 8081
# METRICS_PORT
EXPOSE 9090
# GRPC_PORT
EXPOSE 50051
# Prometheus Exporter Port
EXPOSE 8889
# Web Port
EXPOSE 3001
# Grafana UI Port
EXPOSE 3002
# RabbitMQ Port
EXPOSE 5672
# Fluentd Ports
EXPOSE 24224
# Elasticsearch Ports
EXPOSE 9200

# todo remove if not needed
#EXPOSE 9300
#EXPOSE 9091
#EXPOSE 8080
#EXPOSE 50052

# Entrypoint script
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]