package sql_tools

import (
	"database/sql"
)

func AddAssignee(db *sql.DB, questID int, recurrence int, quest_type_id int) (int, error) {
    var assigneeQuestID int
    sqlZapros := `
        INSERT INTO assignee_quests (quest_id, recurrence, quest_type_id)
        VALUES ($1, $2, $3)
        RETURNING id`
    
    err := db.QueryRow(sqlZapros, questID, recurrence, quest_type_id).Scan(&assigneeQuestID)
    if err != nil {
        return 0, err
    }

    return assigneeQuestID, nil
}
