package main

import "fmt"

func main() {
	a := []int{5, 10, 8, 1, 7, 4}
	fmt.Println(LinearSearch(a, 11))
}

// LinearSearch return true in case key is present in list or else false
func LinearSearch(list []int, key int) bool {
	for _, element := range list {
		if element == key {
			return true
		}
	}
	return false
}
