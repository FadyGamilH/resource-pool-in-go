package main

import (
	"flag"
	"gopool/server"
	"log"
	"time"
)

func main() {
	log.Println("server is running")

	var srvUrl string

	flag.StringVar(&srvUrl, "server", ":8080", "server host:port")
	flag.Parse()

	srv := server.NewTcpSrv(srvUrl)

	// err := srv.Start()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	StartSleepStop(srv)
}

func StartSleepStop(srv *server.TcpSrv) {

	d := time.Duration(time.Second * 5)

	err := srv.Start()
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(d)

	err = srv.StopGracefully()
	if err != nil {
		log.Println(err)
	}
}
