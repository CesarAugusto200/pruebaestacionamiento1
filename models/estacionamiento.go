package models

import (
	"sync"
	"fyne.io/fyne/v2/canvas"
)

type Estacionamiento struct {
	SlotsEstacionamiento chan bool
	PintarCarro chan *canvas.Image
	VehiculosBloqueados chan struct{}
	CarroAbandona chan struct{}
	M sync.Mutex
}

func CreateEstacionamiento(nS int) *Estacionamiento {
	return &Estacionamiento{
		SlotsEstacionamiento: make(chan bool, 100),
		PintarCarro: make(chan *canvas.Image, 100),
		VehiculosBloqueados: make(chan struct{}),
		CarroAbandona: make(chan struct{}), 
	}
}
