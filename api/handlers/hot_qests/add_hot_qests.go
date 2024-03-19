package hot_qests

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/7t1cker/bp/api/models"
	"github.com/gin-gonic/gin"
)

func CreateHotTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newTask models.HotQest
		if err := c.BindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newTask.CreationTime = time.Now().Format("2006-01-02 15:04:05")

		endTime, err := time.Parse("2006-01-02 15:04:05", newTask.EndTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка формата даты"})
			return
		}
		var bp bool
		err = db.QueryRow("SELECT done From assignee_quests WHERE id = $1",newTask.AssigneeQuestID).Scan(&bp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if bp{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Задача уже закрыта"})
			return
		}
		
		var pix bool
		err = db.QueryRow("SELECT $1 > $2", endTime, newTask.CreationTime).Scan(&pix)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Сравнение дат"})
			return
		}
		if !pix {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Дата окончания меньше даты начала"})
			return
		}
		var existingTaskID int
		var dn bool
		err = db.QueryRow(`SELECT hot_tasks.id, assignee_quests.done
		FROM hot_tasks
		JOIN assignee_quests ON hot_tasks.assignee_quest_id = assignee_quests.id
		WHERE assignee_quests.id = $1
		`, newTask.AssigneeQuestID).Scan(&existingTaskID, &dn)
		fmt.Println(dn)
		if dn {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Указанная задача уже закрыта"})
			return
		}
		switch {
		case err == sql.ErrNoRows:
		case err != nil:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "assignee_quest_id"})
			return
		}
		var i int = 2
		if newTask.Fire{
			newTask.Hot *= 2
			i = 3
		}
		sqlStatement := `
            INSERT INTO hot_tasks (assignee_quest_id, creation_time, hot, end_time, fire)
            VALUES ($1, $2, $3, $4, $5)
            RETURNING id`
		err = db.QueryRow(sqlStatement, newTask.AssigneeQuestID, newTask.CreationTime, newTask.Hot, newTask.EndTime, newTask.Fire).Scan(&newTask.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		
			sqlStatement = `
			UPDATE assignee_quests
			SET quest_type_id = $2
			WHERE id = $1`
		_, err = db.Exec(sqlStatement, newTask.AssigneeQuestID, i)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Горящая задача создана", "hot_task": newTask})
	}
}
