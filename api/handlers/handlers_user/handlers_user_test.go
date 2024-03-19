package handlers_user

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func TestCreateUser(t *testing.T) {
    db := prepareTestDB()
    defer db.Close()
    router := gin.Default()
    router.POST("/create-user", CreateUser(db))
    requestBody := `{
        "first_name": "Олег",
        "last_name": "Такой",
        "middle_name": "Олег",
        "division_id": 1,
        "group_id": 1,
        "skill_tasks": [1, 2, 3],
        "login": "loginttp",
        "password": "password12ц3",
        "role": "менеджер"
    }`
    req, err := http.NewRequest("POST", "/create-user", strings.NewReader(requestBody))
    if err != nil {
        t.Fatalf("failed to create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)
    if resp.Code != http.StatusCreated {
        t.Errorf("expected status %d, got %d", http.StatusCreated, resp.Code)
    }
    expectedResponseBody := `{"message":"User created successfully","user_id":1}`
    if resp.Body.String() != expectedResponseBody {
        t.Errorf("expected response body %s, got %s", expectedResponseBody, resp.Body.String())
    }
}

func prepareTestDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:12041982@localhost:5432/vk_v2?sslmode=disable")
    if err != nil {
        panic("failed to connect to test database: " + err.Error())
    }
    return db
}
