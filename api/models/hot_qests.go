package models

type HotQest struct {
	ID              int     `json:"id,omitempty"`
	AssigneeQuestID int     `json:"assignee_quest_id"`
	CreationTime    string  `json:"creation_time,omitempty"`
	Hot             float64 `json:"hot"`
	EndTime         string  `json:"end_time"`
	Fire            bool    `json:"fire"`
}