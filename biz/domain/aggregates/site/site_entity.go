package site

import "time"

type Site struct {
	ID          int64
	SourceID    int64
	SourceType  string
	Url         string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
