# Short URL Service

This project is a URL shortening service built with Go, Redis and the Gin web framework. It allows users to create short URLs and redirect to the original URLs.

## Features

- Create short URLs
- Redirect to original URLs
- Redis store

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/mathleite/go-shorturl.git
cd go-shorturl
```

### Build and Run
To build and run the project using Docker Compose, execute the following command:

```sh
docker-compose up --build
```
This will start the application on `http://localhost:8080`.

### Usage
1. Create a Short URL:
- - The client sends a POST request to /api/v1/ with the original URL.
- - The server generates a unique identifier and stores the original URL in Redis with the identifier as the key.
- - The server responds with the short URL.

2. Redirect to Original URL:
- - The client sends a GET request to /api/v1/:identifier.
- - The server retrieves the original URL from Redis using the identifier.
- - The server redirects the client to the original URL.

### API Endpoints
#### Create Shor URL
Request
```
POST /api/v1/
Content-Type: application/json

{
  "original_url": "https://www.example.com"
}
```
Response
```json
{
  "short_url": "http://localhost:8080/api/v1/abc123"
}
```

#### Redirect to Original URL
Request
```
GET /api/v1/:identifier
```
Response
- Redirects to the original URL.

## Redis
>In this project, Redis is used as a key-value store to manage the mapping between short URLs and their corresponding original URLs.

#### Connecting to Redis
In the Go application, a Redis client is created and connected to the Redis server. This is typically done in a separate package or file, such as redis.go:
```go
package redis

import (
    "context"
    "os"

    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var client *redis.Client

func CreateInstance() {
    if client == nil {
        client = redis.NewClient(&redis.Options{
            Addr:     os.Getenv("REDIS_HOST"),
            Password: "",
            DB:       0,
        })
    }
}

func GetClient() *redis.Client {
    return client
}
```

#### Storing and Retrieving Data
The Redis client is used to store and retrieve data. For example, when creating a short URL, the original URL is stored in Redis with the short URL identifier as the key:
```go
package handler

import (
    "net/http"
    "mathleite/short-url/redis"

    "github.com/gin-gonic/gin"
)

type RedisDataInput struct {
    Identifier   string
    BaseUrl      string
    OriginalUrl  string
    Ttl          time.Duration
}

func CreateShortUrl(c *gin.Context) {
    var input struct {
        OriginalUrl string `json:"original_url"`
    }

    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    identifier := generateIdentifier() // Function to generate a unique identifier
    data := RedisDataInput{
        Identifier:   identifier,
        BaseUrl:      "http://localhost:8080/api/v1/",
        OriginalUrl:  input.OriginalUrl,
        Ttl:          24 * time.Hour, // Example TTL of 24 hours
    }

    redis.Store(data)

    c.JSON(http.StatusOK, gin.H{
        "short_url": data.BaseUrl + data.Identifier,
    })
}

func RedirectOriginalUrl(c *gin.Context) {
    identifier := c.Param("identifier")
    client := redis.GetClient()

    originalUrl, err := client.Get(ctx, identifier).Result()
    if err == redis.Nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Redirect(http.StatusMovedPermanently, originalUrl)
}
```
