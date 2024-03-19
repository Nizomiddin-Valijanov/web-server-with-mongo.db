# Web Server with MongoDB

This project implements a web server using Golang and MongoDB.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Setup](#setup)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

## Features

- **HTTP Server**: Implements a basic HTTP server using Golang's `net/http` package.
- **MongoDB Integration**: Connects to MongoDB database to perform CRUD operations.
- **RESTful API**: Provides RESTful endpoints to interact with MongoDB data.
- **CORS Support**: Implements Cross-Origin Resource Sharing (CORS) for handling requests from web clients.

## Requirements

- Golang installed on your system. You can download it [here](https://golang.org/dl/).
- MongoDB installed and running locally or accessible via URI.

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/web-server-mongodb.git
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Set up MongoDB:
   
   Ensure MongoDB is running locally or update the connection URI in the code.

4. Build the project:

    ```bash
    go build
    ```

## Usage

1. Run the server:

    ```bash
    ./web-server-mongodb
    ```

2. Access the server endpoints through `http://localhost:8080` or your configured port.

## Endpoints

- **GET /people**: Retrieve all people records.
- **POST /people**: Create a new person record.
- **GET /heath**: Health check endpoint.
- **GET /javascript-response**: Sample JavaScript response endpoint.

## Contributing

Contributions are welcome! If you have any ideas, improvements, or bug fixes, feel free to open issues or pull requests.

## License

This project is licensed under the [MIT License](LICENSE).
