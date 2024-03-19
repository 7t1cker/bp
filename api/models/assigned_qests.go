package models

import "database/sql"

type AssignedQuest struct {
    ID                int            `json:"id"`
    AssigneeID        sql.NullInt64  `json:"assignee_id"`
    QuestID           int            `json:"quest_id"`
    RecurrenceLimit   int            `json:"recurrence"`
    CreationTimestamp sql.NullString `json:"creation_timestamp"`
    ClosingTimestamp  sql.NullString `json:"closing_timestamp"`
    Done              bool           `json:"done"`
    QuestsType        int            `json:"quests_type,omitempty"`

}