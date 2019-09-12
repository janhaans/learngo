package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

//DoWork waits for some random time
func DoWork(id int) {
	defer wg.Done()
	fmt.Printf("Worker %d has started\n", id)

	time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)

	fmt.Printf("Worker %d has ended\n", id)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	nWorkers := 4
	wg.Add(nWorkers)
	fmt.Println("The workers are starting")

	for i := 0; i < nWorkers; i++ {
		go DoWork(i)
	}

	wg.Wait()
	fmt.Println("The workers are finished")
}
