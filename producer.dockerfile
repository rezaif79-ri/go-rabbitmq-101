FROM golang:1.21-alpine

RUN apk add --no-cache git

RUN mkdir /producer-app
ADD ./internal /producer-app
ADD ./app/producer /producer-app/app/producer
WORKDIR /producer-app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app ./app/producer/producer.go

EXPOSE 8001

CMD ["./out/app"]