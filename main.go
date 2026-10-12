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

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)

	for i := range data {
		data[i] = rand.Intn(1_000_000) + 1
	}

	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]

	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks returns the maximum number of elements in chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup
	maxValues := make([]int, CHUNKS)

	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		if start >= len(data) {
			break
		}

		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}

		wg.Add(1)

		go func(index int, chunk []int) {
			defer wg.Done()
			maxValues[index] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf(
		"Максимальное значение элемента: %d\nВремя поиска: %d мкс\n",
		maxValue, elapsed,
	)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf(
		"Максимальное значение элемента: %d\nВремя поиска: %d мкс\n",
		maxValue, elapsed,
	)
}
