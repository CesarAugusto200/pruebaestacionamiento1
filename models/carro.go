package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Carro struct {
	Estacionamiento *Estacionamiento
	I               int
}

func CreateCarro(e *Estacionamiento) *Carro {
	return &Carro{
		Estacionamiento: e,
	}
}

func (c *Carro) RunCarro() {

	c.Estacionamiento.M.Lock()
	if c.Estacionamiento.SlotsDisponibles <= 0 {
		c.Estacionamiento.M.Unlock()
		fmt.Println("Carro ", c.I, " Esperando")
		c.Estacionamiento.SlotsEstacionamiento <- true
	} else {
		c.Estacionamiento.M.Unlock()
	}

	// Meter al carro en el estacionamiento
	c.Estacionamiento.M.Lock()
	fmt.Println("Carro ", c.I, " Entra")
	c.Estacionamiento.SlotsDisponibles--
	c.Estacionamiento.M.Unlock()
	//Fyne Moverlo

	TiempoEsperar := rand.Intn(5-1+1) + 1
	time.Sleep(time.Duration(TiempoEsperar) * time.Second)

	c.Estacionamiento.M.Lock()
	fmt.Println("Carro ", c.I, " Sale")
	c.Estacionamiento.SlotsDisponibles++
	c.Estacionamiento.M.Unlock()
}
