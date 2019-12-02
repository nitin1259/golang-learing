package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type stdLogWriter struct{}

func (stdLogWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("these many byte written to standard output: ", len(bs))

	return len(bs), nil
}

func main() {
	fmt.Println("welcome to http packages")

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// fmt.Println("Response: ", resp)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)

	slw := stdLogWriter{}
	io.Copy(slw, resp.Body)

}
