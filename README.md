# Todo API

This is a simple RESTful API for managing a todo list, built with Go and the Gin web framework.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Docker](#docker)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

- Go 1.18 or higher
- Docker (optional, for containerization)

## Installation

1. Clone the repository:

```sh
git clone https://github.com/Ashikurrahaman287/RESTful-API-.git
cd todo-api
```

2. Initialize the Go module and download dependencies:

```sh
go mod tidy
```

3. Install Gin framework:

```sh
go get github.com/gin-gonic/gin
```

## Usage

1. Run the application:

```sh
go run main.go
```

2. The API will be available at `http://localhost:8080`.

## API Endpoints

### Create a new todo

- **URL**: `/todos`
- **Method**: `POST`
- **Body**:
```json
{
"title": "New Todo"
}
```
- **Success Response**:
- **Code**: `201 CREATED`
- **Content**:
```json
{
"id": 1,
"title": "New Todo",
"completed": false
}
```

### Get all todos

- **URL**: `/todos`
- **Method**: `GET`
- **Success Response**:
- **Code**: `200 OK`
- **Content**:
```json
[
{
"id": 1,
"title": "New Todo",
"completed": false
}
]
```

### Get a todo by ID

- **URL**: `/todos/:id`
- **Method**: `GET`
- **Success Response**:
- **Code**: `200 OK`
- **Content**:
```json
{
"id": 1,
"title": "New Todo",
"completed": false
}
```
- **Error Response**:
- **Code**: `404 NOT FOUND`
- **Content**:
```json
{
"error": "Todo not found"
}
```

### Update a todo

- **URL**: `/todos/:id`
- **Method**: `PUT`
- **Body**:
```json
{
"title": "Updated Todo",
"completed": true
}
```
- **Success Response**:
- **Code**: `200 OK`
- **Content**:
```json
{
"id": 1,
"title": "Updated Todo",
"completed": true
}
```
- **Error Response**:
- **Code**: `404 NOT FOUND`
- **Content**:
```json
{
"error": "Todo not found"
}
```

### Delete a todo

- **URL**: `/todos/:id`
- **Method**: `DELETE`
- **Success Response**:
- **Code**: `204 NO CONTENT`
- **Error Response**:
- **Code**: `404 NOT FOUND`
- **Content**:
```json
{
"error": "Todo not found"
}
```

## Docker

You can also run the application using Docker.

1. Build the Docker image:

```sh
docker build -t todo-api .
```

2. Run the Docker container:

```sh
docker run -p 8080:8080 todo-api
```

## Contributing

If you want to contribute to this project, feel free to submit issues or pull requests. Ashik@spudblocks.com 

## License

This project is licensed under the MIT License.
