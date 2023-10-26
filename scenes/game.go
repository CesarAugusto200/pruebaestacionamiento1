package scenes

import (
	"jhsgja/models"
	Poison "jhsgja/poision"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type GameScene struct {
	window fyne.Window
	content        *fyne.Container
}

func NewGameScene(window fyne.Window) *GameScene {
	scene := &GameScene{window: window}
    scene.Render()
    return scene
}

func (s *GameScene) Render() {
	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/bg.png") )
    backgroundImage.Resize(fyne.NewSize(800,600))
	backgroundImage.Move( fyne.NewPos(0,0) )

	s.content = container.NewWithoutLayout(
        backgroundImage, // Fondo
    )
    s.window.SetContent(s.content) 
    s.StartGame()
}

func (s *GameScene) StartGame() {
	e := models.CreateEstacionamiento(20)
	go Poison.GenerateCarros(100, e)
	go s.PintarCarros(e)
}

func (s *GameScene) PintarCarros(e *models.Estacionamiento) {
	for {
		imagen := <- e.PintarCarro
		s.content.Add(imagen)
        s.window.Canvas().Refresh(s.content)
	}
}

func createPeel( fileUri string, posX float32, posY float32 ) *canvas.Image {
	image := canvas.NewImageFromURI(storage.NewFileURI(fileUri))
	image.Resize(fyne.NewSize(50,50))
	image.Move(fyne.NewPos(posX,posY))
	return image
}

