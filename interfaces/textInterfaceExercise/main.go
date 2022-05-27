package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(string(data))

}

//resp, err := http.Get("http://google.com")
//if err != nil {
//	fmt.Printf("An Error has ocurred.\nError: ", err)
//	os.Exit(1)
//}
//
////bs := make([]byte, 99999)
////resp.Body.Read(bs)
//
//lw := logWriter{}
//io.Copy(lw, resp.Body)

//func (logWriter) Write(bs []byte) (int, error) {
//	fmt.Println(string(bs))
//
//}
