# Vaccine Service

This is the Vaccine Service component of the Vaccination Portal microservices architecture. It connects to a database MySQL Database . Follow the steps below to configure and run the service locally.

## Prerequisites

- Go (1.22+ recommended)
- A running database (MySQL)

## Setup Instructions

### 1. Create `.env` File

Create a `.env` file in the root of the project directory and add the following environment variables according to your local or remote configuration:

```env
DB_HOST=
DB_PORT=
DB_USER=
DB_PASS=
DB_NAME=
```
Save the file after entering all values.

### 2. Install Go Dependencies
Run the following command to install all required dependencies:
```
go mod tidy
```
### 3. Run the Vaccine Service
To start the service, run the following command:
```
go run cmd/main.go
```
This will execute the application on your system.

