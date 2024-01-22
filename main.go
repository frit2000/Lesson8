package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Arrays struct {
	Arr []int
}

func (A *Arrays) Init(count int) {
	A.Arr = make([]int, count)
	for i := 0; i < count; i++ {
		A.Arr[i] = rand.Intn(10)
	}
}

func (A *Arrays) QSort(L, R int) {
	if L >= R {
		return
	}
	i := L
	j := R
	p := A.Arr[R]
	for i < j {
		for (A.Arr[i] <= p) && (j > i) {
			i++
		}
		for (A.Arr[j] > p) && (j > i) {
			j--
		}
		A.Arr[i], A.Arr[j] = A.Arr[j], A.Arr[i]
		i++
		j--
	}
	A.QSort(L, j)
	A.QSort(j+1, R)
}

func (A *Arrays) MergeSort(L, R int) {
	if L >= R {
		return
	}
	p := (L + R) / 2
	A.MergeSort(L, p)
	A.MergeSort(p+1, R)

	temp := make([]int, R-L+1)
	m := 0
	a := L
	b := p + 1
	// for a := L; a <= p; a++ {
	// 	for b := p + 1; b <= R; b++ {
	// 		if A.Arr[a] > A.Arr[b] {
	// 			temp[m] = A.Arr[b]
	// 			m++
	// 		} else {
	// 			temp[m] = A.Arr[a]
	// 			m++
	// 		}
	// 	}
	// }

	for a <= p && b <= R {
		if A.Arr[a] > A.Arr[b] {
			temp[m] = A.Arr[b]
			m++
			b++
		} else {
			temp[m] = A.Arr[a]
			m++
			a++
		}
	}

	for a <= p {
		temp[m] = A.Arr[a]
		m++
		a++
	}

	for b <= R {
		temp[m] = A.Arr[b]
		m++
		b++
	}

	for i := L; i <= R; i++ {
		A.Arr[i] = temp[i-L]
	}
}

func main() {
	var Arr Arrays
	for i := 10; i <= 100_000_000_000; i *= 10 {
		Arr.Init(i)
		//	fmt.Println("init=", Arr)
		start := time.Now()
		Arr.QSort(0, i-1)
		//Arr.MergeSort(0, i-1)
		duration := time.Since(start)
		//	fmt.Println("sorted=", Arr)
		fmt.Println("count=", i, "time = ", duration)

	}
}
