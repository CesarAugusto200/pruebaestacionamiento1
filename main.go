package main

import (
	"jhsgja/models"
	Poison "jhsgja/poision"
	"time"
)

func main() {
	p := models.CreateEstacionamiento(3)
	Poison.GenerateCarros(10, p)

	time.Sleep(30 * time.Second)
}
