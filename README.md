# go-microservices

This repository contains three microservices, a front end for user interaction with the backend services, a publisher service with an api used to push messages to Kafka, and a subscription service that subscribes to Kafka and feeds the messages back to the front end through a websocket.

The purpose of this is purely for learning so the code is in no way optimised or even close to production ready. 

With this I've explored pub/sub architecture with Kafka, routing through an API gateway with Kong, Websockets, and the Go language. 

## Design

[Architectural Diagram](./.diagram.jpeg)

## Installation

### API Gateway

Install the API gateway:

```
cd api-gateway
docker-compose up -d
```

Configure the services and routes in the kong api gateway at http://localhost:8000

### Go Microservices

Install the microservices:

```
docker-compose up -d
```


## Limitations

- One websocket client at a time. A map should be created and messages pushed to all clients.