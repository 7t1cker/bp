// handlers/user_handler.go

package handlers_skills

import (
	"database/sql"
	"net/http"

	"github.com/7t1cker/bp/api/models"
	"github.com/gin-gonic/gin"
)

func CreateSkills(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var newSkills models.Skills
        if err := c.BindJSON(&newSkills); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        sqlStatement := `
            INSERT INTO skills (skill)
            VALUES ($1)
            RETURNING id`
        err := db.QueryRow(sqlStatement, newSkills.Name).Scan(&newSkills.ID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, gin.H{"message": "Группа создана", "skills_id": newSkills.ID})
    }
}