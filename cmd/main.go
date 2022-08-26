package main

import (
	"github.com/AlekseySauron/figures/app"
)

func main() {
	application := app.NewApplication()
	application.Run()
	defer application.Stop()
}
