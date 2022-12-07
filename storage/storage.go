package storage

import (
	"reyting_service/storage/postgres"
	"reyting_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Reyting() repo.ReytingStorageI
}

type storagePg struct {
	db          *sqlx.DB
	reytingRepo repo.ReytingStorageI
}

func NewStorage(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		reytingRepo: postgres.NewReytingRepo(db),
	}
}

func (s *storagePg) Reyting() repo.ReytingStorageI {
	return s.reytingRepo
}
