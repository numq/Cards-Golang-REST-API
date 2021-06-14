package core

import (
	"cardsRestApi/entity"
	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetAllCards() ([]*entity.Card, error) {
	cards, err := s.repo.GetAllCards()
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (s *Service) GetCardById(id uuid.UUID) (*entity.Card, error) {
	card, err := s.repo.GetCardById(id)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (s *Service) CreateCard(card *entity.Card) (*entity.Card, error) {
	card.ID = uuid.New()
	card, err := s.repo.CreateCard(card)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (s *Service) UpdateCard(card *entity.Card) (*entity.Card, error) {
	card, err := s.repo.UpdateCard(card)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (s *Service) DeleteCard(card *entity.Card) (*entity.Card, error) {
	card, err := s.repo.DeleteCard(card)
	if err != nil {
		return nil, err
	}
	return card, nil
}
