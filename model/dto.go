package model

import "time"

type Note struct {
	ID         string
	Content    string
	StartTime  time.Time
	UpdateTime time.Time
}
