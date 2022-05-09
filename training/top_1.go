package main

import "fmt"

// get maximum value of list
func max(lst []int) int{
	m := lst[0]
	for i:=1;i<=len(lst) - 1 ;i++{
		if lst[i] > m{
			m = lst[i]
		}
	}
	return m
}

// get keys of a map
func getKeys(m map[int]int) []int{
	keys := make([]int, 0)
	for k := range m{
		keys = append(keys, k)
	}
	return keys
}

// get values of map
func getValues(m map[int]int) []int{
	values := make([]int, 0)
	for value := range m{
		values = append(values, m[value])
	}
	return values
}

// top one algorythm
func top_one(lst []int) []int{

	var mostFrequent int
	result := make([]int, 0)
	values := make(map[int]int)

	for _,elem := range lst{
		if _, ok := values[elem]; ok{  // check if elem in values 
			values[elem] += 1
		}else{
			values[elem] = 1
		}
	}

	mostFrequent = max(getValues(values))

	for _,key := range getKeys(values){
		if values[key] == mostFrequent{
			result = append(result, key)
		}else{
			continue
		}
	}
	return result

}

func main(){
	slice := make([]int, 0)
	slice = append(slice, 1,1,1,2,2,2,2, 1)

	
	fmt.Println(top_one(slice))
	
}