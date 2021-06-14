package core

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	h.initRoutes(router)
	return router
}

func (h *Handler) initRoutes(router *gin.Engine) {
	routerGroup := router.Group("/api")
	routerGroup.GET("/status", h.status)
	routerGroup.GET("/cards", h.getAllCards)
	routerGroup.GET("/cards/:id", h.getCardById)
	routerGroup.POST("/cards", h.createCard)
	routerGroup.POST("/cards/:id", h.updateCard)
	routerGroup.DELETE("/cards/:id", h.deleteCard)
}
