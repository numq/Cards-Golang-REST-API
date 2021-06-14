package core

import (
	"cardsRestApi/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	db.Logger.LogMode(logger.Info)
	return &Repository{db: db}
}

func (r *Repository) GetAllCards() ([]*entity.Card, error) {
	var cards []*entity.Card
	err := r.db.Find(&cards).Error
	return cards, err
}

func (r *Repository) GetCardById(id uuid.UUID) (*entity.Card, error) {
	var card *entity.Card
	err := r.db.Where(entity.Card{ID: id}).Take(&card).Error
	return card, err
}

func (r *Repository) CreateCard(card *entity.Card) (*entity.Card, error) {
	err := r.db.Model(&card).Create(&card).Error
	return card, err
}

func (r *Repository) UpdateCard(card *entity.Card) (*entity.Card, error) {
	err := r.db.Model(&card).Save(entity.Card{
		ID:        uuid.UUID{},
		FrontText: "",
		BackText:  "",
		UpdatedAt: time.Time{},
	}).Error
	return card, err
}

func (r *Repository) DeleteCard(card *entity.Card) (*entity.Card, error) {
	err := r.db.Where(entity.Card{}).Delete(&card).Error
	return card, err
}
