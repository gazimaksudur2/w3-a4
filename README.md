# Rental Property API

A production-style RESTful API for managing, retrieving, and searching rental property information, built with **Go** and the **Beego v2 framework**.

The API provides endpoints for listing rental properties, retrieving individual properties by ID, filtering properties using multiple criteria, and exploring and testing endpoints through **Swagger UI**.

---

## Features

* RESTful API architecture using Go and Beego v2
* Property listing and lookup functionality
* Multi-criteria property filtering
* Price range filtering
* Property type and publication status filtering
* Review, rating, bedroom, and feed filtering
* Multiple amenity filtering
* Result limiting
* JSON-based property data source
* Service-layer business logic
* Swagger/OpenAPI documentation
* Interactive API testing through Swagger UI
* Unit testing for service-layer logic
* Go test coverage reporting
* Static code analysis with `go vet`

---

## Technology Stack

| Technology | Purpose                       |
| ---------- | ----------------------------- |
| Go         | Backend programming language  |
| Beego v2   | REST API and web framework    |
| JSON       | Property data source          |
| Swagger    | Interactive API documentation |
| Go Modules | Dependency management         |
| Go Testing | Unit testing and coverage     |

---

# Project Structure

```text
w3-a4/
│
├── conf/
│   └── app.conf                  # Beego application configuration
│
├── controllers/
│   └── property.go               # HTTP request handling and validation
│
├── data/
│   └── rental_property.json      # Rental property source data
│
├── models/
│   ├── filter.go                 # Property filtering model
│   └── ...                       # Request/response and property models
│
├── routers/
│   ├── router.go                 # API route configuration
│   └── commentsRouter.go         # Generated Beego routes
│
├── services/
│   ├── filter.go                 # Property filtering logic
│   ├── transform.go              # Data transformation logic
│   ├── ...                       # Property service operations
│   └── *_test.go                 # Service-layer unit tests
│
├── swagger/
│   ├── swagger.json              # Generated Swagger specification
│   ├── swagger.yml               # Generated Swagger specification
│   └── ...                       # Swagger UI assets
│
├── main.go                       # Application entry point
├── go.mod                        # Go module configuration
├── go.sum                        # Dependency checksums
├── .gitignore
└── README.md
```

---

# Architecture

The application follows a layered structure:

```text
Client Request
      │
      ▼
    Router
      │
      ▼
  Controller
      │
      ▼
   Service
      │
      ▼
JSON Data Source
      │
      ▼
 Service Transformation
      │
      ▼
 JSON Response
```

Controllers are responsible for HTTP request handling and query-parameter validation, while the service layer contains the main filtering, lookup, and transformation logic.

---

# Prerequisites

Before running the project, make sure you have:

* Go 1.25 or higher
* Beego CLI installed
* Git installed

Check your Go installation:

```bash
go version
```

Install the Beego CLI if necessary:

```bash
go install github.com/beego/bee/v2@latest
```

Verify the installation:

```bash
bee version
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

Install and synchronize dependencies:

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

```text
http://localhost:8080
```

---

# Swagger API Documentation

The project includes Swagger documentation for the available API routes.

After starting the application, open:

```text
http://localhost:8080/swagger/
```

Swagger UI allows you to:

* View available endpoints
* Inspect query and path parameters
* Execute API requests directly from the browser
* Inspect API response structures
* Test different filter combinations

---

# API Base URL

```text
http://localhost:8080/v1
```

---

# API Endpoints

## 1. Get All Properties

Retrieves rental properties.

### Request

```http
GET /v1/properties
```

### Example

```bash
curl "http://localhost:8080/v1/properties"
```

---

## 2. Get Property by ID

Retrieves a specific property using its unique identifier.

### Request

```http
GET /v1/properties/:id
```

### Example

```bash
curl "http://localhost:8080/v1/properties/BC-1000001"
```

If the property does not exist, the API returns a `404` response.

---

# Property Filtering

The `GET /v1/properties` endpoint supports optional query parameters that can be combined to refine the results.

| Parameter          | Description                | Example            |
| ------------------ | -------------------------- | ------------------ |
| `min_price`        | Minimum property price     | `50`               |
| `max_price`        | Maximum property price     | `150`              |
| `min_review_score` | Minimum review score       | `4.0`              |
| `feed`             | Property feed              | `11`               |
| `published`        | Publication status         | `true`             |
| `property_type`    | Property category/type     | `Hotel`            |
| `min_star_rating`  | Minimum star rating        | `4`                |
| `min_reviews`      | Minimum number of reviews  | `100`              |
| `min_bedroom`      | Minimum number of bedrooms | `2`                |
| `amenities`        | Comma-separated amenities  | `Internet,Parking` |
| `limit`            | Maximum number of results  | `5`                |

---

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

## Filter by Price Range

```bash
curl "http://localhost:8080/v1/properties?min_price=50&max_price=150"
```

---

## Filter by Property Type

```bash
curl "http://localhost:8080/v1/properties?property_type=Hotel"
```

---

## Filter by Review Score

```bash
curl "http://localhost:8080/v1/properties?min_review_score=4"
```

---

## Filter by Star Rating

```bash
curl "http://localhost:8080/v1/properties?min_star_rating=4"
```

---

## Filter by Number of Reviews

```bash
curl "http://localhost:8080/v1/properties?min_reviews=100"
```

---

## Filter by Minimum Bedrooms

```bash
curl "http://localhost:8080/v1/properties?min_bedroom=2"
```

---

## Filter by Amenities

Multiple amenities can be provided as comma-separated values.

```bash
curl "http://localhost:8080/v1/properties?amenities=Internet,Parking"
```

The current service implementation uses **OR logic** for multiple amenities. A property is included when it contains at least one of the requested amenities.

---

## Multiple Filters

Filters can be combined in a single request.

```bash
curl "http://localhost:8080/v1/properties?feed=11&published=true&min_price=50&max_price=150"
```

A more complex example:

```bash
curl "http://localhost:8080/v1/properties?feed=11&min_star_rating=4&min_review_score=4&amenities=Internet,Parking&limit=5"
```

---

## Limit Results

Control the maximum number of returned properties:

```bash
curl "http://localhost:8080/v1/properties?limit=5"
```

---

## Testing

The project uses Go's testing framework with table-driven unit tests for core service-layer business logic.

Implemented test coverage includes:

- Property filtering scenarios
- Property lookup scenarios
- Property transformation validation
- Result limiting behavior

Run tests:

```bash
go test ./... -v
```

---

## Run Tests with Coverage

To display test coverage for every package:

```bash
go test ./... -cover
```

Example output:

```text
w3-a4/controllers    coverage: 0.0% of statements
w3-a4/routers        coverage: 0.0% of statements
ok  w3-a4/services   coverage: 84.0% of statements
```

Coverage is currently focused on the **service layer**, where the core business logic resides.

---

## Generate a Coverage Profile

Generate a detailed coverage file:

```bash
go test ./... -coverprofile=coverage.out
```

View the coverage summary by function:

```bash
go tool cover -func=coverage.out
```

---

## View Coverage in the Browser

Generate and open an interactive HTML coverage report:

```bash
go tool cover -html=coverage.out
```

This highlights covered and uncovered statements in the source code.

---

# Static Code Analysis

Run Go's built-in static analyzer:

```bash
go vet ./...
```

---

# API Response Format

The API returns JSON responses.

A property response contains structured property, geographic, pricing, review, amenity, and related information.

Example structure:

```json
{
  "id": "BC-1000001",
  "feed": 11,
  "published": true,
  "geo_info": {
    "city": "Shinjuku",
    "country": "Japan"
  },
  "property_info": {
    "name": "Shinjuku Grand Resort",
    "property_type": "Resort",
    "price": 116.69
  }
}
```

The exact response schema can be inspected through Swagger UI.

---

# Error Handling

The API validates supported query parameters and returns appropriate HTTP error responses for invalid requests.

Examples include:

```text
400 Bad Request
404 Not Found
500 Internal Server Error
```

For example, invalid numeric filter values are rejected with a `400 Bad Request`, while requesting a property ID that does not exist results in `404 Not Found`.

---

# Development Commands

Common commands used during development:

```bash
# Run the application
bee run

# Run all tests
go test ./... -v

# Run tests with coverage
go test ./... -cover

# Run static analysis
go vet ./...

# Synchronize dependencies
go mod tidy
```
---

# Author

Developed by **Gazi Maksudur**

GitHub:
https://github.com/gazimaksudur2/w3-a4
