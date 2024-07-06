package main

import (
	"flag"
	"gopool/server"
	"log"
)

func main() {
	log.Println("server is running")

	var srvUrl string

	flag.StringVar(&srvUrl, "server", ":8080", "server host:port")
	flag.Parse()

	srv := server.NewTcpSrv(srvUrl)

	err := srv.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
