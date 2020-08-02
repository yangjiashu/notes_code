package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	filenames := make(chan string, 10)

	for i := 0; i < 10; i++ {
		filenames <- fmt.Sprintf("file%d", i+1)
	}
	close(filenames)

	result := readFiles(filenames)
	fmt.Println("总共读取的字节数:", result)
}

func readFiles(filenames <-chan string) int {
	var total int
	sizes := make(chan int, 5)
	wg := &sync.WaitGroup{}

	// workers
	for filename := range filenames {
		// 1
		wg.Add(1)
		go func(f string) {
			// 2
			defer wg.Done()
			b, _ := imageFile(f)
			sizes <- len(b)
		}(filename)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	for size := range sizes {
		total += size
	}

	return total
}

func imageFile(filename string) (b []byte, err error) {
	randNum := rand.Intn(10)
	if randNum == 0 {
		err = errors.New("出现错误")
		return
	}
	rand.Seed(2)
	time.Sleep(time.Duration(500*randNum) * time.Millisecond)
	for i := 0; i < randNum; i++ {
		b = append(b, 'b')
	}
	return
}
