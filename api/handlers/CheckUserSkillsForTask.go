package handlers

import (
	"database/sql"
	"net/http"

	"github.com/7t1cker/bp/api/models"
	"github.com/7t1cker/bp/api/utils/sql_tools"
	"github.com/7t1cker/bp/api/utils/validation"
	"github.com/gin-gonic/gin"
)

func PurposeSkillsForTask2(db *sql.DB) gin.HandlerFunc {
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

		sqlStatement := `
            INSERT INTO users (first_name, last_name, middle_name, division_id, group_id, skill_tasks, login, password, role, status, access_token)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
            RETURNING id`
		err := db.QueryRow(sqlStatement,
			newUser.FirstName,
			newUser.LastName,
			newUser.MiddleName,
			newUser.DivisionID,
			newUser.GroupID,
			newUser.SkillTasks,
			newUser.Login,
			newUser.Password,
			newUser.Role,
			newUser.Status,
			newUser.AccessToken).Scan(&newUser.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Пользователь создан", "user_id": newUser.ID})
	}
}