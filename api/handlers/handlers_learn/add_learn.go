package handlers_learn

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/7t1cker/bp/api/models"
	"github.com/gin-gonic/gin"
)

func CreateLearn(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var learn models.Learning
        if err := c.BindJSON(&learn); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        var existingLearn int
        err := db.QueryRow("SELECT COUNT(*) FROM learn WHERE skill = $1", learn.Skill).Scan(&existingLearn)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса"})
            return
        }
        if existingLearn > 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Уже существует обучение с указанным скиллом"})
            return
        }
        var existingSkill int
        err = db.QueryRow("SELECT COUNT(*) FROM skills WHERE id = $1", learn.Skill).Scan(&existingSkill)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса"})
            return
        }
        if existingSkill == 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Указанный скилл не существует"})
            return
        }
        var id int
        err = db.QueryRow("INSERT INTO learn (learn_title, skill) VALUES ($1, $2) RETURNING id", learn.Title, learn.Skill).Scan(&id)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить данные в таблицу learn"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Запись успешно добавлена с ID %d", id)})
    }
}
