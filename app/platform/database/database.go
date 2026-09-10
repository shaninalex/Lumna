package database

import "gitlab.com/shaninalex/lumna/app/platform/config"

type DB struct {
}

func New(cfg *config.Config) *DB {
	return &DB{}
}
