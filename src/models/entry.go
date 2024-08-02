package models

import (
	"time"
)

type Entry struct {
	ID        string
	AddedAt   time.Time
	UpdatedAt time.Time
	Text      string
}
