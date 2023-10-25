package Poison

import (
	"jhsgja/models"
)

func GenerateCarros(n int, estacionamiento *models.Estacionamiento) {
	estacionamiento.SlotsEstacionamiento <- true
	for i := 0; i < n; i++ {
		nuevoCarro := models.CreateCarro(estacionamiento)
		nuevoCarro.I = i + 1
		go nuevoCarro.RunCarro()
	}
}
