# 🚀 CI/CD Pipeline Demo with Notely

A demonstration of building a **complete CI/CD pipeline** for a Go web application using GitHub Actions, with **Notely** (a simple note-taking app) as the working example.

---

## 📌 About

This project showcases how code moves from **development → testing → build → deployment** using automated workflows.
Notely is included to demonstrate a real, deployable backend system.

---

## ✨ Features

* ⚙️ **CI/CD Pipeline (GitHub Actions)**

  * Automated testing & build (CI)
  * Docker-based deployment workflow (CD)

* 📝 **Notely App**

  * User creation with API key auth
  * Create & fetch notes
  * REST API with middleware
  * Embedded static frontend

---

## 🏗️ Architecture

```plaintext
Client → Go Server (chi)
            |
     -------------------
     |                 |
 Static           REST API (/v1)
                     |
                Auth Middleware
                     |
                  Handlers
                     |
                 Database (sqlc)
```

---

## 🛠️ Tech Stack

* Go (1.22+), chi
* PostgreSQL (optional)
* sqlc
* Docker
* GitHub Actions

---

## 🚀 Run Locally

```bash
go build -o notely
./notely
```

Visit: [http://localhost:8080](http://localhost:8080)

---

## 🔐 Example API

```http
POST /v1/users
```

```http
Authorization: ApiKey <token>
```

```http
POST /v1/notes
GET  /v1/notes
```

---

## ⚙️ CI/CD Workflow

* **CI (`ci.yml`)**

  * Runs tests & build on PRs/push

* **CD (`cd.yml`)**

  * Builds Docker image on merge to main
  * Ready for deployment (extensible)

---

## 📂 Structure

```
main.go
handlers/
internal/
static/
Dockerfile
.github/workflows/
```

---

## 🧾 Summary

A minimal Go web app used to demonstrate **real-world CI/CD practices**, including automation, containerization, and clean backend design.
