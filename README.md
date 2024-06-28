# Food Truck API

This Go project serves an API that provides information about food trucks in San Francisco. The data is fetched from a public CSV file hosted by the city of San Francisco. The API allows users to retrieve all food trucks and search for food trucks based on food items.

## Features

- Fetch data from a remote CSV file. Local copy is provider in the repo as well.
- Serve data as JSON over HTTP.
- Search food trucks by food items.
- Ensure unique results based on food truck names.

## Endpoints

- **GET /foodtrucks**: Returns a list of all food trucks.
- **GET /foodtrucks/search?food=<query>**: Searches for food trucks by food items. The query parameter `food` is required.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.22 or later)
- [Docker](https://www.docker.com/get-started) (optional, for containerization)

## Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/foodtrucks-api.git
cd foodtrucks-api
```

## Running the Application Locally

1. Run the application:

```sh

go run main.go
```


2. Access the API:

List all food trucks: http://localhost:8000/foodtrucks

Search food trucks by food type: http://localhost:8000/foodtrucks/search?food=burger

## Running the Application with Docker

1. Build the Docker image:

```sh

docker build -t foodtrucks-api .
```

2. Run the Docker container:

```sh

docker run -p 8000:8000 foodtrucks-api
```

## Testing

1. Run the tests:

```sh
go test -v
```