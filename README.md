# 💰 SimpleBank – Backend System in Go

SimpleBank is a minimal core banking system implemented in Go. It includes secure authentication, transaction-safe logic, and clean architecture principles. The project simulates real-world backend development using modern tools and practices.

## 🧠 Why This Project?

This backend system was built as a hands-on learning project to explore:

- ACID-compliant transactions
- Secure user authentication with JWT and PASETO
- Clean architecture with separation of concerns
- API testing with high coverage
- CI/CD with GitHub Actions
- Docker-based development environments

---

## 📚 Table of Contents

- [🚀 Features](#-features)
- [🧰 Tech Stack](#-tech-stack)
- [📦 Setup Instructions](#-setup-instructions)
- [🧪 Testing](#-testing)
- [📝 License](#-license)
- [👤 Author](#-author)

---

## 🚀 Features

## 🚀 Features

- **9+ RESTful API endpoints** for users, accounts, and money transfers  
- **Secure authentication** with JWT and PASETO token standards  
- **Concurrent transaction handling** using goroutines and channels  
- **Deadlock-resilient logic** using transaction ordering and PostgreSQL isolation levels  
- **4+ normalised tables** designed using DBML and type-safe queries generated via `sqlc`  
- **80%+ unit test coverage** with mocked DB dependencies  
- **CI/CD pipeline** using GitHub Actions for automated testing  
- **Docker & Docker Compose** for easy local development  
- Built a minimal Docker image using multi-stage builds and Go binary stripping, reducing image size by **97% (from 877MB to 25.1MB)**; orchestrated services with Docker Compose for consistent local environments  
- **Performance benchmark:** Sustained **200+ secure money transfers/sec** under **100 concurrent users**, with an average latency of **400ms** and **100% success rate** in full ACID-compliant operations  

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
```
### 2. Create an environment file

Copy the example environment file and rename it to `app.env`:

```bash
cp app.env.example app.env
```
### 3. Run Docker containers (PostgreSQL + API server)

To start your development environment with Docker:

```bash
make compose-up
```
##  Stopping Containers

To stop all running Docker containers managed by Docker Compose, use the following Make command:

```bash
make compose-down
```
## 4. Run Database Migrations Manually (Optional)

If you prefer to apply database migrations manually instead of automatically during container startup, you can use the following Make commands:

### Apply Migrations

To run all pending migrations:

```bash
make migrate-up
```
### Rollback the Most Recent Migration

To undo the most recent database migration, run:

```bash
make migrate-down
```
## 🧪 Testing

### Run All Unit Tests

To run all unit tests in the project:

```bash
make test

### Run Tests with Code Coverage

To run the tests and view code coverage details:

```bash
go test -cover ./...
```
---

## 🐳 Docker Image Optimisation

This project uses a multi-stage Docker build to generate a clean, minimal, production-ready container image.

### 🔍 Image Size Comparison

| Variant                 | Size    | Description                             |
|------------------------|---------|-----------------------------------------|
| simplebank-unoptimized | 877MB   | Single-stage build with full toolchain  |
| simplebank-optimized   | 25.1MB  | ✅ Multi-stage build, stripped Go binary |

✅ Reduced Docker image size by **~97%** using:
- Multi-stage builds
- Stripped Go binaries (`-ldflags="-s -w"`)
- Minimal base image (`alpine:3.19`)


### 🛠 Build Commands

```bash
# Build optimised image
make docker-build-optimized

# View image sizes
make docker-size
```

### 🧪 Run Optimised Container

```bash
docker run --rm -p 8080:8080 simplebank-optimized
```


## License

This project is licensed under the [MIT License](LICENSE).



If you found this project useful, consider giving it a ⭐ on GitHub!
