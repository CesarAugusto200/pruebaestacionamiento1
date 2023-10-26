package main

import (
	"jhsgja/scenes"


	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tux Revenge")

	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(800, 600))
	scenes.NewGameScene(w)
	w.ShowAndRun()
}
