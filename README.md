# Task Reminder API

A RESTful API built with Go Fiber and MongoDB for managing todo items.

## Features

- Create, read, update, and delete todo items
- Mark todo items as completed
- Store todos persistently in MongoDB
- RESTful API design

## Tech Stack

- Go 1.24+
- Fiber v2
- MongoDB
- Air

## Prerequisites

- Go 1.24 or higher
- MongoDB (local installation or MongoDB Atlas)
- Git

## Getting Started

### Clone the repository

```bash
git clone https://github.com/ranieeery/task-reminder-go.git
cd task-reminder-go
```

### Environment Setup

1. Create a `.env` file in the root directory using the provided `.env.example` as a template:

```bash
cp .env.example .env
```

2. Update the MongoDB connection string in the `.env` file

```.env
PORT=4000
MONGODB_URI=your_mongodb_connection_string
```

### Install dependencies

```bash
go mod download
```

### Running the application

```bash
air
```

The API will be available at `http://localhost:4000` (or the port you specified in `.env`).

## API Endpoints

### Get all todos

``` http
GET /api/todos
```

### Create a new todo

```http
POST /api/todos
```

Request body:

```json
{
  "body": "Task description"
}
```

### Mark a todo as completed

```http
PATCH /api/todos/:id
```

### Update todo content

```http
PUT /api/todos/:id
```

Request body:

```json
{
  "body": "Updated task description"
}
```

### Delete a todo

```http
DELETE /api/todos/:id
```

## Project Structure

```structure
task-reminder-go/
├── .air.toml          # Air configuration for live reload
├── .env               # Environment variables (not in repo)
├── .env.example       # Example environment variables
├── .gitignore         # Git ignore rules
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
├── main.go            # Main application code
└── README.md          # This file
```
