# Task Manager API

A simple RESTful API for managing tasks built with Go, Gin framework, and MySQL.

## Features

- Create new tasks 
- List all tasks
- Store tasks in MySQL database

## Tech Stack

- [Go](https://golang.org/) - Programming language
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [MySQL](https://www.mysql.com/) - Database

## Prerequisites

- Go 1.24 or higher
- MySQL server 
- Environment variables set up for database connection:
  - `DBUSER` - MySQL username
  - `DBPASS` - MySQL password


## Installation

1. Clone the repository:
```bash
git clone <repository-url>
```

2. Change to the project directory:
```bash
cd task-manager-go
```

3. Install dependencies:
```bash
go mod download
```

## Database Setup

1. Create a MySQL database:
```sql
CREATE DATABASE task_manager_go;
```
Create schema from MySQL prompt:
```sql
mysql> use task_manager_go;
mysql> source your/path/to/create-table.sql
```

## Running the Application

1. Set up environment variables:
```bash
set DBUSER=your_mysql_username
set DBPASS=your_mysql_password
```

2. Run the server:
```bash
go run .
```

The server will start at http://localhost:8080

## API Endpoints

- `GET /tasks` - Get all tasks
- `POST /tasks` - Create a new task

### Example Request

Create a new task:
```bash
curl -X POST http://localhost:8080/tasks ^
     -H "Content-Type: application/json" ^
     -d "{\"title\": \"My New Task\"}"
```

## Project Structure

- `main.go` - Entry point and server setup
- `db.go` - Database connection and initialization
- `handlers.go` - HTTP request handlers
- `tasksModel.go` - Task data model
- `create-table.sql` - Database schema

## License

MIT License