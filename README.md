# Simple-Kafka-Delivery-System-Go-Kafka-
A Go project demonstrating **asynchronous order processing** using Kafka and a worker pool.   Orders are produced into a Kafka topic, consumed by a Go consumer, and processed in parallel by multiple workers using channels and goroutines.
## Run Locally

### Start Kafka and Zookeeper:

```bash
docker-compose up -d
```
### Run the application:
```
go run cmd/main.go
```

You’ll see output like:
```
Worker 1 Processing Order 1
Worker 2 Processing Order 2
Worker 3 Processing Order 3
```
...

### Project Structure
cmd/main.go          # Entry point <br>
kafka/producer.go    # Kafka producer <br>
kafka/consumer.go    # Kafka consumer <br>
worker/worker.go     # Worker pool & processing <br>
model/order.go       # Order struct <br>
docker-compose.yml   # Kafka + Zookeeper setup <br>

## Features

-> Parallel processing of orders using a worker pool

-> No duplicate processing (each order is processed by one worker)

-> Easy to scale workers by increasing count

-> Fully containerized Kafka environment for local testing
