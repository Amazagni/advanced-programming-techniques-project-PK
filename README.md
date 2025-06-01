# Inventory Management App

A full-stack inventory management application with a Go backend and a modern JavaScript frontend. The app allows users to manage items, track stock, and maintain organized inventories with authentication and user-friendly UI components.

## 🚀 Getting Started

### Backend

#### Prerequisites

- Go 1.20+

#### Run the backend server:

```bash
cd backend
go run cmd/main.go
```

### Frontend

#### Prerequisites

- Node.js 18+
- npm or yarn

#### Install dependencies and run:

```bash
cd frontend\web-app
npm install
npm run dev
```

---

## 📦 Features
🔐 Authentication with route guards

📦 Inventory management with CRUD operations

📊 Dashboard views and item details

🧱 Component-based frontend architecture

📁 Structured backend with separation of concerns

🧪 Vendor-managed dependencies (Go)

---

## 🗂 Project Structure

### Backend (Go)

backend/

├── assets/ # Static assets (e.g., images)

├── bin/ # Compiled binaries or build scripts

├── cmd/ # Application entry points (main package)

├── docs/ # Project documentation

├── handlers/ # HTTP handlers and route logic

├── logger/ # Logging utilities

├── models/ # Database models and data structures

├── vendor/ # Dependency packages (vendored using go mod vendor)

### Frontend (JavaScript/TypeScript)

frontend/src/
├── app/

│ ├── components/ # Reusable UI components

│ │ ├── item-card/

│ │ └── navbar/

│ ├── guards/ # Route guards for authentication

│ ├── models/ # Frontend data models

│ ├── services/ # API and business logic services

│ └── views/ # View pages for routing

│ ├── add-item-view/

│ ├── home-view/

│ ├── inventory-view/

│ └── login-view/

├── assets/ # Static assets like images

---

## 📄 Documentation

- API docs can be found in backend/docs

- Swagger/OpenAPI integration via go-openapi and swaggo

---

## 🧰 Built With
- Backend: Go, Gorilla Mux, Swagger, EasyJSON

- Frontend: (Specify if Angular, React, etc.)

- Database: (Add if applicable)

- Vendor Tools: Testify, Swag, EasyJSON
