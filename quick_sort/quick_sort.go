//package quick_sort
package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int, left int, right int)  {
	l := left
	r := right
	middle := arr[(l+r)/2]
	tmp := 0

	for ;l<r; {
		for ;arr[l] < middle; {
			l ++
		}
		for ;arr[r] > middle; {
			r --
		}
		if l >= r {
			break
		}
		tmp = arr[l]
		arr[l] = arr[r]
		arr[r] = tmp

		if arr[l] == middle {
			r --
		}
		if arr[r] == middle {
			l ++
		}
	}

	if l == r {
		l ++
		r --
	}

	if left < r {
		QuickSort(arr, left, r)
	}
	if right > l {
		QuickSort(arr, l, right)
	}
}

func main() {
	arr := make([]int, 20)
	for i:=0; i<20; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	QuickSort(arr, 0, 19)
	fmt.Println(arr)
}

