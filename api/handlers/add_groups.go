package handlers

import (
	"database/sql"
	"net/http"

	"github.com/7t1cker/bp/api/models"
	"github.com/gin-gonic/gin"
)

func CreateGroups(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var newGroup models.Group
        if err := c.BindJSON(&newGroup); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        sqlStatement := `
            INSERT INTO groups (group_name)
            VALUES ($1)
            RETURNING id`
        err := db.QueryRow(sqlStatement, newGroup.Name).Scan(&newGroup.ID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, gin.H{"message": "Группа создана", "group_id": newGroup.ID})
    }
}
