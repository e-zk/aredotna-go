package main

import (
	"log"
	"os"

	arena "aredotna"
)

func main() {
	//chUrl := "https://api.are.na/v2/channels/urandom"
	//r, err := http.Get(chUrl)
	//if err != nil {
	//	log.Fatal(err)
	//}

	a := arena.New(os.Getenv("ARENA_TOKEN"))
	//ch, err := a.GetChannel("cyber-yxyjp1vymh8")
	block, err := a.GetBlock("25635319")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d => %q (%q)", block.Id, block.Image.Original.Url, block.Image.Filename)
	//log.Printf("name: %q", ch.Title)
}
