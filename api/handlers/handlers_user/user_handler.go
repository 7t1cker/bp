package handlers_user

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"net/http"

	"github.com/7t1cker/bp/api/models"
	"github.com/7t1cker/bp/api/utils/sql_tools"
	"github.com/7t1cker/bp/api/utils/validation"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


	func CreateUser(db *sql.DB) gin.HandlerFunc {
		return func(c *gin.Context) {
			var newUser models.User
			if err := c.BindJSON(&newUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := validation.ValidateUser(newUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := sql_tools.ValidateSkillTasks(db, newUser.SkillTasks); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хешировании пароля"})
				return
			}
			token, err := generateRandomToken(32)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
				return
			}

			sqlStatement := `
				INSERT INTO users (first_name, last_name, middle_name, division_id, group_id, skill_tasks, login, password, role, status, access_token)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, 'не работает', $10)
				RETURNING id`
			err = db.QueryRow(sqlStatement,
				newUser.FirstName,
				newUser.LastName,
				newUser.MiddleName,
				newUser.DivisionID,
				newUser.GroupID,
				newUser.SkillTasks,
				newUser.Login,
				hashedPassword,
				newUser.Role,
				token).Scan(&newUser.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, gin.H{"message": "Пользователь создан", "user_id": newUser.ID})
		}
	}
	func generateRandomToken(length int) (string, error) {
		b := make([]byte, length)
		_, err := rand.Read(b)
		if err != nil {
			return "", err
		}
		return base64.URLEncoding.EncodeToString(b), nil
	}
