package handlers_learn

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteLearning(db *sql.DB) gin.HandlerFunc {
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

		var existingLearn int
		err = db.QueryRow("SELECT COUNT(*) FROM learn WHERE id = $1", id).Scan(&existingLearn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при запросе к базе данных"})
			return
		}
		if existingLearn == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Обучение с указанным идентификатором не найдено"})
			return
		}

		_, err = db.Exec("DELETE FROM learn WHERE id = $1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении обучения из базы данных"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Обучение с идентификатором %d успешно удалено", id)})
	}
}
