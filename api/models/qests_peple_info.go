package models

import "encoding/json"

type AssignedQuestWithTaskInfo struct {
	ID                int             `json:"id"`
	AssigneeID        int             `json:"assignee_id"`
	QuestID           int             `json:"quest_id"`
	RecurrenceLimit   int             `json:"recurrence_limit"`
	CreationTimestamp string          `json:"creation_timestamp"`
	ClosingTimestamp  string          `json:"closing_timestamp"`
	Done              bool            `json:"done"`
    QuestsType int `json:"quests_types ,omitempty"`
	Title             string          `json:"title"`
	Description       string          `json:"description"`
	Deadline          string          `json:"deadline"`
	Cost              float64                `json:"cost"`
	Priority          int             `json:"priority"`
	SkillsRequired    json.RawMessage `json:"skills_required"`
	Recurrence        int             `json:"recurrence"`
}
