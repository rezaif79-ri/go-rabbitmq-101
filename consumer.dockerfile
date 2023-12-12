FROM golang:1.21-alpine

RUN apk add --no-cache git

RUN mkdir /consumer-app
ADD ./internal /consumer-app
ADD ./app/consumer /consumer-app/app/consumer
WORKDIR /consumer-app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app ./app/consumer/consumer.go

EXPOSE 3000

CMD ["./out/app"]