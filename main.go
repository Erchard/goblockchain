package main

import (
	"./mining"
	"fmt"
	"log"
)

func main() {
	log.Println("Go Blockchain")
	log.Println("Author: arsenguzhva@gmail.com")

	ending := make(chan bool, 1)

	go mining.MineLoop(ending)

	var msg string

	_, _ = fmt.Scanln(&msg)
	ending <- true
	log.Println("Stop system...")

	_, _ = fmt.Scanln(&msg)

}
