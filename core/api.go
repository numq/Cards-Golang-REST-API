package core

import (
	"cardsRestApi/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) status(context *gin.Context) {
	context.JSON(http.StatusOK, "Service is OK!")
}

func (h *Handler) getAllCards(context *gin.Context) {

	cards, err := h.service.GetAllCards()
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	context.JSON(http.StatusOK, cards)
}

func (h *Handler) getCardById(context *gin.Context) {

	id, _ := uuid.Parse(context.Param("id"))
	card, err := h.service.GetCardById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	context.JSON(http.StatusOK, card)
}

func (h *Handler) createCard(context *gin.Context) {

	var input *entity.Card
	context.BindJSON(&input)
	card, err := h.service.CreateCard(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, card)
}

func (h *Handler) updateCard(context *gin.Context) {

	var input *entity.Card
	context.BindJSON(&input)
	card, err := h.service.UpdateCard(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, card)
}

func (h *Handler) deleteCard(context *gin.Context) {

	var input *entity.Card
	context.BindJSON(&input)
	card, err := h.service.DeleteCard(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, card)
}
