# 🌐 Node Sphere --- IoT Monitoring & Management Platform

Node Sphere is a lightweight, scalable, and modular platform built to monitor, manage, and analyze IoT devices in real time. It provides a robust backend architecture with clean separation of concerns, seamless health monitoring, and integrations with popular messaging and time-series systems.

---

## 🚀 Features

- ✅ Device registration & lifecycle management  

- 📡 Real-time data ingestion via MQTT / HTTP  

- 📊 Time-series storage with InfluxDB  

- 🔍 Device metadata storage using MongoDB  

- ❤️ Health check system using `health-go` plugins  

- 🧩 Modular, clean architecture with dependency injection  

- 🔐 API-ready for secure device communication and frontend dashboards

---

## 🏗️ Architecture

```bash
📦internal                          # Core of the application
 ┣ 📂application                   # Application logic layer
 ┃ ┗ 📂usecases                    # Business use cases (services)
 ┣ 📂domain                        # Domain layer (enterprise rules)
 ┃ ┣ 📂data_transfer_object       # DTOs for request/response mapping
 ┃ ┣ 📂entities                   # Core entities/models
 ┃ ┣ 📂events                     # Domain events for decoupled logic
 ┃ ┗ 📂repositories               # Repository interfaces contract (ports)
 ┣ 📂infrastructure               # Infrastructure implementations (adapters)
 ┃ ┣ 📂messaging                  # Messaging clients (e.g., MQTT, RabbitMQ)
 ┃ ┣ 📂persistance                # DB adapters (MongoDB, InfluxDB, etc.)
 ┃ ┗ 📂repositories               # Implementations of domain repositories
 ┣ 📂interface                    # Entry points (drivers/controllers)
 ┃ ┣ 📂http                       # HTTP layer
 ┃ ┃ ┣ 📂handler                 # HTTP request handlers/controllers
 ┃ ┃ ┗ 📂router                  # HTTP router setup
 ┃ ┗ 📂mqtt                       # MQTT layer
 ┃ ┃ ┗ 📂consumer                # Consumers for MQTT topics

📦cmd                              # Application entry points
 ┣ 📂web                          # HTTP server entry
 ┃ ┗ 📜main.go                   # Starts HTTP app, injects dependencies
 ┣ 📂worker                       # Worker service entry
 ┃ ┗ 📜main.go                   # Starts background/messaging worker

📦config                           # Centralized configuration
 ┗ 📜config.go                    # Loads env/config variables
```
#### ✨ Built with Clean Architecture for scalability, testability, and separation of concerns.
This system using clean architecture in general but also adapt some concepts and terms of domain driven design to make it more readable, maintainable, structured, and easy to explain. The following picture represent the architecture.
![architecture.png](architecture.png)

#### 🧠 Tech Stack

Layer  Tech/Tool

Language  Go

Routing  net/http

Health Check  health-go

Database  MongoDB

Time-series  InfluxDB

Messaging  MQTT (e.g., Mosquitto, EMQX)

Container  Docker

#### 📦 Setup & Run

Prerequisites

Go ≥ 1.18

Docker & Docker Compose

MongoDB & InfluxDB services running

Optional: MQTT Broker (e.g., EMQX or Mosquitto)

### Run Locally

#### Run HTTP Server
```
go run cmd/web/main.go
```
#### Run MQTT Worker
```
go run cmd/worker/main.go
```
#### 💡 Health Check Example

```http
GET /health
```
healthy response
```json
{
  "status": "OK",
  "timestamp": "2025-04-20T23:27:24.7241761+07:00",
  "system": {
    "version": "go1.22.11",
    "goroutines_count": 17,
    "total_alloc_bytes": 1606928,
    "heap_objects_count": 9437,
    "alloc_bytes": 1606928
  },
  "component": {
    "name": "dms-api-ahmadzeinalwafi",
    "version": "v1"
  }
}
```

## 📈 Roadmap

 WebSocket for real-time dashboard updates & control

 Device grouping and tagging

 User authentication and RBAC

 Alert & anomaly detection module

 Artificial Intelligence Forecasting

## 👨‍💻 Author

@ahmadzeinalwafi

MLOps Engineer | Solution Architect | Creator of Node Sphere

## 📄 License

Apache-2.0 License. See LICENSE for details.