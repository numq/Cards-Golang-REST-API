package domain

import (
	"context"
	"time"
)

type Card struct {
	ID        int64     `json:"id"`
	FrontText string    `json:"front_text"`
	BackText  string    `json:"back_text"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type CardUseCase interface {
	GetAllCards(ctx context.Context, cursor string, num int64) ([]Card, string, error)
	GetByID(ctx context.Context, id int64) (Card, error)
	InsertCard(ctx context.Context, card *Card) error
	DeleteCard(ctx context.Context, id int64) error
}

type CardRepository interface {
	GetAllCards(ctx context.Context, cursor string, num int64) ([]Card, string, error)
	GetByID(ctx context.Context, id int64) (Card, error)
	InsertCard(ctx context.Context, card *Card) error
	DeleteCard(ctx context.Context, id int64) error
}
