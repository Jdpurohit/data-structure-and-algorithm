package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(BinarySearch(a, 11))
}

// BinarySearch return true in case key is present in list or else false, it requires data in list to be sorted
func BinarySearch(list []int, key int) bool {
	if len(list) == 1 {
		return list[0] == key
	}
	middle := len(list) / 2
	if list[middle] == key {
		return true
	} else if list[middle] < key {
		return BinarySearch(list[middle+1:], key)
	} else {
		return BinarySearch(list[:middle], key)
	}
}
