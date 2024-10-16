package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const maxIPv4_2 = 4294967295
const numWorkers = 4

func main() {
	start := time.Now()

	bitmap := make([]bool, maxIPv4_2+1)

	file, err := os.Open("ip_addresses.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	wg := &sync.WaitGroup{}
	ch := make(chan string, 1000)
	mu := &sync.Mutex{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ch, bitmap, wg, mu)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ipStr := scanner.Text()
		ch <- ipStr
	}

	close(ch)

	wg.Wait()

	uniqueCount := 0
	for _, exists := range bitmap {
		if exists {
			uniqueCount++
		}
	}

	fmt.Printf("Number of unique IPs: %d\n", uniqueCount)
	fmt.Printf("Time taken: %s\n", time.Since(start))
}

func worker(ch <-chan string, bitmap []bool, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	for ipStr := range ch {
		ip := net.ParseIP(ipStr).To4()
		if ip == nil {
			continue
		}
		ipInt := ipToInt2(ip)

		mu.Lock()
		if !bitmap[ipInt] {
			bitmap[ipInt] = true
		}
		mu.Unlock()
	}
}

func ipToInt2(ip net.IP) int {
	return int(ip[0])<<24 + int(ip[1])<<16 + int(ip[2])<<8 + int(ip[3])
}
