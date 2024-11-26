# Web Scraper API

This project is a web scraper API built with Go. It scrapes articles from various websites and stores them in a PostgreSQL database. It also logs requests and errors to a MongoDB database.

## Table of Contents

- [Project Structure](#project-structure)
- [Technologies](#technologies)
- [Installation](#installation)
- [Details](#details)
- [Endpoints](#endpoints)
- [Points of Improvement](#points-of-improvement)

## Project Structure

```
web-scraper-api
    │   .env
    │   go.mod
    │   go.sum
    │   main.go
    │
    ├───db
    │       db.go
    │       mongo.go
    │
    ├───middlewares
    │       logger.go
    │       ratelimiter.go
    │
    ├───models
    │       article.go
    │
    ├───routes
    │       articles.go
    │       routes.go
    │
    ├───scraper
    │       scrapehtml.go
    │       scraperss.go
    │
    └───utils
            logging.go
```

## Technologies

- Go
- Gin
- Goquery
- PostgreSQL
- MongoDB

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/MohamedSaidCS/web-scraper-api.git
   cd web-scraper-api
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Set up the environment variables in the `.env` file:

   ```env
   # DB
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=admin
   DB_NAME=go

   # MONGO
   MONGO_HOST=localhost
   MONGO_PORT=27017
   MONGO_DB=go
   ```

4. Run the application:

   ```sh
   go run .
   ```

## Details

A simple server that uses the Gin framework, and has a cron job that scrapes websites either using their RSS feeds or by parsing the HTML using Goquery. The scraped articles are stored in a PostgreSQL database.

The server has a single rate limited endpoint that retrieves articles with pagination, while logging requests and errors to a MongoDB database using a middleware.

The server runs on `http://localhost:8080`.

## Endpoints

- `GET /articles`: Retrieve articles with pagination.

#### Query Parameters

- `page`: The page number to retrieve.
- `per_page`: The number of articles per page (default is 5).

#### Example Response

```json
{
  "data": [
    {
      "id": 1,
      "title": "Article Title",
      "link": "https://example.com/",
      "timestamp": "2024-11-28T12:00:00Z"
    },
    {
      "id": 2,
      "title": "Article Title",
      "link": "https://example.com/",
      "timestamp": "2024-11-28T12:00:00Z"
    },
    ...
  ],
  "paging": {
    "page": 1,
    "per_page": 10,
    "pages": 5,
    "total": 50
  }
}
```

## Points of Improvement

- Implement better pagination logic.
- Make rate limitting user specific by using a cache like Redis.
- Use concurrency more in-depth.
