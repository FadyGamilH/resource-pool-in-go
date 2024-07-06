package server

import (
	"context"
	"fmt"
	"gopool/client"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

type TcpSrv struct {
	NumOfRequests int64
	Router        *gin.Engine
	Server        *http.Server
}

func NewTcpSrv(addr string) *TcpSrv {
	r := gin.Default()
	ts := &TcpSrv{
		NumOfRequests: 0,
		Router:        r,
		Server: &http.Server{
			Addr:         addr,
			Handler:      r,
			WriteTimeout: 1000 * time.Millisecond, // if
			ReadTimeout:  1000 * time.Millisecond,
		},
	}

	ts.Router.POST("/", func(ctx *gin.Context) {
		req := &client.Request{}
		err := ctx.ShouldBind(req)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "couldn't bind the request body",
			})
			return
		}
		log.Println(req.Id)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprint("processed request ", req.Id),
		})

		// add the number of requests we served by 1
		atomic.AddInt64(&ts.NumOfRequests, 1)

		time.Sleep(500 * time.Millisecond)
	})

	return ts
}

func (ts *TcpSrv) Start() error {
	log.Println("server starting")
	var err error
	// notic ethat this method is a blocking method so if you need to run the start() which calls this LIstenAndServe method, you must run it in a go routine or all the code after the start method will not get executed unless the server encounters an error
	go func() {
		err = ts.Server.ListenAndServe()
	}()

	time.Sleep(time.Duration(time.Millisecond * 400))

	return err
}

func (ts *TcpSrv) StopGracefully() error {
	// if its closed
	if ts.Server == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("server shutting down number of served requests is : ", ts.NumOfRequests)

	return ts.Server.Shutdown(ctx)
}

func (ts *TcpSrv) Stop() error {
	// if its closed
	if ts.Server == nil {
		return nil
	}

	log.Println("server shutting down, number of served requests is : ", ts.NumOfRequests)

	return ts.Server.Close()
}
