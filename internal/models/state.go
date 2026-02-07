package models

import (
	"github.com/theunhackable/gator/internal/config"
	"github.com/theunhackable/gator/internal/db"
)

type State struct {
	Db    *db.Queries
	State *config.Config
}
