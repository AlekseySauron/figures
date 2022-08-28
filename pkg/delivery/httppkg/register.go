package httppkg

import (
	"fmt"
	"net/http"

	"github.com/AlekseySauron/figures/pkg/delivery/telegrampkg"
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

	res := mathpkg.Measure(figure)
	c.JSON(http.StatusOK, res)

	newBot := telegrampkg.NewBot()
	// bot, err := tgbotapi.NewBotAPI("1901733643:AAHlKkQJrCaKS1c1SZigHXq6t8CUXO7eeWs")
	// if err != nil {
	// 	return
	// }

	// msg := tgbotapi.NewMessage(421964311, fmt.Sprint(res))
	newBot.Send(fmt.Sprint(res))
	//msg.ReplyToMessageID = update.Message.MessageID
	// bot.Send(msg)

}
