package handlers

import "github.com/Miha-s/it_db_lab1/internal/database"

type HandlersFactory struct {
	database *database.Database
}

func NewHandlersFactory(db *database.Database) *HandlersFactory {
	return &HandlersFactory{
		database: db,
	}
}
