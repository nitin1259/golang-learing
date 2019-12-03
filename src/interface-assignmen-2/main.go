package main

import (
	"fmt"
	"io"
	"os"
)

type mylogger struct{}

func (mylogger) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	fmt.Println("Byte written to standard output: ", len(b))
	return len(b), nil
}

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
		log := mylogger{}
		io.Copy(log, f)
	} else {
		fmt.Println("No argumnent passed")
	}

}
