
# Rental Property API

A RESTful Rental Property API built using **Go** and **Beego**.  
The API loads rental property data from a JSON file into memory and provides endpoints for listing properties with filters and retrieving a single property by ID.

## Technology

- Go
- Beego v2
- JSON file as data source
- In-memory data storage

## Project Setup

### Prerequisites

Make sure you have:

- Go 1.22 or higher
- Beego installed

Check Go version:

```bash
go version
````

## Install Dependencies

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/gazimaksudur2/w3-a4.git

cd w3-a4
```

Install dependencies:

```bash
go mod tidy
```

## Run the Application

Start the Beego server:

```bash
bee run
```

The API will start on:

```
http://localhost:8080
```

---

# API Endpoints

## 1. Get All Properties

### Endpoint

```
GET /v1/properties
```

### Example

```bash
curl http://localhost:8080/v1/properties
```

---

## 2. Get Property By ID

### Endpoint

```
GET /v1/properties/:id
```

### Example

```bash
curl http://localhost:8080/v1/properties/BC-1000001
```

---

# Filtering Examples

## Filter by Feed

```bash
curl "http://localhost:8080/v1/properties?feed=11"
```

---

## Filter by Published Status

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

Amenities use OR logic.

Example:

```bash
curl "http://localhost:8080/v1/properties?amenities=Internet,Parking"
```

---

## Limit Results

```bash
curl "http://localhost:8080/v1/properties?limit=5"
```

---

# Running Tests

Run all tests:

```bash
go test ./... -v
```

Run static analysis:

```bash
go vet ./...
```

---

## Why this README fits your submission

It covers the required submission points:

✅ How to install/run  
✅ `bee run` command  
✅ API examples  
✅ Filter examples  
✅ Test command  
✅ No unnecessary explanation  

The assignment submission guideline asks specifically for:
- README with setup instructions
- sample curl list
- tests passing
- server starts with `bee run`  
:contentReference[oaicite:1]{index=1}

Next after adding this README, I would do:

```bash
git status
git add README.md
git commit -m "finalize project readme with setup and api examples"
````