package main

import "fmt"

func linearSearch(arr [5]int, elem int) int {
	for i,e := range arr{
		if e == elem {
			return i+1
		}
	}
	return 0
}

func main(){
	arr := [5]int{1,2,3,4,5}
	fmt.Println(linearSearch(arr, 10))
}