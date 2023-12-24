package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex // pointer

func main() {
	websitelist := []string{
		"https://massivetechinterview.blogspot.com/2015/06/how-to-ace-systems-design-interview.html",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait() 
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	start := time.Now().UnixNano() / int64(time.Millisecond)
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Print("Problem with url")
		fmt.Println(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d statsdue code for %s\n", res.StatusCode, endpoint)
	}

	end := time.Now().UnixNano() / int64(time.Millisecond)
	diff := end - start
	fmt.Printf("Duration(ms): %d", diff)
	fmt.Println("\n")

}
