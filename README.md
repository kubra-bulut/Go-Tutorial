# Go-Tutorial 🚀

This repository contains a Go project built by following the tutorial [Build a REST API with Go](https://www.youtube.com/watch?v=lNd7XlXwlho&ab_channel=freeCodeCamp.org) from freeCodeCamp.org. The project demonstrates how to build a REST API using Go and integrates with PostgreSQL for database operations.

## Project Overview 📚

The project implements a simple REST API for managing a list of todos. It includes functionality to create, read, update, and delete todos. The API is built using the Go programming language with the Fiber web framework and PostgreSQL for data storage.

## Technologies Used 🛠️

- **Go (Golang)**: A statically typed, compiled language designed for simplicity and efficiency.
- **Fiber**: A fast and lightweight web framework for Go, inspired by Express.js.
- **PostgreSQL**: An open-source relational database system used for data storage.
- **Vite**: A build tool that provides a fast development environment for modern web projects.
- **npm**: A package manager for JavaScript used to manage frontend dependencies.

## Installation ⚙️

To set up and run the project, follow these steps:

1. **Install Go dependencies:**
   
   ``` go mod tidy ```

2. **Set up PostgreSQL:**
  - Ensure PostgreSQL is installed and running on your machine.

  - Create a database and a table named todos with the following schema:

  ```
  CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    body TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE
  );
  ```

3. **Update the database connection string:**

    Edit `main.go` and set the correct PostgreSQL connection string in the `connStr` variable.

## Running the Project 🚀
- **Run the Go server:**
   
   ```go run main.go```

- **Start the development server:**
   
    `npm run dev`

## Screenshots 🖼️

Here’s a screenshot of the frontend in action:

![image](https://github.com/user-attachments/assets/4541e2da-e790-46d4-8d43-21ad625c2803)
