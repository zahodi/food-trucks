# Food Truck API

This Go project serves an API that provides information about food trucks in San Francisco. The data is fetched from a public CSV file hosted by the city of San Francisco. The API allows users to retrieve all food trucks and search for food trucks based on food items.

## Features

- Fetch data from a remote CSV file.
- Serve data as JSON over HTTP.
- Search food trucks by food items.
- Ensure unique results based on food truck names.

## Endpoints

- **GET /foodtrucks**: Returns a list of all food trucks.
- **GET /foodtrucks/search?food=<query>**: Searches for food trucks by food items. The query parameter `food` is required.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or later)
- [Docker](https://www.docker.com/get-started) (optional, for containerization)

## Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/foodtruck-api.git
cd foodtruck-api
```

## Running the Application Locally

# Run the application:

```sh

go run main.go
```


Access the API:

List all food trucks: http://localhost:8000/foodtrucks
Search food trucks by food type: http://localhost:8000/foodtrucks/search?food=burger
