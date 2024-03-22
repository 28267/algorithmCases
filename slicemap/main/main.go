package main

import (
	"fmt"
)

func main() {

	slice1 := []int{0, 1, 2, 3, 4}
	map1 := make(map[int]*int)
	var slice3 []int
	for k1, v1 := range slice1 {
		map1[k1] = &v1
	}
	for k2, v2 := range map1 {
		fmt.Println(k2, "->", *v2)
	}
	fmt.Println()

	slice2 := []int{1, 2, 3, 4, 5}
	map2 := make(map[int]*int)
	for k3, v3 := range slice2 {
		value := v3
		map2[k3] = &value
	}
	for k4 := range map2 {
		slice := k4
		slice3 = append(slice3, slice)
	}

	for j := 0; j < 4; j++ {
		for i := 0; i < len(slice3)-1-j; i++ {
			if temp := 0; slice3[i] > slice3[i+1] {
				temp = slice3[i]
				slice3[i] = slice3[i+1]
				slice3[i+1] = temp
			}
		}
	}

	for _, v5 := range slice3 {
		fmt.Println(v5, "---->", *map2[v5])
	}
}

// for i := 0; i < len(arr1) - 1; i ++ {
// 	if temp := 0; arr1[i] > arr1[i + 1] {
// 		temp = arr1[i]
// 		arr1[i] = arr1[i+1]
// 		arr1[i+1] = temp
// 	}
// }
// for i := 0; i < len(arr1) - 2; i ++ {
// 	if temp := 0; arr1[i] > arr1[i + 1] {
// 		temp = arr1[i]
// 		arr1[i] = arr1[i+1]
// 		arr1[i+1] = temp
// 	}
// }
// for i := 0; i < len(arr1) - 3; i ++ {
// 	if temp := 0; arr1[i] > arr1[i + 1] {
// 		temp = arr1[i]
// 		arr1[i] = arr1[i+1]
// 		arr1[i+1] = temp
// 	}
// }
// fmt.Println(k4, "--->", v4)
