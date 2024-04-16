package models

import "encoding/json"
type User struct {
	ID          int     `json:"id,omitempty"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	MiddleName  string  `json:"middle_name"`
	DivisionID  int64   `json:"division_id"`
	GroupID     int64   `json:"group_id"`
	SkillTasks  json.RawMessage  `json:"skill_tasks"`
	Login       string  `json:"login"`
	Password    string  `json:"password"`
	Role        string  `json:"role"`
	Status      string  `json:"status,omitempty"`
	Balance     float64 `json:"balance,omitempty"`
	AccessToken string  `json:"access_token,omitempty"`
}