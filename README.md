# 🏦 Simple Bank

A backend service for a banking system built with Go, PostgreSQL, and Docker.  
Includes secure authentication, account management, money transfers, and robust concurrency control with deadlock prevention.

[![Go](https://img.shields.io/badge/Go-1.20+-blue)](https://golang.org)  
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue)](https://www.postgresql.org/)  
[![Build](https://img.shields.io/github/actions/workflow/status/abrar-mashuk/simplebank/ci.yml?branch=main)](https://github.com/abrar-mashuk/simplebank/actions)  
[![License](https://img.shields.io/github/license/abrar-mashuk/simplebank)](LICENSE)

---

## 📚 Table of Contents

- [🚀 Features](#-features)
- [🧰 Tech Stack](#-tech-stack)
- [📦 Setup Instructions](#-setup-instructions)
- [🧪 Testing](#-testing)
- [🔁 API Usage Examples](#-api-usage-examples)
- [📬 Postman Collection](#-postman-collection)
- [🎥 Terminal Demo](#-terminal-demo)
- [📝 License](#-license)
- [👤 Author](#-author)

---

## 🚀 Features

- **9+ RESTful API endpoints** for users, accounts, and money transfers  
- **Secure authentication** with JWT and PASETO token standards  
- **Concurrent transaction handling** using goroutines and channels  
- **Deadlock-resilient logic** using transaction ordering and PostgreSQL isolation levels  
- **4+ normalized tables** designed using DBML and type-safe queries generated via `sqlc`  
- **80%+ unit test coverage** with mocked DB dependencies  
- **CI/CD pipeline** using GitHub Actions for automated testing  
- **Docker & Docker Compose** for easy local development  
- **Clean architecture** for scalable, maintainable code  

---

## Tech Stack

- **Language**: Go  
- **Database**: PostgreSQL  
- **Framework**: Gin (HTTP router)  
- **ORM/Query Builder**: sqlc (type-safe SQL)  
- **Auth**: JWT, PASETO  
- **Testing**: Go test, Testify  
- **CI/CD**: GitHub Actions  
- **DevOps**: Docker, Docker Compose  
- **Migration Tool**: golang-migrate  

---

## Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/abrar-mashuk/simplebank.git
cd simplebank

### 2. Create environment file

Copy the example environment file and rename it to `app.env`:

```bash
cp app.env.example app.env

### 3. Run Docker containers (PostgreSQL + API server)

To start your development environment with Docker:

```bash
make compose-up

##  Stopping Containers

To stop all running Docker containers managed by Docker Compose, use the following Make command:

```bash
make compose-down

## 4. Run Database Migrations Manually (Optional)

If you prefer to apply database migrations manually instead of automatically during container startup, you can use the following Make commands:

### Apply Migrations

To run all pending migrations:

```bash
make migrate-up

### Rollback the Most Recent Migration

To undo the most recent database migration, run:

```bash
make migrate-down

## 🧪 Testing

### Run All Unit Tests

To run all unit tests in the project:

```bash
make test

### Run Tests with Code Coverage

To run the tests and view code coverage details:

```bash
go test -cover ./...


## License

This project is licensed under the [MIT License](LICENSE).



If you found this project useful, consider giving it a ⭐ on GitHub!
