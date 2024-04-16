package handlers_qests

import (
	"database/sql"
	"net/http"

	"github.com/7t1cker/bp/api/models"
	"github.com/gin-gonic/gin"
)

func GetQests(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
		headers := c.Request.Header
        for key, values := range headers {
            for _, value := range values {
                println(key+":", value)
            }
        }
        rows, err := db.Query("SELECT * FROM quests")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer rows.Close()
        var tasks []models.Quest
        for rows.Next() {
            var task models.Quest
            err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Deadline, &task.CreatorID, &task.Cost, &task.Priority, &task.SkillsRequired, &task.RecurrenceLimit)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            assignedTasks, err := getAssignedTasksForTask(db, task.ID)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            task.AssignedQuests = assignedTasks
            tasks = append(tasks, task)
        }
        if err := rows.Err(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"tasks": tasks})
    }
}
func getAssignedTasksForTask(db *sql.DB, taskID int) ([]models.AssignedQuest, error) {
    rows, err := db.Query("SELECT * FROM assignee_quests WHERE quest_id = $1", taskID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var assignedTasks []models.AssignedQuest
    for rows.Next() {
        var assignedTask models.AssignedQuest
        err := rows.Scan(&assignedTask.ID, &assignedTask.AssigneeID, &assignedTask.QuestID, &assignedTask.RecurrenceLimit, &assignedTask.CreationTimestamp, &assignedTask.ClosingTimestamp, &assignedTask.Done, &assignedTask.QuestsType)
        if err != nil {
            return nil, err
        }
        assignedTasks = append(assignedTasks, assignedTask)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return assignedTasks, nil
}