package sql_tools

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type Skills struct {
	SkillTasks     *[]int `json:"skill_tasks"`
	SkillsRequired *[]int `json:"skills_required"`
}

func CheckSkills(db *sql.DB, QuestID int, userID int) (bool, error) {
	var SkillUser Skills
	var SkillTasks Skills

	if userID <= 0 || QuestID <= 0 {
		return false, fmt.Errorf("userID и QuestID должны быть положительными числами")
	}
	var assigneeID sql.NullInt64
	sqlStep := `SELECT assignee_id FROM assignee_quests WHERE id = $1`
	err := db.QueryRow(sqlStep, QuestID).Scan(&assigneeID)
	switch {
	case err == sql.ErrNoRows:
		return false, fmt.Errorf("назначение с QuestID=%d не найдено", QuestID)
	case err != nil:
		return false, err
	}

	if assigneeID.Valid && int(assigneeID.Int64) != userID {
		return false, fmt.Errorf("назначенный пользователь (assignee_id) не совпадает с userID")
	}
	var done bool
	sqlCheckDone := `SELECT done FROM assignee_quests WHERE id = $1`
	err = db.QueryRow(sqlCheckDone, QuestID).Scan(&done)
	switch {
	case err == sql.ErrNoRows:
		return false, fmt.Errorf("назначение с QuestID=%d не найдено", QuestID)
	case err != nil:
		return false, err
	}

	if done {
		return false, fmt.Errorf("задача с QuestID=%d уже выполнена", QuestID)
	}
	sqlQuery := `
		SELECT u.skill_tasks, q.skills_required
		FROM users AS u
		CROSS JOIN (
			SELECT q.skills_required
			FROM assignee_quests AS aq
			JOIN quests AS q ON aq.quest_id = q.id
			WHERE aq.id = $1
		) AS q
		WHERE u.id = $2;
	`
	var skillTasksJSON, skillsRequiredJSON []byte
	err = db.QueryRow(sqlQuery, QuestID, userID).Scan(&skillTasksJSON, &skillsRequiredJSON)
	switch {
	case err == sql.ErrNoRows:
		return false, fmt.Errorf("нет данных о навыках для пользователя с ID=%d и задачи с ID=%d", userID, QuestID)
	case err != nil:
		return false, err
	}
	if err := json.Unmarshal(skillTasksJSON, &SkillUser.SkillTasks); err != nil {
		return false, err
	}
	if err := json.Unmarshal(skillsRequiredJSON, &SkillTasks.SkillsRequired); err != nil {
		return false, err
	}
	if haveCommonSkill(*SkillUser.SkillTasks, *SkillTasks.SkillsRequired) {
		return true, nil
	}

	return false, nil
}
func haveCommonSkill(tasks []int, required []int) bool {
	for _, task := range tasks {
		for _, req := range required {
			if task == req {
				return true
			}
		}
	}
	return false
}
