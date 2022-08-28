package httppkg

import (
	"net/http"

	"github.com/AlekseySauron/figures/pkg/services/mathpkg"
	"github.com/gin-gonic/gin"
)

// type Task struct {
// 	Figure string  `json:"figure"`
// 	H      float64 `json:"h"`
// 	W      float64 `json:"w"`
// }

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

	//figure := mathpkg.Square{H: newTask.H, W: newTask.W}

	var figure mathpkg.Geometry

	if newTask.Figure == "square" {
		//figure = mathpkg.Square{H: newTask.H, W: newTask.W}
		// figure = mathpkg.NewSquare(newTask)
		figure = mathpkg.NewSquare(newTask.H, newTask.W)
	} else if newTask.Figure == "circle" {
		// figure = mathpkg.Circle{R: newTask.H}
		figure = mathpkg.NewCircle(newTask.H)
	} else if newTask.Figure == "rectangle" {
		// figure = mathpkg.Rectangle{H: newTask.H, W: newTask.W}
		figure = mathpkg.NewRectangle(newTask.H, newTask.W)
	} else if newTask.Figure == "triangle" {
		// figure = mathpkg.Triangle{H: newTask.H, W: newTask.W}
		figure = mathpkg.NewTriangle(newTask.H, newTask.W)
	} else {
		c.JSON(http.StatusBadRequest, "unknow Figure")
		return
	}

	c.JSON(http.StatusOK, mathpkg.Measure(figure))

}
