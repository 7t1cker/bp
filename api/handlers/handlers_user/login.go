// auth_handlers.go

package handlers_user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func Login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest LoginRequest
		if err := c.BindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка JSON формата"})
			return
		}
		var userID int
		var hashedPassword string
		sqlStatement := `SELECT id, password FROM users WHERE login = $1`
		err := db.QueryRow(sqlStatement, loginRequest.Login).Scan(&userID, &hashedPassword)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка полномочий"})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginRequest.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка полномочий"})
			return
		}
		token, err := generateRandomToken(32)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
			return
		}
		updateTokenStatement := `UPDATE users SET access_token = $1, status = 'перерыв' WHERE id = $2`
		_, err = db.Exec(updateTokenStatement, token, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления токена"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
