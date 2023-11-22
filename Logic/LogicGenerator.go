package logic

import (
	"Simulator/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"math/rand"
	"sync"
	"time"
)

var entranceSemaphore = &sync.Mutex{}


func GenerateCars(n int, estacionamiento *models.Estacionamiento) {
	generatedCars := 0

	
	espaciosEstacionamiento := []fyne.Position{
		fyne.NewPos(100, 100), fyne.NewPos(200, 100), fyne.NewPos(300, 100),
		fyne.NewPos(400, 100), fyne.NewPos(500, 100), fyne.NewPos(600, 100),
		fyne.NewPos(100, 200), fyne.NewPos(200, 200), fyne.NewPos(300, 200),
		fyne.NewPos(400, 200), fyne.NewPos(500, 200), fyne.NewPos(600, 200),
		fyne.NewPos(100, 300), fyne.NewPos(200, 300), fyne.NewPos(300, 300),
		fyne.NewPos(400, 300), fyne.NewPos(500, 300), fyne.NewPos(600, 300),
		fyne.NewPos(100, 400), fyne.NewPos(200, 400), fyne.NewPos(300, 400),
	}

	for i := 0; i < n; i++ {
		select {
		case estacionamiento.SlotsEstacionamiento <- true:
			entranceSemaphore.Lock()
			entranceSemaphore.Unlock()

			// Crear una nueva imagen de carro y ajustar su tamaño
			carroImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/Car.png"))
			carroImage.Resize(fyne.NewSize(100, 100))

			// Mover la imagen del carro a un espacio de estacionamiento
			carroImage.Move(espaciosEstacionamiento[i%len(espaciosEstacionamiento)])

			// Crear un nuevo carro y asignarle un número
			nuevoCarro := models.CreateCarro(estacionamiento, carroImage)
			nuevoCarro.I = generatedCars + 1

			// Pintar la imagen del carro en el estacionamiento
			estacionamiento.PintarCarro <- carroImage
			go nuevoCarro.RunCarro()

			// Tiempo de espera aleatorio antes de generar otro carro
			TiempoEsperar := rand.Intn(5000-1000+1) + 1000
			time.Sleep(time.Duration(TiempoEsperar) * time.Millisecond)

			generatedCars++
		default:
			// Si no hay espacio disponible, esperar y continuar
			<-estacionamiento.VehiculosBloqueados
			continue
		}
	}
}
