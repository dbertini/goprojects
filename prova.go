package main

import (
	"fmt"
	"log"
)

func main() {

	log.Default().Printf("prova")
	pippo()
	var variabile = "Prova Var"
	funzione("Diavolone!")
	funzione(variabile)
}

func pippo() {
	fmt.Println("Prova")
}

func funzione(arg string) {
	fmt.Println(arg)
}
