package main

import (
	"log"

	"github.com/gen2brain/beeep"
)

func main() {
	err := beeep.Notify("Judul Notif", "Pesan notifikasi dari Go", "indo.png")
	if err != nil {
		log.Fatal(err)
	}
}
