package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(subarray_sum(arr, 3)) // teste para retornar true
	fmt.Println(subarray_sum(arr, 4)) // teste para retornar false
}

func subarray_sum(arr []int, target int) bool {
	var bool bool
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j]+arr[j+1] == target {
				bool = true
			}
		}
	}
	return bool
}
