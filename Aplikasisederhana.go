package main

import (
	"fmt"
	"time"
)

// ===============================
// Binary Search Iteratif
// ===============================
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

// ===============================
// Binary Search Rekursif
// ===============================
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

// ===============================
// MAIN PROGRAM (CMD APP)
// ===============================
func main() {

	// DATA CONTOH
	// Jumlah data bisa diubah: 20 / 100 / 1000
	var data []int
	for i := 0; i < 100; i++ {
		data = append(data, 23100001+i)
	}

	var target int
	fmt.Println("=== Aplikasi Binary Search (Go) ===")
	fmt.Println("Jumlah data:", len(data))
	fmt.Print("Masukkan NIM yang dicari: ")
	fmt.Scan(&target)

	// -------------------------------
	// Iteratif
	// -------------------------------
	startIter := time.Now()
	indexIter := binarySearchIterative(data, target)
	elapsedIter := time.Since(startIter)

	// -------------------------------
	// Rekursif
	// -------------------------------
	startRec := time.Now()
	indexRec := binarySearchRecursive(data, 0, len(data)-1, target)
	elapsedRec := time.Since(startRec)

	// -------------------------------
	// OUTPUT
	// -------------------------------
	fmt.Println("\n=== HASIL PENCARIAN ===")

	if indexIter != -1 {
		fmt.Printf("Iteratif  -> Index: %d | Waktu: %d mikrodetik\n",
			indexIter, elapsedIter.Microseconds())
	} else {
		fmt.Println("Iteratif  -> Data tidak ditemukan")
	}

	if indexRec != -1 {
		fmt.Printf("Rekursif -> Index: %d | Waktu: %d mikrodetik\n",
			indexRec, elapsedRec.Microseconds())
	} else {
		fmt.Println("Rekursif -> Data tidak ditemukan")
	}
}
