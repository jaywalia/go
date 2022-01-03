package main

import (
	"fmt"
	"sync"
)

func printArray(arr []int) {
	fmt.Print("[")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d,", arr[i])
	}
	fmt.Println("\b]")
}

func duplicateZeroes(arr []int) {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	ch := make(chan int, 2)
	l := len(arr)
	fmt.Println("Duplicate zeroes")
	// array to channel: 2 values
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < l; i += 2 {
			mutex.Lock()
			// fmt.Print("read :")
			// printArray(arr)
			a := arr[i]
			b := 0
			if i+1 < l {
				b = arr[i+1]
			}
			mutex.Unlock()
			// fmt.Printf("i: %d | i+1: %d\n", i, i+1)
			// fmt.Printf("ch <- %d, %d\n", a, b)
			ch <- a
			ch <- b
		}
	}()

	// channel to array: 1 value
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < l; j++ {
			v := <-ch
			// fmt.Printf("%d <- ch\n", v)
			mutex.Lock()
			arr[j] = v
			mutex.Unlock()
			if v == 0 {
				// throw away one value from ch
				// <-ch
				// fill array with an extra 0
				j++
				mutex.Lock()
				arr[j] = 0
				// fmt.Print("Write: ")
				// printArray(arr)
				mutex.Unlock()
			}
		}
		close(ch)
	}()

	wg.Wait()
}

func main() {
	arr := []int{1, 2, 0, 3, 4, 5, 6}
	printArray(arr)
	duplicateZeroes(arr)
	// fmt.Scanln()
	// fmt.Println("done")
	printArray(arr)
}
