package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)
type User struct {
    ID          int            `json:"id"`
    FirstName   string         `json:"first_name"`
    LastName    string         `json:"last_name"`
    MiddleName  sql.NullString `json:"middle_name"`
    DivisionID  sql.NullInt64  `json:"division_id"`
    GroupID     sql.NullInt64  `json:"group_id"`
    SkillTasks  []byte         `json:"skill_tasks"`
    Login       string         `json:"login"`
    Password    string         `json:"password"`
    Role        string         `json:"role"`
    Status      string         `json:"status"`
    Balance     float64        `json:"balance"`
    AccessToken string         `json:"access_token"`
}