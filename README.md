# 🧱 Go Starter Project

A reusable and scalable Go boilerplate for building RESTful APIs or web services. This template is designed to help you spin up new projects faster with a clear structure, pre-configured tools, and a clean starting point.

---

## 🚀 Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### 🔧 Prerequisites

- Go 1.21 or higher installed
- MySQL or PostgreSQL database
- Git installed

---

## ⚙️ Setup Guide

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/go-starter.git my-new-app
cd my-new-app

```

---

### 2. Re-initialize Git (optional)

```bash
rm -rf .git
git init
go mod init github.com/yourusername/my-new-app
```

### edit the mod

```bash
go mod edit -module=github.com/your-username/your-new-project
```

### to install required dependencies

```bash
go mod tidy
```

### setup database

from .env.example create .env and add the database credentails

### run migrations

go run migrate.go
