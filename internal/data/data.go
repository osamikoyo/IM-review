package data

import (
	"github.com/osamikoyo/IM-review/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct{
	db *gorm.DB
}

func New(cfg *config.Config) (*Storage, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Dsn))
	if err != nil{
		return nil, err
	}

	
}