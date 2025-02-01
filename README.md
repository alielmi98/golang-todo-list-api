# golang-todo-list-api
 This project implements a RESTful API for managing a to-do list with user authentication, using Go. It includes endpoints for user registration, login, and CRUD operations for to-do items, with features like pagination, filtering, and error handling to ensure a secure and user-friendly experience.
This project is part of the Todo List API project on [roadmap.sh](https://roadmap.sh/projects/todo-list-api).

## Features

- User registration and login
- CRUD operations for to-do items
- User authentication and authorization
- refresh token mechanism for the authentication
- Pagination, filtering and sorting for to-do items
- Error handling and security measures
- rate limiting and throttling for the API
- Database integration (using Docker)
## Endpoints

The API provides the following endpoints:

1. **User Registration**:
- `POST /register`
- Request: `{ "name": "John Doe", "email": "john@doe.com", "password": "password" }`
- Response: `{ "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" }`

2. **User Login**:
- `POST /login`
- Request: `{ "email": "john@doe.com", "password": "password" }`
- Response: `{ "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" }`

3. **Create a To-Do Item**:
- `POST /todos`
- Request: `{ "completed": "true", "title": "Buy groceries", "description": "Buy milk, eggs, and bread" }`
- Response: `{ "id": 1, "completed": "true", "title": "Buy groceries", "description": "Buy milk, eggs, and bread" }`

4. **Update a To-Do Item**:
- `PUT /todos/1`
- Request: `{ "completed": "true" "title": "Buy groceries", "description": "Buy milk, eggs, bread, and cheese" }`
- Response: `{ "id": 1,"completed": "true", "title": "Buy groceries", "description": "Buy milk, eggs, bread, and cheese" }`

5. **Delete a To-Do Item**:
- `DELETE /todos/1`
- Response: 204 No Content

6. **Get To-Do Items**:
- `GET /todos?page=1&limit=10`
- Response: `{ "data": [...], "page": 1, "limit": 10, "total": 2 }`

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/alielmi98/golang-todo-list-api.git
```

2. Navigate to the project directory:

```bash
cd golang-todo-list-api
```
3. Run the application using Docker:

```bash
docker-compose up -d
```

5. This will build and start the Docker containers for the application and the database.
The API will be available at http://localhost:8080.

## Contributing

Contributions are welcome! If you find any issues or have ideas for improvements, please feel free to create a new issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).