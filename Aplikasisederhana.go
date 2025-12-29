package main

import (
	"fmt"
	"time"
)

const ITERASI = 100000

// Binary Search Iteratif
func binarySearchIterative(data []int, target int) int {
	low, high := 0, len(data)-1

	for low <= high {
		mid := (low + high) / 2
		if data[mid] == target {
			return mid
		} else if data[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Binary Search Rekursif
func binarySearchRecursive(data []int, low, high, target int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if data[mid] == target {
		return mid
	} else if data[mid] < target {
		return binarySearchRecursive(data, mid+1, high, target)
	} else {
		return binarySearchRecursive(data, low, mid-1, target)
	}
}

// MAIN PROGRAM
func main() {

	var N int
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&N)

	data := make([]int, N)

	fmt.Println("Masukkan data NIM (satu per baris, HARUS terurut):")
	for i := 0; i < N; i++ {
		fmt.Printf("Data ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	var target int
	fmt.Print("Masukkan NIM yang dicari: ")
	fmt.Scan(&target)

	// Iteratif
	startIter := time.Now()
	var indexIter int
	for i := 0; i < ITERASI; i++ {
		indexIter = binarySearchIterative(data, target)
	}
	elapsedIter := time.Since(startIter)

	// Rekursif
	startRec := time.Now()
	var indexRec int
	for i := 0; i < ITERASI; i++ {
		indexRec = binarySearchRecursive(data, 0, len(data)-1, target)
	}
	elapsedRec := time.Since(startRec)

	// OUTPUT
	fmt.Println("\n=== HASIL PENCARIAN ===")
	fmt.Println("Jumlah Data:", N)
	fmt.Println("Jumlah Iterasi:", ITERASI)

	if indexIter != -1 {
		fmt.Printf("Iteratif  -> Index: %d | Waktu: %d mikrodetik\n",
			indexIter, elapsedIter.Microseconds())
	} else {
		fmt.Printf("Iteratif  -> Data tidak ditemukan | Waktu: %d mikrodetik\n",
			elapsedIter.Microseconds())
	}

	if indexRec != -1 {
		fmt.Printf("Rekursif -> Index: %d | Waktu: %d mikrodetik\n",
			indexRec, elapsedRec.Microseconds())
	} else {
		fmt.Printf("Rekursif -> Data tidak ditemukan | Waktu: %d mikrodetik\n",
			elapsedRec.Microseconds())
	}
}
