package store

import "errors"

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNoRows       = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)
