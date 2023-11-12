package Poison

import (
    "Simulator/models"
    "math/rand"
    "time"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/storage"
    "sync"
)

var entranceSemaphore = &sync.Mutex{}

func GenerateCarros(n int, estacionamiento *models.Estacionamiento) {
    generatedCars := 0

    //  coordenadas de los espacios de estacionamiento
    espaciosEstacionamiento := []fyne.Position{
        fyne.NewPos(100, 100),
        fyne.NewPos(200, 100),
        fyne.NewPos(300, 100),
        fyne.NewPos(400, 100),
        fyne.NewPos(500, 100),
        fyne.NewPos(600, 100),
        fyne.NewPos(100, 200),
        fyne.NewPos(200, 200),
        fyne.NewPos(300, 200),
        fyne.NewPos(400, 200),
        fyne.NewPos(500, 200),
        fyne.NewPos(600, 200),
        fyne.NewPos(100, 300),
        fyne.NewPos(200, 300),
        fyne.NewPos(300, 300),
        fyne.NewPos(400, 300),
        fyne.NewPos(500, 300),
        fyne.NewPos(600, 300),
        fyne.NewPos(100, 400),
        fyne.NewPos(200, 400),
        fyne.NewPos(300, 400),
    }

    for i := 0; i < n; i++ {
        select {
        case estacionamiento.SlotsEstacionamiento <- true:
            entranceSemaphore.Lock()
          
            entranceSemaphore.Unlock()

        
            carroImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/Car.png"))
            carroImage.Resize(fyne.NewSize(100, 100))

           
            carroImage.Move(espaciosEstacionamiento[i % len(espaciosEstacionamiento)])

            nuevoCarro := models.CreateCarro(estacionamiento, carroImage)
            nuevoCarro.I = generatedCars + 1

            estacionamiento.PintarCarro <- carroImage
            go nuevoCarro.RunCarro()

            TiempoEsperar := rand.Intn(5000-1000+1) + 1000
            time.Sleep(time.Duration(TiempoEsperar) * time.Millisecond)

            generatedCars++
        default:
          
            <-estacionamiento.VehiculosBloqueados
            continue 
        }
    }
}

func GenerateCarsContinuously(estacionamiento *models.Estacionamiento) {
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
