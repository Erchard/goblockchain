package main

import (
	"./mining"
	"fmt"
	"log"
)

func main() {
	log.Println("Go Blockchain")
	log.Println("Author: arsenguzhva@gmail.com")

	mining.MineLoop()

	var msg string

	_, _ = fmt.Scanln(&msg)

}
