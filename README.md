# go-microservices

This repository contains three microservices, a front end for user interaction with the backend services, a publisher service with an api used to push messages to Kafka, and a subscription service that subscribes to Kafka and feeds the messages back to the front end through a websocket. There is also a name generator service that provides randomly generated names.

Communication to the backend services is routed through a Kong API Gateway and messages are either streamed to the frontend through a websocket or called to via REST API.

The purpose of this is purely for learning, so the code is in no way optimised or even close to production ready.

## Design

[Architectural Diagram](./.diagram.jpeg)

## Configuration

Copy the `.env.example` file to `.env` and set relevant values.

Note: The config/kong.yaml configuration still uses hardcoded port numbers to setup each service's route. 

## Installation

Start the API Gateway, Kafka, and microservices with docker-compose:

```
docker-compose up -d
```

Navigate to [http://localhost:8080](localhost:8080) to view the front-end.

## Limitations

- One websocket client at a time. A map should be created and messages pushed to all clients.
