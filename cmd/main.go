package main

func main() {
	application := app.NewApplication(ctx)
	application.Run()
	defer application.Stop()
}
