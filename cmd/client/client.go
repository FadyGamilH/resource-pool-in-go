package main

import (
	"flag"
	"gopool/client"
	"log"
)

const (
	srv = ":8080"
)

func main() {
	log.Println("client start running ... ")

	var srvUrl string

	flag.StringVar(&srvUrl, "server", srv, "server host:port")
	flag.Parse()

	client.SendRequestsInBatches(srvUrl)
}
