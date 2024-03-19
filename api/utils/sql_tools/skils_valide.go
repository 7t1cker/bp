package sql_tools

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)
func ValidateSkillTasks(db *sql.DB, skillTasks json.RawMessage) error {
	var skills []int64
	err := json.Unmarshal(skillTasks, &skills)
	if err != nil {
		return err
	}
    fmt.Println("Skills:", skills)
	inQuery := "("
	for _, id := range skills {
		inQuery += fmt.Sprintf("%d,", id)
	}
	inQuery = strings.TrimSuffix(inQuery, ",") + ")"
	query := fmt.Sprintf("SELECT COUNT(*) FROM skills WHERE id IN %s", inQuery)

	var count int
	err = db.QueryRow(query).Scan(&count)
	if err != nil {
		return err
	}
	if count != len(skills) {
		return errors.New("One or more skill IDs are invalid")
	}

	return nil
}
