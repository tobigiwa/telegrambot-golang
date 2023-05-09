package store

import "errors"

var (
	ErrDuplicateID       = errors.New("id already exists")
	ErrDuplicateUsername = errors.New("username already exists")
	ErrNoRows            = errors.New("row not exists")
	ErrUpdateFailed      = errors.New("update failed")
	ErrDeleteFailed      = errors.New("delete failed")
)
