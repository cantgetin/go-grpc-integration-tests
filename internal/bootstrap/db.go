package bootstrap

import (
	"fmt"
	"go-integration-tests/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	gormCfg := &gorm.Config{
		DisableAutomaticPing: true,
	}

	db, err := gorm.Open(postgres.Open(cfg.PgDSN), gormCfg)
	if err != nil {
		return nil, fmt.Errorf("failed connect to db, %v", err)
	}

	return db, nil
}
