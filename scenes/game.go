package scenes

import (
    "fmt"
    "Simulator/models"
    Poison "Simulator/poision"
    "math/rand"
    "time"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/storage"
)

type GameScene struct {
    window  fyne.Window
    content *fyne.Container
}

func NewGameScene(window fyne.Window) *GameScene {
    scene := &GameScene{window: window}
    scene.Render()
    return scene
}

func GenerateCarros(coordenadasFyne []Poison.CoordenadasFyne, estacionamiento *models.Estacionamiento) {
    // Verifica la capacidad máxima del estacionamiento
    if estacionamiento.VehiculosEnEstacionamiento >= estacionamiento.Capacidad {
        return
    }

    // Intenta obtener un espacio disponible, bloqueándose si no hay espacios
    estacionamiento.EspaciosDisponibles <- struct{}{}
    defer func() { <-estacionamiento.EspaciosDisponibles }()

    totalCarros := len(coordenadasFyne)
    for i := 0; i < 100; i++ {
        coords := coordenadasFyne[i%totalCarros] // Reutiliza las coordenadas si se agotan

        // Verifica la capacidad máxima del estacionamiento antes de permitir que un vehículo entre
        if estacionamiento.VehiculosEnEstacionamiento >= estacionamiento.Capacidad {
            break
        }

        carroImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/Car.png"))
        carroImage.Resize(fyne.NewSize(100, 100))
        carroImage.Move(fyne.NewPos(coords.X, coords.Y))

        nuevoCarro := models.CreateCarro(estacionamiento, carroImage)
        nuevoCarro.I = i + 1

        estacionamiento.PintarCarro <- carroImage
        go nuevoCarro.RunCarro()
        TiempoEsperar := rand.Intn(700-100+1) + 1
        time.Sleep(time.Duration(TiempoEsperar) * time.Millisecond)

        // Actualiza el registro de vehículos en el estacionamiento
        estacionamiento.VehiculosEnEstacionamiento++
        fmt.Println("Carro ", nuevoCarro.I, " Entra") // Agrega esta línea
    }
}

func (s *GameScene) Render() {
    backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/Group2.png"))
    backgroundImage.Resize(fyne.NewSize(800, 600))
    backgroundImage.Move(fyne.NewPos(0, 0))

    s.content = container.NewWithoutLayout(
        backgroundImage,
    )
    s.window.SetContent(s.content)
    s.StartGame()
}

func (s *GameScene) StartGame() {
	cajones := make(Poison.CajonesEstacionamiento, 20)
    e := models.CreateEstacionamiento(20)

  
    coordenadasFyne := []Poison.CoordenadasFyne{
        {X: 100, Y: 200},
        {X: 300, Y: 400},
       
    }

    go func() {
        for {
            
            <-e.AccesoDisponible

           
            Poison.GenerateCarros(coordenadasFyne, e, cajones)// Pasa coordenadasFyne como argumento

            
            e.AccesoDisponible <- struct{}{}
        }
    }()

    go s.PintarCarros(e)
}

func (s *GameScene) PintarCarros(e *models.Estacionamiento) {
    for {
        imagen := <-e.PintarCarro
        s.content.Add(imagen)
        s.window.Canvas().Refresh(s.content)
    }
}

func createPeel(fileUri string, posX float32, posY float32) *canvas.Image {
    image := canvas.NewImageFromURI(storage.NewFileURI(fileUri))
    image.Resize(fyne.NewSize(50, 50))
    image.Move(fyne.NewPos(posX, posY))
    return image
}
