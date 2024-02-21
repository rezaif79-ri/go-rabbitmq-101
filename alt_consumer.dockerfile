FROM golang:1.21-alpine

RUN apk add --no-cache git

RUN mkdir /alt-consumer-app
ADD ./internal /alt-consumer-app
ADD ./app/alt_consumer /alt-consumer-app/app/consumer
WORKDIR /alt-consumer-app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app ./app/alt_consumer/alt_consumer.go

EXPOSE 3000

CMD ["./out/app"]