# Rest API Users Management CRUD Application

The **Rest API Users Management CRUD Application** is a full-stack project that includes:

- A REST API for managing users.
- Integration with AWS SQS for message handling.
- A React-based frontend for managing users.
- A Dockerized setup for running the application locally.

This README provides the requirements, setup instructions, and details for the task.

## Requirements

### 1. Backend API

- Build a REST API for managing users with the following features:
  - CRUD operations:
    - `GET /users`: Retrieve a list of users.
    - `POST /users`: Create a new user.
    - `PUT /users/{id}`: Update an existing user.
    - `DELETE /users/{id}`: Delete a user.
  - User model fields:
    - `id` (UUID)
    - `name` (string)
    - `email` (string)
    - `created_at` (timestamp)
- Use **SOLID principles** to ensure a clean architecture.
- Generate **Swagger/OpenAPI documentation** accessible at `/docs`.

### 2. Messaging

- Add the following endpoints for messaging:
  - `GET /message`: Retrieve messages.
  - `POST /message`: Send a message to **AWS SQS**.
- Messages should include:
  - `id` (UUID)
  - `content` (string)
  - `timestamp` (ISO8601 format).
- Use AWS SDK for SQS integration with environment-based credentials.

### 3. Frontend

- Develop a React application using **Vite** with the following features:
  - A table displaying all users with options to:
    - Edit users.
    - Delete users.
  - A form for adding a new user with basic validation (e.g., email format validation).

### 4. Functional Tests

- Write **functional tests** for the API using the **GIVEN-WHEN-THEN** principle:
  - Example:
    - GIVEN a valid user payload.
    - WHEN the `POST /users` endpoint is called.
    - THEN a new user is created and returned with a 201 status code.
- Use testing frameworks like Jest or Mocha.

### 5. Hosting

- Host the application locally using **Docker**:
  - A container for the backend (Node.js, Python, or a language of your choice).
  - A container for the frontend (React + Vite).
  - Use **Docker Compose** to orchestrate the services.
- Provide environment variables for configuration (e.g., AWS credentials, database URL).
- Host the application remotely using a cloud provider like AWS, GCP, or Azure.
  - Provide a link to the hosted application.
  - Use a CI/CD pipeline for deployment.
  - Use any kind of IaC tool like Terraform, CloudFormation, or Pulumi for deployment.

### 6. Version Control

- Follow the **Conventional Commits** standard for all commits in the repository.

## Deliverables

1. **Source Code**: A Git repository with:
   - Complete backend and frontend code.
   - Functional tests.
2. **Documentation**:
   - A `README.md` file with:
     - Setup instructions.
     - Description of the architecture.
     - How to run the application and tests.
   - Swagger/OpenAPI documentation available at `http://localhost:<PORT>/docs`.
3. **Docker Setup**:
   - A `docker-compose.yml` file to run the application locally.
   - Properly configured backend, frontend, and optional database containers.
4. **Tests**:
   - Functional tests for the backend API using the GIVEN-WHEN-THEN principle.

## Access the application:

Backend API: http://localhost:<BACKEND_PORT>/api
Swagger Docs: http://localhost:<BACKEND_PORT>/docs
Frontend: http://localhost:<FRONTEND_PORT>

## Bonus Features

- Add authentication to the API using **JWT** or **OAuth**.
- Use a database (e.g., PostgreSQL or MongoDB) for persistence, with a container for it in Docker Compose.
- Add real-time updates to the frontend via **WebSockets** or **polling**.
- Configure CI/CD pipelines using GitHub Actions, GitLab CI, or similar tools.

## Development Guidelines

Follow SOLID principles and MVC Layered pattern for clean and maintainable code.

Middleware layer is responsible for handling the requests and responses. Middleware sends the request to the controller. Controller has the constructor with the service layer which handles the business logic. Service layer interacts with the repository layer which is responsible for interacting with the database. Model layer is responsible for defining the schema of the data.

Read the [Coding Principles and Guidelines for clean code](https://principles.wisepace.org/)

### Commit Messages

Follow the Conventional Commits format:

- feat: Introduce a new feature.
- fix: Fix a bug.
- docs: Update documentation.
- style: Code style changes (e.g., formatting).
- refactor: Code restructuring without functional changes.
- test: Add or modify tests.

## Testing Guidelines

Write tests using the GIVEN-WHEN-THEN principle for functional clarity.
