# Rental Property API

A production-style RESTful API for managing and searching rental property information, built with **Go** and the **Beego v2 framework**.

The API provides endpoints for retrieving rental property data, searching properties using multiple filters, and accessing interactive API documentation through **Swagger UI**.

---

## Features

* RESTful API architecture using Go and Beego v2
* Property listing and lookup functionality
* Search and filtering support
* In-memory data processing from JSON source data
* Swagger API documentation
* API testing support through Swagger UI
* Unit testing and static code analysis support

---

## Technology Stack

| Technology        | Purpose                                |
| ----------------- | -------------------------------------- |
| Go                | Backend programming language           |
| Beego v2          | Web framework and REST API development |
| JSON              | Property data source                   |
| Swagger           | API documentation                      |
| Go Modules        | Dependency management                  |

---

# Project Structure

```
w3-a4/
│
├── controllers/        # API controllers
├── routers/            # API route configuration
├── models/             # Data models
├── swagger/            # Swagger generated documentation assets
├── main.go             # Application entry point
├── go.mod              # Go dependency configuration
└── README.md
```

---

# Prerequisites

Before running the project, make sure you have:

* Go 1.22 or higher
* Beego framework installed
* Git installed

Check your Go installation:

```bash
go version
```

---

# Installation

Clone the repository:

```bash
git clone https://github.com/gazimaksudur2/w3-a4.git
```

Navigate into the project directory:

```bash
cd w3-a4
```

Install dependencies:

```bash
go mod tidy
```

---

# Running the Application

Start the Beego development server:

```bash
bee run
```

The API will be available at:

```
http://localhost:8080
```

---

# Swagger API Documentation

This project includes Swagger documentation for all available API routes.

After starting the application, open:

```
http://localhost:8080/swagger/
```

Swagger UI allows you to:

* View all available endpoints
* Check request parameters
* Test API calls directly from the browser
* Review API responses

---

# API Base URL

```
http://localhost:8080/v1
```

---

# Available API Endpoints

## 1. Get All Properties

Retrieves a list of rental properties.

### Request

```
GET /v1/properties
```

### Example

```bash
curl http://localhost:8080/v1/properties
```

---

## 2. Get Property By ID

Retrieves a specific property using its unique identifier.

### Request

```
GET /v1/properties/:id
```

### Example

```bash
curl http://localhost:8080/v1/properties/BC-1000001
```

---

# Filtering API

The property listing endpoint supports multiple query filters.

---

## Filter By Feed

```bash
curl "http://localhost:8080/v1/properties?feed=11"
```

---

## Filter By Published Status

```bash
curl "http://localhost:8080/v1/properties?published=false"
```

---

## Multiple Filters

Example:

```bash
curl "http://localhost:8080/v1/properties?feed=11&published=false"
```

---

## Price Range Filter

```bash
curl "http://localhost:8080/v1/properties?min_price=50&max_price=150"
```

---

## Property Type Filter

```bash
curl "http://localhost:8080/v1/properties?property_type=Hotel"
```

---

## Amenities Filter

Amenities filtering supports multiple values.

Example:

```bash
curl "http://localhost:8080/v1/properties?amenities=Internet,Parking"
```

---

## Limit Results

Control the number of returned records:

```bash
curl "http://localhost:8080/v1/properties?limit=5"
```

---

# Testing

Run all tests:

```bash
go test ./... -v
```

Run Go static analysis:

```bash
go vet ./...
```

---

# API Response Format

The API returns JSON responses.

Example:

```json
{
    "id": "BC-1000001",
    "property_type": "Hotel",
    "published": true,
    "price": 120,
    "amenities": [
        "Internet",
        "Parking"
    ]
}
```

---

# Future Improvements

Possible future enhancements:

* Database integration (PostgreSQL/MySQL)
* Authentication and authorization
* JWT-based security
* Pagination support
* Docker containerization
* Cloud deployment
* Automated CI/CD pipeline

---

# Author

Developed by **Gazi Maksudur**

GitHub:

https://github.com/gazimaksudur2/w3-a4

---

# License

This project is created for educational and development purposes.
