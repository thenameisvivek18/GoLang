package main

import (
	"fmt"
	"net/http"
	
	"sync"
)
var signals=[]string{"Test"}
var wg sync.WaitGroup
var mut sync.Mutex
func main() {
	
	ws := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
	}

	for _, web := range ws {
		go getStatusCode(web)
		// go sum()
		// go sub()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS is Endpoints")
	} else {
		mut.Lock()
		signals = append(signals,endpoint)
		mut.Unlock()
		fmt.Printf("%d Status Code For %s\n", res.StatusCode, endpoint)
	}
}

// func sum() {
// 	defer wg.Done()	
// 	fmt.Println("Sum is", 2+2)
// }
// func sub() {
// 	defer wg.Done()
// 	fmt.Println("Sub is", 4-2)
// }
