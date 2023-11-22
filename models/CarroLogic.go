package models

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
)

func (c *Carro) RunCarro() {
	for {
		c.Estacionamiento.M.Lock()
		select {
		case slot := <-c.Estacionamiento.SlotsEstacionamiento:
			
			x := float32(rand.Intn(650-150+1) + 150)
			y := float32(rand.Intn(300-50+1) + 50)
			c.skin.Move(fyne.NewPos(x, y))
			fmt.Println("Carro", c.I, "Entra al estacionamiento")
			c.Estacionamiento.M.Unlock()

			TiempoEsperar := rand.Intn(5-1+1) + 1
			time.Sleep(time.Duration(TiempoEsperar) * time.Second)

			c.Estacionamiento.M.Lock()
			c.skin.Move(fyne.NewPos(0, 0))
			fmt.Println("Carro", c.I, "Abandona el estacionamiento")
			c.Estacionamiento.CarroAbandona <- struct{}{}
			c.Estacionamiento.SlotsEstacionamiento <- slot
			c.Estacionamiento.M.Unlock()
		default:
			fmt.Println("Carro", c.I, "Intenta buscar el siguiente cajón disponible")
			c.Estacionamiento.M.Unlock()
			time.Sleep(time.Millisecond * 100)
		}
	}
}
