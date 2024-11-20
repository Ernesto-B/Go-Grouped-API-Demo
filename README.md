# GIN REST API With Middleware and Sorting Query

## Overview
The purpose of this project is to demonstrate how to build a REST API with the Gin framework in Go, using groups of routes, middleware for logging, and a sorting query for sorting responses.

## Project Structure
```plaintext
.
├── handlers
│   ├── albumsEndpoints.go    # Contains endpoints for managing albums
│   ├── nonAlbumEndpoints.go  # Contains additional endpoints (e.g., homepage)
├── middleware
│   ├── fileLogging.go        # Middleware for logging to a file
│   ├── logging.go            # Middleware for logging to the console
├── models
│   ├── data.go               # Album data model and sample album data
├── .gitignore                # Git ignore file
├── go.mod                    # Go module file
├── go.sum                    # Dependency lock file
├── README.md                 # Project documentation
├── requests.rest             # Sample HTTP requests for testing
├── routes.go                 # Route group definitions
├── server.go                 # Main server file
└── server.log                # Log file created by the file logging middleware
```

## Features
- **Album Management**: Provides endpoints for managing albums (get all albums, get album by ID, add album, update album, delete album).
- **Sorting Query**: Sorts albums by `id`, `title`, or `artist`.
- **Grouping of Routes**: Routes are grouped into `albumsEndpoint` and `nonAlbumEndpoint`.
- **Logging Middleware**: Logs details of requests and responses to the console.
- **File Logging Middleware**: Logs details of requests and responses to a file (`server.log`).

## Endpoints
| Method | Endpoint | Description |
| --- | --- | --- |
| GET | `/albums` | Get all albums, can also query |
| POST | `/albums` | Add album |
| GET | `/albums/:id` | Get album by ID |
| GET | `/other` | Homepage (displays simple text) |

## Usage
1. Clone the repository:
   ```bash
   git clone https://github.com/Ernesto-B/Go-Grouped-API-Demo.git
   cd Go-Grouped-API-Demo
    ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the server:
   ```bash
    go run .
    ```
4. Test the endpoints using the provided `requests.rest` file if you have the Thunder Client extension in VSCode, or use other tools like Postman.