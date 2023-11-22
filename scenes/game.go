package scenes

import (
	"Simulator/models"
	Poison "Simulator/poision"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type GameScene struct {
	window fyne.Window
	content *fyne.Container
}

func NewGameScene(window fyne.Window) *GameScene {
	scene := &GameScene{window: window}
    scene.Render()
    return scene
}

func (s *GameScene) Render() {
	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/Group2.png") )
    backgroundImage.Resize(fyne.NewSize(800,600))
	backgroundImage.Move( fyne.NewPos(0,0) )

	s.content = container.NewWithoutLayout(
        backgroundImage, 
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
    var mu sync.Mutex

    for {
        imagen := <-e.PintarCarro

    
        mu.Lock()
        s.content.Add(imagen)
        mu.Unlock()

        // Espera a que el carro abandone el estacionamiento
        <-e.CarroAbandona

       
        s.window.Canvas().Refresh(s.content)
    }
}



