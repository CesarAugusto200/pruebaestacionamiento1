package models

/*import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"math/rand"
	"sync"
	"time"

 )

 
 var entranceSemaphore = &sync.Mutex{}

 func GenerateCarsContinuously(estacionamiento , *models.Estacionamiento) {
    generatedCars := 0
    for {
        select {
        case estacionamiento.SlotsEstacionamiento <- true:
         
            entranceSemaphore.Lock()
         
            entranceSemaphore.Unlock()
        default:
            
            <-estacionamiento.VehiculosBloqueados
            continue 
        }

        carroImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/Car.png"))
        carroImage.Resize(fyne.NewSize(100, 100))
        x := float32(rand.Intn(700-100+1) + 1)
        carroImage.Move(fyne.NewPos(x, 500))

        nuevoCarro := models.CreateCarro(estacionamiento, carroImage)
        nuevoCarro.I = generatedCars + 1

        estacionamiento.PintarCarro <- carroImage
        go nuevoCarro.RunCarro()
        TiempoEsperar := rand.Intn(700-100+1) + 1
        time.Sleep(time.Duration(TiempoEsperar) * time.Millisecond)

        generatedCars++
        if generatedCars >= 100 {
            break
        }
    }
}
*/