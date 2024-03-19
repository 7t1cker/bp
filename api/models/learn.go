package models

type Learning struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"learn_title"`
	Skill int    `json:"skill"`
}