package main

import (
	"log"
	"os"
	"image/png"

	"./identicon"
)

func main() {
	text := "nguyenthanh1290@gmail.com"
	icon, err := identicon.Identicon(text)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("identicon.png")
	if err != nil {
		log.Fatalf("Fail to create identicon.png: %v", err)
	}

	if err := png.Encode(file, icon); err != nil {
		file.Close()
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("Identicon successfully created.")
}
