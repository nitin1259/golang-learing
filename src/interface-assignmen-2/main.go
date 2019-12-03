package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("welcome to assignment no 2")
	ss := os.Args

	// fmt.Println(ss)
	if len(ss) > 1 {
		fileName := ss[1]

		fmt.Println("fileName: ", fileName)

		f, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		io.Copy(os.Stdout, f)
	} else {
		fmt.Println("No argumnent passed")
	}

}
