FROM golang:1.24

WORKDIR /app

COPY . .

RUN go get -v

RUN go build -o main .

CMD ["./main"]