## Go RabbitMQ 101

First encounter with wild rabbit mq in golang universe

I use docker to run rabbitmq server with this step-by-step setup:
- Run `docker pull rabbitmq:management` to pull the images
- Once its ready, Run `docker run -d — name dev-rabbit — hostname rabbitmq-dev -p 15672:15672 -p 5672:5672 rabbitmq:management`

## Summary

- Create a producer message and publish to rabbitmq 
- To consume the message, create consumer and listen to rabbitmq

