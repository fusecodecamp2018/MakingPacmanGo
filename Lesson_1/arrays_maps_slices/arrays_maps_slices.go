package main

import "fmt"

func main()  {
	//an array
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	//a map
	var amap = make(map[int]string)
	amap[1] = "thing"
	amap[2] = "stuff"
	fmt.Println(amap)

	//a slice
	slice := make([]string, 3)
	slice[0] = "first"
	slice[1] = "second"
	slice[2] = "third"
	fmt.Println(slice)
}
