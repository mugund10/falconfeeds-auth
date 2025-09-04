# FalconFeeds User Authentication Service

This project is my submission for the **Backend/Full-stack Engineering Assignment** at FalconFeeds.io.  
It implements a **User Authentication API** with support for **user signup, login, and JWT-based authentication**, built using **Go** and **MongoDB**.

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
- [ ] **Integration Tests** (basic test scaffolding only)

---

## Installing mongodb

### Installing Mongodb with Docker

```
docker run --name mymongo -p 27017:27017 -d mongo
```
