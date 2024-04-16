package handlers_user

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/7t1cker/bp/api/models"
	"github.com/7t1cker/bp/api/utils/sql_tools"
	"github.com/gin-gonic/gin"
)

func AddSkillFromLearning(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("Api-Key")
		statusCode, userID := sql_tools.GetUserRequest(db, apiKey)
		if statusCode != http.StatusOK {
			c.JSON(statusCode, gin.H{"error": "Ошибка токена"})
			return
		}
		var req LearningRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}
		learning, err := getLearningByID(db, req.LearningID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Обучение с указанным идентификатором не найдено"})
			return
		}
		userSkills, err := getUserSkills(db, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении навыков пользователя"})
			return
		}
		hasSkill := haveCommonSkill(learning.Skill, userSkills)
		if hasSkill {
			c.JSON(http.StatusOK, gin.H{"message": "У пользователя уже есть этот скилл"})
			return
		}
		err = addSkillToUser(db, userID, learning.Skill)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении скилла сотруднику"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Скилл успешно добавлен сотруднику из обучения"})
	}
}
func getLearningByID(db *sql.DB, learningID int) (*models.Learning, error) {
	var learning models.Learning
	err := db.QueryRow("SELECT skill FROM learn WHERE id = $1", learningID).Scan(&learning.Skill)
	if err != nil {
		return nil, err
	}
	return &learning, nil
}
func addSkillToUser(db *sql.DB, userID int, skillID int) error {
	var skillTasksJSON []byte
	err := db.QueryRow("SELECT skill_tasks FROM users WHERE id = $1", userID).Scan(&skillTasksJSON)
	if err != nil {
		return err
	}
	var skillTasks []int
	if err = json.Unmarshal(skillTasksJSON, &skillTasks); err != nil {
		return err
	}
	skillTasks = append(skillTasks, skillID)
	updatedSkillTasksJSON, err := json.Marshal(skillTasks)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE users SET skill_tasks = $1 WHERE id = $2", updatedSkillTasksJSON, userID)
	if err != nil {
		return err
	}
	return nil
}

func haveCommonSkill(task int, required []int) bool {
	for _, req := range required {
		if task == req {
			return true
		}
	}
	return false
}
func getUserSkills(db *sql.DB, userID int) ([]int, error) {
	var skillTasksJSON []byte
	err := db.QueryRow("SELECT skill_tasks FROM users WHERE id = $1", userID).Scan(&skillTasksJSON)
	if err != nil {
		return nil, err
	}
	var skillTasks []int
	if err = json.Unmarshal(skillTasksJSON, &skillTasks); err != nil {
		return nil, err
	}
	return skillTasks, nil
}

// Структура для запроса на добавление скилла из обучения
type LearningRequest struct {
	LearningID int `json:"learning_id"`
}
