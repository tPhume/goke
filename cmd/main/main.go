package main

import (
	"github.com/tPhume/goke/poke"
	"log"
	"os"
)

func main() {
	if err := poke.RunApp(os.Args); err != nil {
		log.Fatal(err)
	}
}
