# Backend Assignment 1 — Go (Gin + Kafka + MySQL)

A RESTful backend service built with **Go**, using the **Gin** web framework, **Kafka** for event streaming, and **MySQL** as the database.

---

## 🗂 Project Structure

```
Backend_assigment_1/
├── cmd/
│    └── main.go                      # Application entry point
├── controller/
│   ├── OrdersController.go           # Order HTTP handlers
│   └── ProductController.go          # Product HTTP handlers
├── DatabaseConnection/
│   └── DatabaseConnection.go         # MySQL connection setup
├── docs/                             # Documentation files
├── kafka/
│   └── producer.go                   # Kafka producer initialization
├── migration/
│   ├── migrate.go                    # Database migrations
│   └── Seed.go                       # Product seeding
├── models/
│   ├── OrderModels.go                # Order entity/model
│   ├── ProductModels.go              # Product entity/model
│   └── User.go                       # User entity/model
├── routes/
│   ├── OrderRoutes.go                # Order route definitions
│   └── ProductRoutes.go              # Product route definitions
├── services/
│   ├── OrderServices.go              # Order business logic
│   └── ProductServices.go            # Product business logic
├── .env                              # Environment variables
├── .gitignore
└── go.mod
```

---

## ⚙️ Prerequisites

Make sure the following are installed on your machine:

- [Go](https://golang.org/dl/) `>= 1.21`
- [MySQL](https://dev.mysql.com/downloads/) `>= 8.0`
- [Apache Kafka](https://kafka.apache.org/downloads)
- [Git](https://git-scm.com/)

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Ayushkm26/backend_assigment_1.git
cd Backend_assigment_1
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Set Up Environment Variables

The `.env` file is already present in the root. Fill in your values:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name

KAFKA_BROKER=localhost:9092
```

### 4. Start Kafka (if running locally)

```bash
# Start Zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties

# Start Kafka broker (in a new terminal)
bin/kafka-server-start.sh config/server.properties
```

---

## 🔨 Build

> ⚠️ Note: `main.go` is located inside `cmd/cmd/` directory

```bash
# Build the binary
go build -o app ./cmd/cmd/main.go
```

This creates an executable named `app` in the current directory.

---

## ▶️ Run

### Option 1 — Run directly without building

```bash
go run ./cmd/cmd/main.go
```

### Option 2 — Run the built binary

```bash
# Step 1: Build
go build -o app ./cmd/cmd/main.go

# Step 2: Run
./app
```

The server starts on:
```
http://localhost:8080
```

---

## 🌐 API Endpoints

Base URL: `http://localhost:8080/api`

### Products

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/products/` | Get all products |

### Orders

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/orders/createorder` | Create a new order |
| GET | `/api/orders/getOrdersById/:id` | Get order by ID |

---

## 🔒 CORS Configuration

This service allows cross-origin requests from the Java Spring backend running on port `9000`:

```
Allowed Origin:  http://localhost:9000
Allowed Methods: GET, POST, PUT, DELETE, OPTIONS
Allowed Headers: Origin, Content-Type, Authorization
Max Age:         12 hours
```

---

## 🧰 Tech Stack

| Technology | Purpose |
|------------|---------|
| [Go](https://golang.org/) | Core language |
| [Gin](https://github.com/gin-gonic/gin) | HTTP web framework |
| [gin-contrib/cors](https://github.com/gin-contrib/cors) | CORS middleware |
| [Apache Kafka](https://kafka.apache.org/) | Event streaming |
| [MySQL](https://www.mysql.com/) | Relational database |

---

## 📦 Useful Go Commands

```bash
# Download / tidy dependencies
go mod tidy

# Run tests
go test ./...

# Check for issues
go vet ./...

# Format all code
go fmt ./...
```

---

## 📬 API Request Examples

Base URL: `http://localhost:8080/api`

---

### POST `/api/orders/createorder`

**✅ Valid Request**
```json
POST /api/orders/createorder
Content-Type: application/json

{
"userId": 1,
"items": [
{ "productId": 101, "quantity": 2 },
{ "productId": 205, "quantity": 1 }
]
}
```
**Response — 200 OK**
```json
{
  "orderId": 55,
  "userId": 1,
  "status": "PENDING",
  "items": [
    { "productId": 101, "quantity": 2 },
    { "productId": 205, "quantity": 1 }
  ]
}
```

**❌ Invalid — Missing `userId`**
```json
{
  "items": [{ "productId": 101, "quantity": 2 }]
}
```
```json
{ "error": "userId is required" }
```

**❌ Invalid — Wrong type (`userId` as string)**
```json
{
  "userId": "john",
  "items": [{ "productId": 101, "quantity": 2 }]
}
```
```json
{ "error": "json: cannot unmarshal string into Go struct field .userId of type uint" }
```

**❌ Invalid — Empty `items` array**
```json
{
  "userId": 1,
  "items": []
}
```
```json
{ "error": "items cannot be empty" }
```

**❌ Invalid — Completely empty body**
```json
{}
```
```json
{ "error": "invalid request body" }
```

---

### GET `/api/orders/getOrdersById/:id`

**✅ Valid Request**
```
GET /api/orders/getOrdersById/55
```
**Response — 200 OK**
```json
{
  "orderId": 55,
  "userId": 1,
  "status": "PENDING",
  "totalAmount": 199.99,
  "items": [
    { "productId": 101, "quantity": 2 },
    { "productId": 205, "quantity": 1 }
  ]
}
```

**❌ Invalid — Order ID does not exist**
```
GET /api/orders/getOrdersById/9999
```
```json
{ "error": "order not found" }
```

**❌ Invalid — ID is a string**
```
GET /api/orders/getOrdersById/abc
```
```json
{ "error": "invalid order ID format" }
```

**❌ Invalid — Negative ID**
```
GET /api/orders/getOrdersById/-1
```
```json
{ "error": "ID must be a positive integer" }
```

---

### GET `/api/products/`

**✅ Valid Request**
```
GET /api/products/
```
**Response — 200 OK**
```json
[
  { "id": 101, "name": "Wireless Mouse",       "price": 29.99, "stock": 150 },
  { "id": 102, "name": "Mechanical Keyboard",  "price": 89.99, "stock": 75  }
]
```

**❌ Invalid — DB is down / no products seeded**
```json
{ "error": "failed to fetch products" }
```

**❌ Invalid — Wrong HTTP method (POST instead of GET)**
```
POST /api/products/
```
```
404 page not found
```

---

## 📝 Notes

- On first run, **migrations are applied and products are seeded automatically** via `migrate.go` and `Seed.go`.
- **Kafka producer** is initialized at startup — ensure Kafka is running before starting the app.
- The main entry point is `cmd/main.go`.
- CORS is configured to allow requests from the **Java Spring backend** on `localhost:9000`.
