package main

import (
	"Algorithm/sort"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	// var a []int
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 10; i++ {
	// 	a = append(a, rand.Intn(100))
	// }
	// b := make([]int, len(a))
	// c := make([]int, len(a))
	// // sort.QuickSort(a)
	// copy(b, a)
	// copy(c, a)
	// fmt.Println("a", a, len(a), cap(a))
	// fmt.Println("b", b, len(b), cap(b))
	// fmt.Println("c", c, len(c), cap(c))
	// sort.SelectSort(a)
	// sort.InsertionSort(b)
	// // sort.QuickSort(c)
	// c = sort.MergeSort(c)
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	fmt.Println("测试结束!")
}

func BenchmarkB(b *testing.B) {
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
		sort.SelectSort(q[i])
	}
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
		sort.InsertionSort(q[i])
	}
}

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		fmt.Println(result, num)
		result ^= num
	}
	return result
}

func singleNumber2(nums []int) (int, int) {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	/*两个数相异或后的数取最低位的1然后再和数组中的数分别做与运算，
	排除所有相与为0的（一种是双次出现的和这个1对应的数字，还有不包括这个1的数字，
	当中有一个是我们要的只出现一次的，其他的因为都是多次出现可以忽略掉了）目的是将2个单次出现的排除掉一个，将数组变成只有一个单多出现的再运算
	*/
	value := result & (^result + 1)
	ww := 0
	for _, num := range nums {
		if value&num != 0 {
			ww ^= num
		}
	}
	return result ^ ww, ww
}
