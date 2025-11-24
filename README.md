# Collaborative Task Manager

A **Go** based collaborative task‑manager API built with **Gin**, **GORM**, **PostgreSQL**, and **JWT** authentication.  It supports role‑based access for **Managers** and **Developers**, provides Swagger UI for API exploration, and is ready for containerised deployment.

---

## 📦 Project Structure

```
.
├── cmd/                # Entry points (main.go)
├── internal/           # Application code (controllers, services, routes)
├── pkg/                # Shared packages (db, models, utils)
├── docs/               # Additional documentation
├── .env                # Environment variables (ignored by git)
├── go.mod / go.sum     # Go module definition
├── API_DOCUMENTATION.md
├── TEST_RESULTS.md
└── README.md           # <‑‑ you are here!
```

---

## 🚀 Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/Swarnadip-Dey/Collaborative-taskmanager.git
   cd Collaborative-taskmanager
   ```

2. **Install dependencies** (if you haven't already)
   ```bash
   go mod tidy
   ```

3. **Configure the database**
   - Create a PostgreSQL database.
     ```dotenv
     DATABASE_URL=postgres://user:password@localhost:5432/taskdb?sslmode=require
     ```
   - Remember to add *PORT* and *JWT_SECRET* in env

4. **Run the API**
   ```bash
   go run ./cmd/api
   # or, if you prefer the single‑file form:
   go run ./cmd/api/main.go
   ```
   The server starts on **port 8080** and Swagger UI is available at `http://localhost:8080/swagger/index.html`.

---

## 🔐 Authentication & RBAC

- **Register** – `POST /api/register`
- **Login** – `POST /api/login` (returns a JWT)
- **Profile** – `GET /api/profile` (requires JWT)
- **Manager routes** – `/api/manager/*` (requires `role=manager`)
- **Developer routes** – `/api/dev/*` (requires `role=dev`)
- **Admin routes** – `/api/admin/*` (requires `role=admin`)

---

## 📚 API Endpoints (excerpt)

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/api/register` | Create a new user |
| `POST` | `/api/login` | Authenticate and receive JWT |
| `GET`  | `/api/ping` | Health check |
| `POST` | `/api/manager/workspaces` | Create a workspace (manager) |
| `GET`  | `/api/manager/workspaces/:workspace_id/projects` | List projects in a workspace |
| `POST` | `/api/manager/projects` | Create a project |
| `PUT`  | `/api/manager/tasks/:id/assign` | Assign a task |
| `GET`  | `/api/dev/projects/:id` | Get project details (developer) |
| `POST` | `/api/dev/tasks` | Create a task |
| `PUT`  | `/api/dev/tasks/:id` | Update a task |
| `GET`  | `/api/admin/users` | List all users (admin) |

---

## 🧪 Testing

A helper script `test_api.sh` is provided.  Run it with:
```bash
bash test_api.sh
```
Check `TEST_RESULTS.md` for the latest test run output.

---

## 📦 Building a Docker Image (optional)

```Dockerfile
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o /app/api ./cmd/api/main.go

FROM alpine:latest
COPY --from=builder /app/api /usr/local/bin/api
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/api"]
```
Build and run:
```bash
docker build -t taskmanager .
docker run -p 8080:8080 taskmanager
```

---

## 📄 License

This project is licensed under the **MIT License**.

---

*Happy coding!*
