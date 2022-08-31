package httppkg

import (
	"fmt"
	"net/http"

	"github.com/AlekseySauron/figures/pkg/delivery/telegrampkg"
	"github.com/AlekseySauron/figures/pkg/services/mathpkg"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func Register(router *gin.Engine) {
	h := NewHandler()

	router.POST("", h.Posting)
}

func (h *Handler) Posting(c *gin.Context) {
	var newTask Task

	err := c.BindJSON(&newTask)
	if err != nil {
		return
	}

	var figure mathpkg.Geometry

	if newTask.Figure == "square" {
		figure = mathpkg.NewSquare(newTask.H, newTask.W)
	} else if newTask.Figure == "circle" {
		figure = mathpkg.NewCircle(newTask.H)
	} else if newTask.Figure == "rectangle" {
		figure = mathpkg.NewRectangle(newTask.H, newTask.W)
	} else if newTask.Figure == "triangle" {
		figure = mathpkg.NewTriangle(newTask.H, newTask.W)
	} else {
		c.JSON(http.StatusBadRequest, "unknow Figure")
		return
	}

	chat_id := c.Request.Header.Get("chat_id")

	go func(chat_id string, figure mathpkg.Geometry, figureName string) {
		res := mathpkg.Measure(figure)

		telegrampkg.Send(chat_id, fmt.Sprintf("Result for figure %s: %f", figureName, res))
	}(chat_id, figure, newTask.Figure)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
