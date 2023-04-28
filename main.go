package main

import (
	"fmt"
	"sync"
)


func main()  {
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	var score = []int{0}
	wg.Add(3)
	go func (wg *sync.WaitGroup, mut *sync.RWMutex)  {
		mut.Lock()
		fmt.Println("First")
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	go func (wg *sync.WaitGroup,mut *sync.RWMutex)  {
		mut.Lock()
		fmt.Println("second")
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	go func (wg *sync.WaitGroup,mut *sync.RWMutex)  {
		fmt.Println("third")
		mut.RLock()
		score = append(score, 3)
		mut.RUnlock()
		wg.Done()
	}(wg,mut)
	wg.Wait()
	fmt.Println(score)



}