// handlers/quest_handler.go

package handlers_qests

import (
	"database/sql"
	"net/http"

	"github.com/7t1cker/bp/api/utils/sql_tools"
	"github.com/gin-gonic/gin"
)


func MarkQuestAsDone(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
	
		var requestData struct {
			QuestID int    `json:"quest_id"`
		}

		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		apiKey := c.Request.Header.Get("Api-Key")
		statusCode, userID := sql_tools.GetUserRequest(db, apiKey)
		if statusCode != http.StatusOK {
			c.JSON(statusCode, gin.H{"error": "Ошибка токена"})
			return
		}

		sqlStep := `SELECT assignee_id FROM assignee_quests WHERE id = $1`
		var assigneeID *int
		err := db.QueryRow(sqlStep, requestData.QuestID).Scan(&assigneeID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if assigneeID != nil && *assigneeID != userID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Данная задача назначена на другого сотрудника."})
			return
		}

		var done bool
		sqlCheckDone := `SELECT done FROM assignee_quests WHERE id = $1`
		err = db.QueryRow(sqlCheckDone, requestData.QuestID).Scan(&done)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if done {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Эта задача уже выполнена"})
			return
		}
		sqlStatement := `UPDATE assignee_quests
		SET done = true, assignee_id = $1, closing_timestamp = CURRENT_TIMESTAMP
		WHERE id = $2;
		`
		_, err = db.Exec(sqlStatement, userID,requestData.QuestID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка с статусом задачи"})
			return
		}
		sqlPay :=`UPDATE users
		SET balance = users.balance + quests.cost
		FROM assignee_quests
		JOIN quests ON assignee_quests.quest_id = quests.id
		WHERE users.id = assignee_quests.assignee_id
		AND assignee_quests.id = $1;`
		_, err = db.Exec(sqlPay, requestData.QuestID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка начисления"})
			return
		}
		sqlReplik := `SELECT q.recurrence_limit - aq.recurrence AS difference,
		aq.recurrence,
		aq.quest_id
		FROM assignee_quests AS aq
		JOIN quests AS q ON aq.quest_id = q.id
		WHERE aq.id = $1;
		`
 	var limit int
 	var recurrence int
 	var questID int
 	err = db.QueryRow(sqlReplik, requestData.QuestID).Scan(&limit, &recurrence, &questID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var questTypeID int
	sqlQuestType := `SELECT quest_type_id FROM assignee_quests WHERE id = $1`
	err = db.QueryRow(sqlQuestType, requestData.QuestID).Scan(&questTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var percentage float64 = 1
	if questTypeID == 2 || questTypeID == 3{
		
		sqlBalanceUpdate := `
		UPDATE users
		SET balance = balance + (
    	SELECT CASE 
        	WHEN CURRENT_TIMESTAMP > ht.end_time THEN 0
        	ELSE ht.hot * q.cost *$2
    	END
    	FROM hot_tasks ht
    	JOIN assignee_quests aq ON ht.assignee_quest_id = aq.id
    	JOIN quests q ON aq.quest_id = q.id
    	WHERE aq.id = $1)
		WHERE id = (SELECT assignee_id FROM assignee_quests WHERE id = $1)
		`
    _, err = db.Exec(sqlBalanceUpdate, requestData.QuestID,)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка начисления на баланс"})
        return
    }
	}
	

	if limit > 0 {
			_, err = sql_tools.AddAssignee(db, questID, recurrence + 1, 1)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    }
		}
		c.JSON(http.StatusOK, gin.H{"message": "Здача выполнена", "pr" : percentage})
	}
}
