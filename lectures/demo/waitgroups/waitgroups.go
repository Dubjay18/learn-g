package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a new WaitGroup
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		counter += 1
		go func() {
			defer func() {

				fmt.Println(counter, "goroutines remaining")
				counter -= 1
				wg.Done()
			}()
			duration := time.Duration(rand.Intn(1000)) * time.Millisecond
			fmt.Println("Sleeping for", duration)
			time.Sleep(duration)
		}()
		fmt.Println("waiting for goroutines")
		wg.Wait()
		fmt.Println("done!")
	}
}
