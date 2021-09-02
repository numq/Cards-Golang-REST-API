package main

import (
	card "cardsRestApi/core"
	"cardsRestApi/database"
	"cardsRestApi/entity"
	"cardsRestApi/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	cfg := database.LoadConfig("./database")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: cfg.GetDbUrl(),
	}), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to the DB.")
	}

	db.AutoMigrate(&entity.Card{})

	repo := card.NewRepository(db)
	service := card.NewService(repo)
	handler := card.NewHandler(service)

	err = server.NewServer(handler.Init(), cfg.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}
}
