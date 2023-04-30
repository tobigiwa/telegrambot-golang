package store

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

type USER struct {
	ID       int64
	Username string
}

type SQLiteRespository struct {
	db *sql.DB
}

func NewSQLiteRespository(db *sql.DB) *SQLiteRespository {
	return &SQLiteRespository{
		db: db,
	}
}

func (DB *SQLiteRespository) Insert(userID int64, userUsername string) error {
	query := `insert into users_tbl values(?, ?)`
	_, err := DB.db.Exec(query, userID, userUsername)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return ErrDuplicate
			} else {
				return err
			}
		}
	}
	return nil
}

func (DB *SQLiteRespository) All() ([]USER, error) {
	query := `select * from users_tbl`
	rows, err := DB.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var UserList []USER
	for rows.Next() {
		var user USER
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrNoRows
			} else {
				return nil, err
			}
		}
		UserList = append(UserList, user)
	}
	return UserList, nil
}

// IsUser returns true if user is found in the db and false otherwise.
func (DB *SQLiteRespository) IsUser(id int64) bool {
	query := `select * from users_tbl where id = ?`
	row := DB.db.QueryRow(query, id)
	var x USER
	if err := row.Scan(&x.ID, &x.Username); err != nil {
		return false
	}
	return true
}

func (DB *SQLiteRespository) Delete(id int64) error {
	query := `DELETE FROM users_tbl WHERE id = ?`
	res, err := DB.db.Exec(query, id)
	if err != nil {
		return err
	}
	if rowsAffected, err := res.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return nil
}
