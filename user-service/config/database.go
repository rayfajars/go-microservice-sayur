package config

import (
	"fmt"
	"user-service/database/seeds"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgres() (*Postgres, error) {
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.PqslDB.User, cfg.PqslDB.Password, cfg.PqslDB.Host, cfg.PqslDB.Port, cfg.PqslDB.DBName)

	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-1] Failed to connect to database " + cfg.PqslDB.Host)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-2] Failed to get database connection")
		return nil, err
	}

	seeds.SeedRole(db)
	seeds.SeedAdmin(db)

	sqlDB.SetMaxOpenConns(cfg.PqslDB.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.PqslDB.DBMaxIdle)

	return &Postgres{DB: db}, nil
}
