package models

import "sync"

type Estacionamiento struct {
	SlotsDisponibles     int
	SlotsEstacionamiento chan bool

	M sync.Mutex
}

func CreateEstacionamiento(nS int) *Estacionamiento {
	return &Estacionamiento{
		SlotsDisponibles:     nS,
		SlotsEstacionamiento: make(chan bool, 1),
	}
}
