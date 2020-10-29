FROM golang:alpine

COPY . /app

WORKDIR /app

RUN go build -o bin/run

CMD ["./bin/run"]