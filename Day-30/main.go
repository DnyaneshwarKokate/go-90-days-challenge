package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"day-30/config"
	"day-30/di"
	"day-30/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("==========================================================================")
	fmt.Println("🚀 Day 30: Student Management REST API Milestone Project in Go")
	fmt.Println("==========================================================================")

	cfg := config.LoadConfig()
	cfg.Environment = "development"

	container, err := di.NewContainer(cfg)
	if err != nil {
		fmt.Printf("Failed to initialize DI Container: %v\n", err)
		os.Exit(1)
	}

	router := container.SetupRouter()
	ctx := context.Background()
	container.Logger.Info(ctx, "Student Management REST API started successfully", "port", cfg.Port, "env", cfg.Environment)

	fmt.Println("\n--- 1️⃣ Registering Admin, Teacher, and Student Accounts ---")
	adminToken := registerAndLogin(router, "admin@school.com", "AdminPass123", "ADMIN")
	studentToken := registerAndLogin(router, "john.doe@student.com", "StudentPass123", "STUDENT")

	fmt.Println("\n--- 2️⃣ Creating Student Records (Role: ADMIN) ---")
	s1ID := createStudent(router, adminToken, "John Doe", "john.doe@student.com", "COMPUTER_SCIENCE", 3.85)
	s2ID := createStudent(router, adminToken, "Jane Smith", "jane.smith@student.com", "ELECTRICAL", 3.92)

	fmt.Println("\n--- 3️⃣ Listing All Students (Role: STUDENT) ---")
	listStudents(router, studentToken)

	fmt.Println("\n--- 4️⃣ Updating Student GPA & Status (Role: ADMIN) ---")
	updateStudent(router, adminToken, s1ID, "John Doe", "COMPUTER_SCIENCE", 3.95, "GRADUATED")

	fmt.Println("\n--- 5️⃣ Testing RBAC Authorization (Forbidden Student Deletion by STUDENT Role) ---")
	deleteStudent(router, studentToken, s2ID) // Expected: 403 Forbidden

	fmt.Println("\n--- 6️⃣ Deleting Student Record (Role: ADMIN) ---")
	deleteStudent(router, adminToken, s2ID) // Expected: 200 OK

	fmt.Println("\n✅ Day 30 Milestone REST API Project executed successfully! Check student_app.log for logs.")
}

func registerAndLogin(router *gin.Engine, email, password, role string) string {
	regPayload := domain.RegisterInput{Email: email, Password: password, Role: role}
	regBody, _ := json.Marshal(regPayload)
	req1, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(regBody))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	loginPayload := domain.LoginInput{Email: email, Password: password}
	loginBody, _ := json.Marshal(loginPayload)
	req2, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(loginBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	var res map[string]interface{}
	json.Unmarshal(w2.Body.Bytes(), &res)

	dataMap, ok := res["data"].(map[string]interface{})
	if !ok {
		fmt.Printf("Login Failed for %s: %s\n", email, w2.Body.String())
		return ""
	}

	token, _ := dataMap["token"].(string)
	fmt.Printf("Authenticated %s (%s) -> Token: %s...\n", email, role, token[:20])
	return token
}

func createStudent(router *gin.Engine, token, name, email, dept string, gpa float64) string {
	payload := domain.CreateStudentInput{FullName: name, Email: email, Department: dept, GPA: gpa}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/students", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST /students -> Status %d | Body: %s\n", w.Code, w.Body.String())

	var res map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &res)
	if dataMap, ok := res["data"].(map[string]interface{}); ok {
		id, _ := dataMap["id"].(string)
		return id
	}
	return ""
}

func listStudents(router *gin.Engine, token string) {
	req, _ := http.NewRequest("GET", "/api/v1/students", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("GET /students -> Status %d | Body: %s\n", w.Code, w.Body.String())
}

func updateStudent(router *gin.Engine, token, id, name, dept string, gpa float64, status string) {
	payload := domain.UpdateStudentInput{FullName: name, Department: dept, GPA: gpa, Status: status}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "/api/v1/students/"+id, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("PUT /students/%s -> Status %d | Body: %s\n", id, w.Code, w.Body.String())
}

func deleteStudent(router *gin.Engine, token, id string) {
	req, _ := http.NewRequest("DELETE", "/api/v1/students/"+id, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("DELETE /students/%s -> Status %d | Body: %s\n", id, w.Code, w.Body.String())
}
