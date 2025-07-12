📦 student-dashboard
A modular Go project for managing students, teachers, classes, and related data — built with clean architecture, gRPC, and protobuf.

🚀 Features


✅ Student service

Register student

List students

Store class & section info




✅ Teacher service

Register teacher (with subjects & class-sections)

Update teacher

Delete teacher

List teachers

Get teacher by ID


✅ Clean architecture

Handler → Service → Repository → Model

Separation of concerns for better maintainability


✅ Proto / gRPC APIs

Protobuf definitions (api/)

Auto-generated Go gRPC code


⚙️ Tech stack

Go

gRPC

Protocol Buffers (proto3)

In-memory storage (using sync.Mutex + map)

Clean architecture style

🛠 How it works (simple flow)
1️⃣ Handler receives gRPC request
2️⃣ Handler converts request into model object
3️⃣ Calls method on Service
4️⃣ Service applies business logic, then calls Repository
5️⃣ Repository saves / retrieves data from in-memory store
6️⃣ Service returns result → Handler returns gRPC response

📦 Build & run
Make sure you have Go installed and protoc + Go plugins set up.

bash
Copy
Edit
# Clone the repo
git clone https://github.com/yourusername/student-dashboard.git
cd student-dashboard

# Generate gRPC code
protoc --go_out=. --go-grpc_out=. api/student/student.proto
protoc --go_out=. --go-grpc_out=. api/teacher/teacher.proto

# Run the project
go run cmd/main.go
(Add real build instructions once you finalize deployment or dockerize it)

✏ Notes
Data is stored in-memory (for demo/testing)

Later you can add real database (MySQL, PostgreSQL, etc.)

Auth service & marks/grades service planned for future

🤝 Contribute / Learn
Feel free to fork, open issues, or suggest improvements!
This is a learning project to practice Go, gRPC, and clean architecture.
.