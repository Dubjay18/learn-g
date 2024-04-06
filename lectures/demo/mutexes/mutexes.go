package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

type Hits struct {
	count int
	sync.Mutex
}

func serve(wg *sync.WaitGroup, h *Hits, iteration int) {
	wait()
	h.Lock()
	defer h.Unlock()
	defer wg.Done()
	h.count++
	fmt.Println("serving", iteration)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var wg sync.WaitGroup
	hitCounter := Hits{}
	for i := 0; i < 20; i++ {
		iteration := i
		wg.Add(1)
		go serve(&wg, &hitCounter, iteration)
	}

	fmt.Printf("waiting for goroutines..\n\n")
	wg.Wait()
	hitCounter.Lock()
	totalHits := hitCounter.count
	hitCounter.Unlock()

	fmt.Println(totalHits)
}
