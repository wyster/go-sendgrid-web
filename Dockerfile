FROM golang:alpine

COPY . /app

WORKDIR /app

RUN go build -o bin/run

EXPOSE 80/tcp

CMD ["./bin/run"]