package store

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type USER struct {
	ID       int64
	Username string
}

type Respository struct {
	db *pgxpool.Pool
}

func (DB *Respository) Migrate() error {
	query := `CREATE TABLE IF NOT EXISTS users_tbl(id INTEGER PRIMARY KEY, username TEXT);`
	_, err := DB.db.Exec(context.Background(), query)
	return err
}

func NewRespository(db *pgxpool.Pool) *Respository {
	return &Respository{
		db: db,
	}
}

func (DB *Respository) Insert(userID int64, userUsername string) error {
	query := `insert into users_tbl values($1, $2)`
	_, err := DB.db.Exec(context.Background(), query, userID, userUsername)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				switch {
				case strings.Contains(pgxError.Detail, "id"):
					return ErrDuplicateID
				case strings.Contains(pgxError.Detail, "username"):
					return ErrDuplicateUsername
				}
			} else {
				return err
			}
		}
	}
	return nil
}

// AllIDs returns a slice of all user id and an error
func (DB *Respository) AllIDs() ([]int64, error) {
	query := `select id from users_tbl`
	rows, err := DB.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var userIDs []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, id)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return userIDs, nil
}

// IsUser returns true if user is found in the db and false otherwise.
func (DB *Respository) IsUser(id int64) bool {
	query := `select * from users_tbl where id = $1`
	row := DB.db.QueryRow(context.Background(), query, id)
	var x USER
	if err := row.Scan(&x.ID, &x.Username); err != nil {
		return false
	}
	return true
}

func (DB *Respository) Delete(id int64) error {
	query := `DELETE FROM users_tbl WHERE id = $1`
	commandTag, err := DB.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return ErrDeleteFailed
	}
	return nil
}
