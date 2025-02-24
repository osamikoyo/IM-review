package data

import (
	"github.com/osamikoyo/IM-review/internal/config"
	"github.com/osamikoyo/IM-review/internal/data/models"
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

	err = db.AutoMigrate(&models.Review{})

	return &Storage{db: db}, err	
}

func (s *Storage) Add(review *models.Review) error {
	return s.db.Create(review).Error
}

func (s *Storage) Get(id uint64) ([]models.Review, error) {
	var reviews []models.Review

	result := s.db.Where(&models.Review{
		ProductID: id,
	}).Find(&reviews)

	return reviews, result.Error
}