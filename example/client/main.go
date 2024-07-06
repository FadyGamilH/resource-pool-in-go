package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	request := Request{
		Name: "John",
		Age:  30,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	resp, err := http.Post("http://localhost:9080/greet", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)
	}

	fmt.Println("Server response:", response.Message)
}
