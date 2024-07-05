package main

import (
	"gopool/client"
	"log"
)

func main() {
	log.Println("client start running ... ")

	client.SendRequestsInBatches()
}
