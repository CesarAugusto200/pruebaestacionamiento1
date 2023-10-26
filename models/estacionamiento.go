package models

import (
	"sync"

	"fyne.io/fyne/v2/canvas"
)

type Estacionamiento struct {
	SlotsDisponibles     int
	SlotsEstacionamiento chan bool
	PintarCarro chan *canvas.Image

	M sync.Mutex
}

func CreateEstacionamiento(nS int) *Estacionamiento {
	return &Estacionamiento{
		SlotsDisponibles:     nS,
		SlotsEstacionamiento: make(chan bool, 1),
		PintarCarro: make(chan *canvas.Image, 1),
	}
}
