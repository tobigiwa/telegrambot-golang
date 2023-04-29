package store

import "database/sql"

type SQLiteRespository struct {
	db *sql.DB
}

func NewSQLiteRespository(db *sql.DB) *SQLiteRespository {
	return &SQLiteRespository{
		db: db,
	}
}

