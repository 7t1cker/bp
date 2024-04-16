package handlers_learn

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/7t1cker/bp/api/models"
	"github.com/gin-gonic/gin"
)

func GetAllLearnings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, learn_title, skill FROM learn")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при запросе к базе данных"})
			return
		}
		defer rows.Close()
		learnings := []models.Learning{}
		for rows.Next() {
			var learning models.Learning
			err := rows.Scan(&learning.ID, &learning.Title, &learning.Skill)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сканировании результатов запроса"})
				return
			}
			learnings = append(learnings, learning)
		}
		c.JSON(http.StatusOK, learnings)
	}
}

func GetLearningByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		if idStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Не указан идентификатор обучения"})
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат идентификатора обучения"})
			return
		}
		var learning models.Learning
		err = db.QueryRow("SELECT id, learn_title, skill FROM learn WHERE id = $1", id).Scan(&learning.ID, &learning.Title, &learning.Skill)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Обучение с указанным идентификатором не найдено"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при запросе к базе данных"})
			return
		}

		c.JSON(http.StatusOK, learning)
	}
}
