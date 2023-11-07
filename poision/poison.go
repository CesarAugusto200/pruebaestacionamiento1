package Poison

import (
	"Simulator/models"
	"math/rand"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type CoordenadasFyne struct {
    X float32
    Y float32
}

type CajonesEstacionamiento []bool


func GenerateCarros(coordenadasFyne []CoordenadasFyne, estacionamiento *models.Estacionamiento, cajones CajonesEstacionamiento) {
    if estacionamiento.VehiculosEnEstacionamiento >= estacionamiento.Capacidad {
        return
    }

    <-estacionamiento.AccesoDisponible

    totalCarros := len(coordenadasFyne)
    for i := 0; i < 100; i++ {
        coords := coordenadasFyne[i%totalCarros]

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

        // Verifica si hay cajones de estacionamiento disponibles
        cajonDisponible := -1
        for j, ocupado := range cajones {
            if !ocupado {
                cajonDisponible = j
                break
            }
        }

        if cajonDisponible != -1 {
            // si encuentra un cajon disponible lo ocupa
            cajones[cajonDisponible] = true

            // El vehículo ocupara el cajan durante un tiempo aleatorio entre 1 y 5 segundos
            tiempoOcupacion := rand.Intn(5-1+1) + 1
            time.Sleep(time.Duration(tiempoOcupacion) * time.Second)

            // Desocupa el cajón
            cajones[cajonDisponible] = false
        } else {
			tiempoEspera := rand.Intn(5-1+1) + 1 
			time.Sleep(time.Duration(tiempoEspera) * time.Second)
        }

        estacionamiento.VehiculosEnEstacionamiento++
    }

    estacionamiento.AccesoDisponible <- struct{}{}
}

