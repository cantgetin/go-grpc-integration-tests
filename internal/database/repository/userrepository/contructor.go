package userrepository

import "gorm.io/gorm"

type RepositoryImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}
