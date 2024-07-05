package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	s = rand.NewSource(time.Now().UnixMilli())
	r = rand.New(s)
)

const (
	AddReq = iota
	AvgReq
	RandomReq
	SpellCheckReq
	SearchReq
)

const (
	maxNumberOfRequestsPerClient = 100
	maxBatchSize                 = (maxNumberOfRequestsPerClient / 10) * 2 // 20 percent of the max number of requests per client will be the max number of msgs per batch, we might send less than that of course
)

const MaxReqDataSize = 1 * 1024 // kb of data per request is the max request size

type Request struct {
	Id   int
	Type int
	Data [MaxReqDataSize]byte // slice of 1 KB of bytes
	Size int                  // the size of our actual data that we sent of the Data [] we might send less than 1KB
}

func EncodeStructToJson(req *Request) io.Reader {
	buf := &bytes.Buffer{}
	// create an encoder initialized by this buffer
	enc := json.NewEncoder(buf)
	enc.Encode(req) // turn req to json and write it into the buf

	// we can read from the buffer as following
	io.Copy(os.Stdout, buf) // --> this swill prints {"Id" : 1}

	return buf
}

func SendRequestsInBatches(url string) {
	reqId := 0

	msgsLeftToSent := maxNumberOfRequestsPerClient

	// as long as we still have something to send
	for msgsLeftToSent > 0 {

		batch := r.Intn(maxBatchSize) // random number between 0 -> 20 % of msgs left to sent
		if batch <= msgsLeftToSent {
			batch = msgsLeftToSent
		}

		for i := 0; i < batch; i++ {
			req := &Request{}
			reqId++
			req.Id = reqId
			req.Size = r.Intn(MaxReqDataSize)
			req.Type = r.Intn(5)

			buf := EncodeStructToJson(req)

			res, err := http.Post(url, "text/json", buf)
			if err != nil {
				log.Fatalln(err)
			}
			defer res.Body.Close()
		}

		msgsLeftToSent = msgsLeftToSent - batch
	}

	// sleep randomly between each batch
	time.Sleep(time.Duration(r.Intn(500) * int(time.Millisecond)))
}
