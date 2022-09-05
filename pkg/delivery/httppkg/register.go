package httppkg

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	//router.POST("/json", h.Posting)

	router.LoadHTMLGlob("../templates/*")
	router.GET("/", h.GettingHtml)
	router.POST("/", h.PostingHtml)
}

func (h *Handler) GettingHtml(c *gin.Context) {
	// tmpl, _ := template.ParseFiles("../templates/index.html")
	// tmpl.Execute(w, nil)
	// c.HTML(http.StatusOK, "index.html", gin.H{"result": "100"})
	c.HTML(http.StatusOK, "index.html", gin.H{"result": "0"})
	// h.GettingHtml("../templates/index.html")

}

func (h *Handler) PostingHtml(c *gin.Context) {
	figure := c.PostForm("Figures")

	chat_id := c.PostForm("chat_id")
	if chat_id == "" {
		c.JSON(http.StatusBadRequest, "chat_id не указан")
		return
	}

	var newTask Task
	if figure == "circle" {
		radius, err := strconv.ParseFloat(c.PostForm("radius"), 64)
		if err != nil {
			log.Fatal("Ошибка формата radius")
		}

		newTask = Task{figure, radius, 0}
	} else {
		width, err := strconv.ParseFloat(c.PostForm("width"), 64)
		if err != nil {
			log.Fatal("Ошибка формата width")
		}

		height, err := strconv.ParseFloat(c.PostForm("height"), 64)
		if err != nil {
			log.Fatal("Ошибка формата height")
		}

		newTask = Task{figure, height, width}
	}

	result := procInData(newTask, c, chat_id)

	c.HTML(http.StatusOK, "index.html", gin.H{"result": result})
}

func procInData(newTask Task, c *gin.Context, chat_id string) float64 {
	var figure mathpkg.Geometry
	var Result float64

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
		return 0
	}

	// chat_id := c.Request.Header.Get("chat_id")
	// if chat_id == "" {
	// 	c.JSON(http.StatusBadRequest, "chat_id не указан")
	// 	return
	// }

	// ch := make(chan float64)

	go func(chat_id string, figure mathpkg.Geometry, figureName string) {
		Result = mathpkg.Measure(figure)

		telegrampkg.Send(chat_id, fmt.Sprintf("Result for figure %s: %f", figureName, Result))

		// ch <- result
		// close(ch)
	}(chat_id, figure, newTask.Figure)

	// result <- ch

	//fmt.Println("result = ", Result)
	return Result
}

func (h *Handler) Posting(c *gin.Context) {
	var newTask Task

	err := c.BindJSON(&newTask)
	if err != nil {
		return
	}

	chat_id := c.Request.Header.Get("chat_id")
	if chat_id == "" {
		c.JSON(http.StatusBadRequest, "chat_id не указан")
		return
	}

	procInData(newTask, c, chat_id)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
