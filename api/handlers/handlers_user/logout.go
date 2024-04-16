package handlers_user

import (
	"database/sql"
	"net/http"

	"github.com/7t1cker/bp/api/utils/sql_tools"
	"github.com/gin-gonic/gin"
)


func Logout(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("Api-Key")
		statusCode, userID := sql_tools.GetUserRequest(db, apiKey)
		if statusCode != http.StatusOK {
			c.JSON(statusCode, gin.H{"error": "Ошибка токена"})
			return
		}
		updateStatusStatement := `UPDATE users SET status = 'не работает' WHERE id = $1`
		_, err := db.Exec(updateStatusStatement, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user status"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
	}
}
