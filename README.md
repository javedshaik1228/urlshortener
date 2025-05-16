# URL Shortener Service

A microservices-based URL shortener application built with Go, gRPC, and MongoDB. This project exposes REST API endpoints to shorten URLs and retrieve the original URLs.

## Features
- Shorten long URLs into short, unique URLs.
- Retrieve the original URL from a shortened URL.
- Built with a scalable microservices architecture.
- Uses MongoDB as the NoSQL database for storing URL mappings.

## Setup Instructions

### Prerequisites
- Go 1.18+ installed
- MongoDB Atlas or a local MongoDB instance

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/urlshortener.git
   cd urlshortener
2. Create a .env file in the root directory and configure the following environment variables:
    ```bash
    # Application ports
    GATEWAY_SERVER_ADDR=<gateway_server_port>
    SHORTEN_SERVER_ADDR=<shorten_server_port>
    RETRIEVE_SERVER_ADDR=<retrieve_server_port>

    # NoSQL MongoDB config
    DB_USER=<db_user>
    DB_PASSWORD=<db_pwd>
    HOST=<db_host>
    DATABASE_NAME=<db_name>
    COLLECTION_NAME=<db_url_collection_name>
3. Install dependencies:
    ```bash
    go mod tidy
4. Generate Protobuf code
    ```bash
    mkdir proto/genproto/retrievepb
    mkdir proto/genproto/shortenpb
    make gen
5. Run the services in different terminals:
    ```bash
    - make run-gateway
    - make run-shortener
    - make run-retriever
6. Test the application using postman or `curl`

## API Endpoints
### Gateway service
- Shorten URL: `POST /shorten`
    - Request Body: `{ "longUrl" : "https://example.com" }`
    - Response: `{ "shortUrl" : "abc123" }`

- Retrieve URL: `GET /retrieve/:shortUrl`
    - Response: `{ "longUrl" : "https://example.com" }`

## Technologies Used
- **Go**: Programming language for building the services.
- **Rest** Communication protocol between client and gateway server
- **gRPC**: Communication protocol between microservices.
- **MongoDB**: NoSQL database for storing URL mappings.
- **Gin**: Web framework for the Gateway service.
