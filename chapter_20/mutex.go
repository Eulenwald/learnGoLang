package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mu sync.Mutex

func mutex() {
	fmt.Println("Kapitel 20 Mutex")

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines before start:", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// time.Sleep(time.Second * 1)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Printf("%v ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("\nGoroutines after running:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
