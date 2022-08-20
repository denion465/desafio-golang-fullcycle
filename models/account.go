package models

import "time"

type Account struct {
	ID             string
	Account_number string
	Balance        float64
	Created_at     time.Time
}
