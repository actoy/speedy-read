package symbol

import "time"

type Symbol struct {
	ID        int64
	Symbol    string
	Company   string
	Source    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
