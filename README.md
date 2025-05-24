# Go Tasks API

This is a Go API designed to manage tasks. Currently in its early development stages, I'm using this project to learn Go.

---

## Table of Contents

* [Features (Planned/Under Development)](#features-plannedunder-development)
* [Getting Started](#getting-started)
    * [Prerequisites](#prerequisites)
    * [Installation](#installation)
    * [Running the API](#running-the-api)
* [API Endpoints (Planned)](#api-endpoints-planned)

---

## Features (Planned/Under Development)

* **Task Creation:** Allow users to create new tasks with details such as title, description, due date, and priority.
* **Task Retrieval:** Fetch individual tasks or a list of tasks.
* **Task Update:** Modify existing tasks.
* **Task Completion:** Mark tasks as completed.
* **Task Deletion:** Remove tasks.
* **User Authentication (Planned):** Secure API endpoints.
* **Task Filtering & Sorting (Planned):** Enable advanced querying of tasks.

---

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

* Go (version 1.22 or higher) - [https://golang.org/doc/install](https://golang.org/doc/install)

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/migvas/go-tasks-api.git
    cd go-tasks-api
    ```
2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

### Running the API

1.  **Run the application:**
    ```bash
    go run cmd/api/main.go
    ```
    The API should now be running locally, typically on `http://localhost:8080`.

---

## API Endpoints (Planned)

* `POST /tasks`: Create a new task.
* `POST /complete_task`: Mark a task as completed.
* `GET /tasks`: Retrieve all tasks.
* `GET /tasks/{id}`: Retrieve a specific task by ID.
* `PUT /tasks/{id}`: Update an existing task.
* `DELETE /tasks/{id}`: Delete a task.

---