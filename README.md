# 📜 Receipt Processor API

🚀 **A RESTful API that processes store receipts and calculates points based on predefined rules.**

[![Go Version](https://img.shields.io/github/go-mod/go-version/Mohit-Jawale/receipt-processor)](https://golang.org)  
[![Docker Ready](https://img.shields.io/badge/Docker-Supported-blue)](https://www.docker.com/)

---

## 📌 Features

✅ **Submit receipts via `POST /receipts/process`**  
✅ **Retrieve receipt points via `GET /receipts/{id}/points`**  
✅ **Docker support for easy deployment**  
✅ **In-memory storage**

---

## 📌 Installation

### 1️⃣ Clone the Repository

```sh
git clone https://github.com/Mohit-Jawale/receipt-processor.git
cd receipt-processor
```

### 2️⃣ Install Dependencies

```sh
go mod tidy
```

### 3️⃣ Run the Server

```sh
go run cmd/server/main.go
```

🎯 **Now the API is running at `http://localhost:8080`**

---

## 📌 API Usage

### 1️⃣ Submit a Receipt

📌 **`POST /receipts/process`**

```sh
curl -X POST "http://localhost:8080/receipts/process"      -H "Content-Type: application/json"      -d '{
         "retailer": "Target",
         "purchaseDate": "2022-01-01",
         "purchaseTime": "13:01",
         "items": [{"shortDescription": "Item 1", "price": "10.00"}],
         "total": "10.00"
     }'
```

🎯 **Response:**

```json
{ "id": "b4a9f89d-7e35-4c9a-86c5-50a5f25d3f37" }
```

---

### 2️⃣ Retrieve Points for a Receipt

📌 **`GET /receipts/{id}/points`**

```sh
curl -X GET "http://localhost:8080/receipts/b4a9f89d-7e35-4c9a-86c5-50a5f25d3f37/points"
```

🎯 **Response:**

```json
{ "points": 88 }
```

---

## 📌 Running with Docker

📌 **Build and Run the Docker Container**

```sh
docker build -t receipt-processor .
docker run -p 8080:8080 receipt-processor
```

🎯 **API now runs inside a Docker container!** 🐳

## 📌 Project Structure

```
receipt-processor/
│── cmd/server/main.go    # Entry point of the application
│── internal/
│   ├── handlers/         # API request handlers
│   ├── models/           # Data models
│   ├── services/         # Business logic
│   ├── storage/          # In-memory storage
│── test/                 # Unit tests
│── Dockerfile            # Docker setup
│── go.mod                # Go dependencies
│── README.md             # Project documentation
```

---

## 📌 Running Tests

```sh
go test ./test
```

🎯 **This ensures all API functionalities are working correctly!** ✅

---

## 📌 License

📜 **MIT License** - Feel free to use and modify the project.

---

🚀 **Happy Coding!** 💻  
For questions, open an issue or reach out via **[LinkedIn](https://www.linkedin.com/in/mohit-jawale-01a48a1aa/)**.
