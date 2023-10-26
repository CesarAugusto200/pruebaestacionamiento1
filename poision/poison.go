package Poison

import (
	"jhsgja/models"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

func GenerateCarros(n int, estacionamiento *models.Estacionamiento) {
	estacionamiento.SlotsEstacionamiento <- true
	for i := 0; i < n; i++ {
		carroImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/carro.png") )
		carroImage.Resize(fyne.NewSize(100,100))
		x := rand.Intn(700-100+1) + 1
		carroImage.Move( fyne.NewPos(float32(x), 500) )

		nuevoCarro := models.CreateCarro(estacionamiento, carroImage)
		nuevoCarro.I = i + 1

		estacionamiento.PintarCarro <- carroImage
		go nuevoCarro.RunCarro()
		TiempoEsperar := rand.Intn(700-100+1) + 1
		time.Sleep(time.Duration(TiempoEsperar) * time.Millisecond)
	}
}
