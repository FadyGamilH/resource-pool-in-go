package server

import (
	"context"
	"fmt"
	"gopool/client"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TcpSrv struct {
	NumOfRequests int
	Router        *gin.Engine
	Server        *http.Server
}

func NewTcpSrv(addr string) *TcpSrv {
	r := gin.Default()
	return &TcpSrv{
		NumOfRequests: 0,
		Router:        r,
		Server: &http.Server{
			Addr:         addr,
			Handler:      r,
			WriteTimeout: 1000 * time.Millisecond, // if
			ReadTimeout:  1000 * time.Millisecond,
		},
	}
}

func (ts *TcpSrv) Start() error {

	ts.Router.POST("/", func(ctx *gin.Context) {
		req := &client.Request{}
		err := ctx.ShouldBind(req)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "couldn't bind the request body",
			})
			time.Sleep(500 * time.Millisecond)
			return
		}
		log.Println(req.Id)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprint("processed request ", req.Id),
		})
		time.Sleep(500 * time.Millisecond)
	})

	log.Println("server starting")

	return ts.Server.ListenAndServe()
}

func (ts *TcpSrv) StopGracefully() error {
	// if its closed
	if ts.Server == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("server shutting down")

	return ts.Server.Shutdown(ctx)
}

func (ts *TcpSrv) Stop() error {
	// if its closed
	if ts.Server == nil {
		return nil
	}

	log.Println("server shutting down")

	return ts.Server.Close()
}
