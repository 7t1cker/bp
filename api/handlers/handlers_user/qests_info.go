package handlers_user

import (
	"database/sql"
	"net/http"

	"github.com/7t1cker/bp/api/models"
	"github.com/gin-gonic/gin"
)


func GetAssignedQuests(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.Param("user_id")
        rows, err := db.Query(`
            SELECT aq.*, q.title, q.description, q.deadline, q.cost, q.priority, q.skills_required, q.recurrence_limit
            FROM assignee_quests aq
            JOIN quests q ON aq.quest_id = q.id
            WHERE aq.assignee_id = $1
        `, userID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer rows.Close()
        var assignedQuests []models.AssignedQuestWithTaskInfo
        for rows.Next() {
            var assignedTask models.AssignedQuestWithTaskInfo
            err := rows.Scan(
                &assignedTask.ID,
                &assignedTask.AssigneeID,
                &assignedTask.QuestID,
                &assignedTask.RecurrenceLimit,
                &assignedTask.CreationTimestamp,
                &assignedTask.ClosingTimestamp,
                &assignedTask.Done,
                &assignedTask.QuestsType,
                &assignedTask.Title,
                &assignedTask.Description,
                &assignedTask.Deadline,
                &assignedTask.Cost,
                &assignedTask.Priority,
                &assignedTask.SkillsRequired,
                &assignedTask.RecurrenceLimit,
            )
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            assignedQuests = append(assignedQuests, assignedTask)
        }
        if err := rows.Err(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"assigned_quests": assignedQuests})
    }
}
