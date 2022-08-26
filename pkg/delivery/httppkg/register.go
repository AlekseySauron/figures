package httppkg

import (
	"net/http"

	"github.com/AlekseySauron/figures/pkg/services/mathpkg"
	"github.com/gin-gonic/gin"
)

type task struct {
	Figure string  `json:"figure"`
	H      float64 `json:"h"`
	W      float64 `json:"w"`
}

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
	var newTask task

	err := c.BindJSON(&newTask)
	if err != nil {
		return
	}

	if newTask.Figure == "square" {
		// g := mathpkg.Square{newTask.H, newTask.W}
		g := mathpkg.Square{newTask.H, newTask.W}
	} else if newTask.Figure == "circle" {
		g := mathpkg.Circle{newTask.H}
	} else if newTask.Figure == "rectangle" {
		g := mathpkg.Rectangle{newTask.H, newTask.W}
	} else if newTask.Figure == "triangle" {
		g := mathpkg.Triangle{newTask.H, newTask.W}
	}

	// c.JSON(http.StatusOK, g.area())
	c.JSON(http.StatusOK, measure(g))

}

func measure(g mathpkg.Geometry) {
	return g.Area()
}
