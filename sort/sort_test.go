package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	var a []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(100))
	}
	b := make([]int, len(a))
	c := make([]int, len(a))
	d := make([]int, len(a))
	e := make([]int, len(a))
	copy(b, a)
	copy(c, a)
	copy(d, a)
	copy(e, a)
	fmt.Println("a", a, len(a), cap(a))
	fmt.Println("b", b, len(b), cap(b))
	fmt.Println("c", c, len(c), cap(c))
	fmt.Println("d", d, len(d), cap(d))
	BubbleSort(a)
	SelectSort(b)
	InsertionSort(c)
	QuickSort(d)
	e = MergeSort(e)
	fmt.Println("BubbleSort:", a)
	fmt.Println("SelectSort:", b)
	fmt.Println("InsertionSort:", c)
	fmt.Println("QuickSort:", d)
	fmt.Println("Mergesort:", e)
}

func BenchmarkC(b *testing.B) {
	b.ResetTimer()
	var q [][]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		var a []int
		for i := 0; i < 20; i++ {
			a = append(a, rand.Intn(100))
		}
		q = append(q, a)

	}
	for i := 0; i < b.N; i++ {
		InsertionSort(q[i])
	}
}
