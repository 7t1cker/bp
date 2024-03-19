package handlers_user

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/7t1cker/bp/api/utils/sql_tools"
	"github.com/gin-gonic/gin"
)

type User struct {
    ID          int              `json:"id,omitempty"`
    FirstName   string           `json:"first_name"`
    LastName    string           `json:"last_name"`
    MiddleName  string           `json:"middle_name"`
    DivisionID  int64            `json:"division_id"`
    GroupID     int64            `json:"group_id"`
    SkillTasks  json.RawMessage           `json:"skill_tasks"`
    Login       string           `json:"login"`
    Role        string           `json:"role"`
    Bl float64 `json:"balance"`
}

func KL(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        apiKey := c.Request.Header.Get("Api-Key")

        statusCode, userID := sql_tools.GetUserRequest(db, apiKey)
        if statusCode != http.StatusOK {
            c.JSON(statusCode, gin.H{"error": "Ошибка токена"})
            return
        }
        var user User
        err := db.QueryRow("SELECT id, first_name, last_name, middle_name, division_id, group_id, skill_tasks, login, role, balance  FROM users WHERE id = $1", userID).Scan(
            &user.ID,
            &user.FirstName,
            &user.LastName,
            &user.MiddleName,
            &user.DivisionID,
            &user.GroupID,
            &user.SkillTasks,
            &user.Login,
            &user.Role,
            &user.Bl,
        )
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"data": user})
    }
}
