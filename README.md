# 📦 Student Dashboard

A modular Go project to manage students, teachers, classes, and related data — built using clean architecture, gRPC, and protobuf.

---

### 🚀 Features

#### ✅ Student Service
- Register student  
- List students  
- Store class & section info

#### ✅ Teacher Service
- Register teacher (with subjects & class-sections)
- Update teacher
- Delete teacher
- List teachers
- Get teacher by ID

#### ✅ Clean Architecture
- Handler → Service → Repository → Model
- Clear separation of concerns for maintainability

#### ✅ gRPC APIs
- Protobuf definitions in `api/`
- Auto-generated Go gRPC code

---

### 📡 REST / HTTP endpoints (via gRPC Gateway)

| Service   | Method | Endpoint                | Body                                      |
|----------|--------|------------------------|-------------------------------------------|
| Teacher  | POST   | `/v1/teachers`         | `{ "name": "...", "gender": "...", "subjects": [...], "class_sections": [...] }` |
| Teacher  | PUT    | `/v1/teachers/{id}`    | `{ "name": "...", "gender": "...", "subjects": [...], "class_sections": [...] }` |
| Teacher  | DELETE | `/v1/teachers/{id}`    | –                                         |
| Student  | POST   | `/v1/students`         | `{ "name": "...", "age": ..., "roll_number": "...", "class": "...", "section": "..." }` |

Runs on port: **8080**

---

### ⚙️ Tech stack
- Go
- gRPC & Protocol Buffers (proto3)
- In-memory storage (map + sync.Mutex)
- Clean architecture

---

### 🛠 Build & run

Make sure you have Go and `protoc` (with Go plugins) installed.

```bash
# Clone the repo
git clone https://github.com/yourusername/student-dashboard.git
cd student-dashboard

# Generate gRPC code
protoc --go_out=. --go-grpc_out=. api/student/student.proto
protoc --go_out=. --go-grpc_out=. api/teacher/teacher.proto

# Run the server
go run cmd/main.go
