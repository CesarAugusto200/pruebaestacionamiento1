package models

import (
	"sync"

	"fyne.io/fyne/v2/canvas"
)

type Estacionamiento struct {
	SlotsEstacionamiento chan bool
	PintarCarro chan *canvas.Image

	M sync.Mutex
}

func CreateEstacionamiento(nS int) *Estacionamiento {
	return &Estacionamiento{
		SlotsEstacionamiento: make(chan bool, nS+1),
		PintarCarro: make(chan *canvas.Image, 1),
	}
}
