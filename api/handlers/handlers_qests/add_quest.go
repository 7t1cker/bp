package handlers_qests

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/7t1cker/bp/api/models"
	"github.com/7t1cker/bp/api/utils/sql_tools"
	"github.com/gin-gonic/gin"
)

func CreateQuest(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
   
        var newQuest models.Quest
        if err := c.BindJSON(&newQuest); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        apiKey := c.Request.Header.Get("Api-Key")

        statusCode, userID := sql_tools.GetUserRequest( db, apiKey)
        if statusCode != http.StatusOK {
            c.JSON(statusCode, gin.H{"error": "Ошибка токена"})
            return
        }
        newQuest.CreatorID = userID
		skillsRequiredJSON, err := json.Marshal(newQuest.SkillsRequired)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки скилов"})
			return
		}
		if err := sql_tools.ValidateSkillTasks(db, newQuest.SkillsRequired); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		sqlStatement := `
            INSERT INTO quests (title, description, deadline, creator_id, cost, priority, skills_required,recurrence_limit)
            VALUES ($1, $2, $3, $4, $5, $6, $7,$8 )
            RETURNING id`
		var questID int
        var assigneeQuestID int
        if newQuest.RecurrenceLimit <= 0{
            newQuest.RecurrenceLimit = 1
        }
		err = db.QueryRow(sqlStatement,
			newQuest.Title,
			newQuest.Description,
			newQuest.Deadline,
			newQuest.CreatorID,
			newQuest.Cost,
			newQuest.Priority,
			skillsRequiredJSON,
            newQuest.RecurrenceLimit).Scan(&questID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
        

                assigneeQuestID, err = sql_tools.AddAssignee(db, questID, 1, 1)
                if err != nil{
                    c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания задачи"})
                    return
                }
		c.JSON(http.StatusCreated, gin.H{ "quest_id": assigneeQuestID })
	}
}
