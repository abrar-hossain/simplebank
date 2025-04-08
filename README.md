# 🏦 Simple Bank

A backend service for a banking system built with Go, PostgreSQL, and Docker.  
Includes secure authentication, account management, money transfers, and robust concurrency control with deadlock prevention.

[![Go](https://img.shields.io/badge/Go-1.20+-blue)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue)](https://www.postgresql.org/)
[![Build](https://img.shields.io/github/actions/workflow/status/abrar-mashuk/simplebank/ci.yml?branch=main)](https://github.com/abrar-mashuk/simplebank/actions)
[![License](https://img.shields.io/github/license/abrar-mashuk/simplebank)](LICENSE)

---

## 🚀 Features

- ✅ **9+ RESTful API endpoints** for users, accounts, and money transfers  
- 🔐 **Secure authentication** with JWT and PASETO token standards  
- 🔄 **Concurrent transaction handling** using goroutines and channels  
- 🔒 **Deadlock-resilient logic** using transaction ordering and PostgreSQL isolation levels  
- 🔧 **4+ normalized tables** designed using DBML and type-safe queries generated via `sqlc`  
- 🧪 **80%+ unit test coverage** with mocked DB dependencies  
- ⚙️ **CI/CD pipeline** using GitHub Actions for automated testing  
- 🐳 **Docker & Docker Compose** for easy local development  
- 🧱 **Clean architecture** for scalable, maintainable code  

---

## 🧰 Tech Stack

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

## 📦 Setup Instructions

### 1. Clone the repo

```bash
git clone https://github.com/abrar-mashuk/simplebank.git
cd simplebank
