# Docker-Based Client-Server Application

This project demonstrates how to set up a client-server application using Docker and Docker Compose, perform HTTP requests with `curl`, and access the Swagger UI for API documentation.

## Prerequisites

Before running the application, make sure you have the following installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- A terminal (Linux, macOS, or WSL on Windows)
- If on Windows [TDM-GCC Compiler](https://sourceforge.net/projects/tdm-gcc/)

## Step-by-Step Setup Guide

### 1. Create a New Docker Network

First, create a new network bridge using Docker to allow the services to communicate with each other.

```bash
docker network create web
```

### 2. Start the Services
Use Docker Compose to bring up the services:

```bash
docker-compose up -d
```

### 3. Create a New Order (POST)
Send a POST request to create a new order:

```bash
curl -X POST http://localhost:9090/orders \
-H "Content-Type: application/json" \
-d '{
    "customer_name": "Amit",
    "quantity": 123,
    "sku": "abc-def-ghi"
}'
```

### 4. Create a New Shipping (POST)
Send a POST request to create a new shipping:

```bash
curl -X POST http://localhost:9091/api/orders/?num_orders=100 \
-H "Content-Type: application/json" \
-d '{
    "customer_name":"amit",
    "quantity":123,
    "Sku": "abc-def-ghi"
}'
```

### Teardown
To stop the services and remove the containers, you can run:
```bash
docker-compose down
```

### Key Points:
- **Clear setup steps**: For creating the network, starting services, and making API calls.
- **`curl` examples**: Provided for both POST and GET requests with headers and payloads.