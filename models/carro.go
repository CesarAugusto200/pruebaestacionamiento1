package models

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Carro struct {
	Estacionamiento *Estacionamiento
	I               int
	skin			*canvas.Image
}

func CreateCarro(e *Estacionamiento, s *canvas.Image) *Carro {
	return &Carro{
		Estacionamiento: e,
		skin: s,
	}
}

func (c *Carro) RunCarro() {

	c.Estacionamiento.M.Lock()
	if c.Estacionamiento.SlotsDisponibles <= 0 {
		c.Estacionamiento.M.Unlock()
		c.Estacionamiento.SlotsEstacionamiento <- true
		c.Estacionamiento.M.Lock()
	} else {
		c.Estacionamiento.SlotsDisponibles--
	}

	// Meter al carro en el estacionamiento

	
	time.Sleep(1 *time.Second)
	x := float32( rand.Intn(650-150+1) )
	y := float32( rand.Intn(300-50+1) )
	c.skin.Move(fyne.NewPos( x, y ))
	fmt.Println("Carro ", c.I, " Entra")

	c.Estacionamiento.M.Unlock()






	

	TiempoEsperar := rand.Intn(5-1+1) + 1
	time.Sleep(time.Duration(TiempoEsperar) * time.Second)

	c.Estacionamiento.M.Lock()
	
	time.Sleep(1 *time.Second)
	fmt.Println("Carro ", c.I, " Sale")
	c.skin.Move(fyne.NewPos( 0,0 ))
	if c.Estacionamiento.SlotsDisponibles == 0 {
		c.Estacionamiento.SlotsDisponibles--
		<-c.Estacionamiento.SlotsEstacionamiento
	}
	
	c.Estacionamiento.SlotsDisponibles++
	c.Estacionamiento.M.Unlock()
}
