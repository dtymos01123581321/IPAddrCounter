package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

const maxIPv4 = 4294967295

func main() {
	start := time.Now()

	bitmap := make([]bool, maxIPv4+1)
	file, err := os.Open("ip_addresses.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	uniqueCount := 0

	for scanner.Scan() {
		ipStr := scanner.Text()
		ip := net.ParseIP(ipStr).To4()
		if ip == nil {
			continue
		}
		ipInt := ipToInt(ip)
		if !bitmap[ipInt] {
			bitmap[ipInt] = true
			uniqueCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Number of unique IPs: %d\n", uniqueCount)
	fmt.Printf("Time taken: %s\n", time.Since(start))
}

func ipToInt(ip net.IP) int {
	return int(ip[0])<<24 + int(ip[1])<<16 + int(ip[2])<<8 + int(ip[3])
}
