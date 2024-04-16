package handlers_qests

import (
	"database/sql"
	"net/http"

	"github.com/7t1cker/bp/api/utils/sql_tools"
	"github.com/gin-gonic/gin"
)


func UpdateAssignee(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("Api-Key")
		statusCode, userID := sql_tools.GetUserRequest(db, apiKey)
		if statusCode != http.StatusOK {
			c.JSON(statusCode, gin.H{"error": "Ошибка токена"})
			return
		}
		var requestData struct {
			QuestID int `json:"quest_id"`
		}
		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ok, err := sql_tools.CheckSkills(db, requestData.QuestID, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Недостаточно прав для изменения задачи или недостаточно навыков"})
			return
		}
		sqlStatement := `
			UPDATE assignee_quests
			SET assignee_id = $1
			WHERE id = $2`
		_, err = db.Exec(sqlStatement, userID, requestData.QuestID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Задача назначена"})
	}
}
