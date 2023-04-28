package main

import (
	"fmt"
	"sync"
)



func main()  {
	fmt.Println("Go Channels")

	mych := make(chan int,2)
	wg := &sync.WaitGroup{}
	

	wg.Add(2)
	//receive only
	go func(ch <-chan int, wg *sync.WaitGroup){
		val,ischanopen:= <-ch
		
		fmt.Println(val)
		fmt.Println(ischanopen)
		// fmt.Println(<-mych)
		wg.Done()	
	}(mych,wg)
	//send Only
	go func(ch chan<- int, wg *sync.WaitGroup){
		
		// mych <- 0
		mych <- 6
		close(mych)
		wg.Done()	
	}(mych,wg)
	wg.Wait()	





	
}



