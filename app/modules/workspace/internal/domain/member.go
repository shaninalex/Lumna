package domain

import "time"

type Member struct {
	ID         int
	Email      string
	FullName   string
	Image      string
	Role       string
	DateJoined time.Time
}
