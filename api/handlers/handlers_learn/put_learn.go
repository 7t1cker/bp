package handlers_learn

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateLearnTitle(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var requestData struct {
            LearnTitle string `json:"learn_title"`
        }
        if err := c.BindJSON(&requestData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
            return
        }
        var existingLearn int
        err := db.QueryRow("SELECT COUNT(*) FROM learn WHERE id = $1", id).Scan(&existingLearn)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса"})
            return
        }
        if existingLearn == 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Обучение с указанным идентификатором не найдено"})
            return
        }
        if requestData.LearnTitle == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Заголовок обучения не может быть пустым"})
            return
        }
        _, err = db.Exec("UPDATE learn SET learn_title = $1 WHERE id = $2", requestData.LearnTitle, id)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении заголовка обучения"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Заголовок обучения успешно обновлен на '%s'", requestData.LearnTitle)})
    }
}
