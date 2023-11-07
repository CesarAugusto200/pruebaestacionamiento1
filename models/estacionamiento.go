package models

import (
	"sync"
	"fyne.io/fyne/v2/canvas"
)

type Estacionamiento struct {
	SlotsEstacionamiento chan bool
	PintarCarro chan *canvas.Image
	M sync.Mutex
	VehiculosEnEstacionamiento int
	Capacidad int 
	EspaciosDisponibles chan struct{}
	PuertaAccesoOcupada bool
	AccesoDisponible chan struct{}
}

func CreateEstacionamiento(nS int) *Estacionamiento {
	return &Estacionamiento{
		Capacidad: nS,
		EspaciosDisponibles: make(chan struct{}, nS),
		SlotsEstacionamiento: make(chan bool, nS),
		PintarCarro:           make(chan *canvas.Image, 1),
		VehiculosEnEstacionamiento: 0,
	}
}
