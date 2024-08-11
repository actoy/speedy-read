package symbol

import "time"

type Symbol struct {
	ID              int64
	Symbol          string
	Company         string
	CompanyZH       string
	CompanyUrl      string
	CompanyAddress  string
	Description     string
	CompanyBusiness string
	Source          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
