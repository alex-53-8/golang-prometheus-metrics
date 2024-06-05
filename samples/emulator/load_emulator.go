package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"time"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8081/processing/create", nil)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	for {
		_, err = client.Do(req)
		if err != nil {
			// handle error
			fmt.Printf(err.Error())
		}

		var awaitBeforeSendingMs = randRange(10, 1600)
		var awaitBeforeSending = time.Duration(awaitBeforeSendingMs) * time.Millisecond
		time.Sleep(awaitBeforeSending)
	}
}
