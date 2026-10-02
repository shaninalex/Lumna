package email

import "gitlab.com/shaninalex/lumna/app/platform/database"

type Sender struct {
	db   *database.DB
	repo Repository
}

func NewSender(db *database.DB, repo Repository) *Sender {
	s := &Sender{db: db, repo: repo}
	s.init()
	return s
}

func (s *Sender) init() {
	go s.process()
}

func (s *Sender) process() {
	for {
		// get pending emails from database
		// walk through the list
		// send each of that emails
		// mark email as sent and save it
		continue
	}
}
