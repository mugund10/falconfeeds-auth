# FalconFeeds User Authentication Service

This project is my submission for the **Backend/Full-stack Engineering Assignment** at FalconFeeds.io.

---

## Implementation Status

### Core Requirements

- [x] **POST /signup** (with validation, duplicate email check, DB insert)
- [x] **POST /login** (with password hashing + verification)
- [x] **JWT session management** (ID + email claims)
- [x] **NoSQL Database** (MongoDB used)
- [x] **Proper error handling** (400, 401, 409 responses)

### Bonus Implementations

- [x] **Containerization (Docker + Buildx)**
- [x] **Code Structure** (modular packages: `api/`, `storage/`, `types/`)
- [x] **Password Security (bcrypt)**
- [x] **Abuse Prevention - ratelimits**
- [x] **Swagger/OpenAPI Docs**
- [ ] **Integration Tests**

---

## Installation & Setup

### 1. Installing Mongodb with Docker

```
docker run --name mymongo -p 27017:27017 -d mongo
```

### 2. Run the FalconFeeds Auth Service

```
docker run -d   --name falconfeeds-auth   -p 8080:8080   -e DB_ADDR="mongodb://mymongo:27017"   -e DB_NAME="ff"   -e JWT_SECRET="mugu2000ispass"   --link mymongo:mymongo   mugundhan10/falconfeeds:latest
```

## Documentation

Live Swagger API Docs: https://ff.of.mugund10.dev/swagger/index.html

## Live Deployment

Base URL: https://ff.of.mugund10.dev