package handlers

import (
	"database/sql"
	"net/http"

	"github.com/7t1cker/bp/api/models"
	"github.com/gin-gonic/gin"
)


func CreateDivisions(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var newDivision models.Division
        if err := c.BindJSON(&newDivision); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        sqlStatement := `
            INSERT INTO divisions (division_name)
            VALUES ($1)
            RETURNING id`
        err := db.QueryRow(sqlStatement, newDivision.Name).Scan(&newDivision.ID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, gin.H{"message": "Подразделение создано", "division_id": newDivision.ID})
    }
}
