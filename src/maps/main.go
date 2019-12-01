package main

import "fmt"

func main() {
	fmt.Println("welcome to go Maps")

	// simple and plain way of declaring map
	color := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf723",
	}

	fmt.Println(color)

	// 2nd way of declaring the map
	var color2 map[string]string
	// color2["white"] = "#ffffff" // this is valid but color2.white is not

	fmt.Println(color2)

	//3rd way is to use inbuild function make

	color3 := make(map[int]string)

	color3[10] = "#999999"
	color3[2] = "#333333"

	fmt.Println(color3)

	// deletion of key value in map

	delete(color3, 10)

	fmt.Println(color3)
}
