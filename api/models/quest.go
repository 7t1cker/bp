package models

import "encoding/json"
type Quest struct {

	ID              int             `json:"id,omitempty"`
    Title           string          `json:"title"`
    Description     string          `json:"description"`
    Deadline        string          `json:"deadline"`
    CreatorID       int             `json:"creator_id,omitempty"`
    Cost            float64         `json:"cost"`
    Priority        int             `json:"priority,omitempty"`
    SkillsRequired  json.RawMessage `json:"skills_required"`
    RecurrenceLimit int             `json:"recurrence_limit,omitempty"`
    AssignedQuests  []AssignedQuest `json:"assigned_quests ,omitempty"`
    
}