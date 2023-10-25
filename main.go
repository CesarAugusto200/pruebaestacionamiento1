package main

import (
	"jhsgja/models"
	Poison "jhsgja/poision"
	"time"
)

func main() {
	p := models.CreateEstacionamiento(3)
	Poison.GenerateCarros(5, p)

	time.Sleep(10 * time.Second)
}
