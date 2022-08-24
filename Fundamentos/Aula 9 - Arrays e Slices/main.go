package main

import "fmt"

func main() {

	//---------------------------- ARRAYS ----------------------------
	var array1 [5]int
	array2 := [1]string{"a"}
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(array1, array2, array3)

	//---------------------------- SLICES ----------------------------
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 18)

	fmt.Println(slice)

}
