package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type QueuedChange struct {
	bun.BaseModel `bun:"table:queued_change,alias:qc"`
	ID            string    `bun:",pk" json:"id"`
	Release       Release   `json:"release"`
	Type          string    `json:"type"`
	TargetApp     string    `json:"targetApp"`
	LastChecked   time.Time `json:"lastChecked"`
	Details       string    `json:"details"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
}

func (qc QueuedChange) ToSC() StoredChange {
	return StoredChange{
		ID:          qc.ID,
		Release:     qc.Release,
		Type:        qc.Type,
		TargetApp:   qc.TargetApp,
		LastChecked: qc.LastChecked,
		Details:     qc.Details,
		CreatedAt:   qc.CreatedAt,
	}
}

type StoredChange struct {
	bun.BaseModel `bun:"table:stored_change,alias:sc"`
	ID            string    `bun:",pk" json:"id"`
	Release       Release   `json:"release"`
	Type          string    `json:"type"`
	TargetApp     string    `json:"targetApp"`
	LastChecked   time.Time `json:"lastChecked"`
	Details       string    `json:"details"`
	CreatedAt     time.Time `json:"createdAt"`
	CompletedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"completedAt"`
}
