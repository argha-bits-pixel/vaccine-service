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
### 3. Create Database Schema
Before running the service, make sure the following table exists in your database. You can create it using the SQL command below:
```
CREATE TABLE vaccination_drives (
  id INT AUTO_INCREMENT PRIMARY KEY,
  vaccine_name VARCHAR(100) NOT NULL,
  drive_date DATE NOT NULL,
  doses INT NOT NULL,
  classes TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
);
```
### 4. Run the Vaccine Service
To start the service, run the following command:
```
go run cmd/main.go
```
This will execute the application on your system.


