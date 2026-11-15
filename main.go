package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)

	for index := range data {
		data[index] = rand.Intn(1_000_000) + 1
	}

	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	result := data[0]

	for _, value := range data[1:] {
		if value > result {
			result = value
		}
	}

	return result
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxi := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	var workers sync.WaitGroup
	workers.Add(CHUNKS)

	for index := 0; index < CHUNKS; index++ {
		start := index * chunkSize
		end := start + chunkSize

		if index == CHUNKS-1 {
			end = len(data)
		}

		go func(chunk []int, resultIndex int) {
			defer workers.Done()

			if len(chunk) == 0 {
				maxi[resultIndex] = data[0]
				return
			}

			maxi[resultIndex] = maximum(chunk)
		}(data[start:end], index)
	}

	workers.Wait()

	return maximum(maxi)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)

	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
