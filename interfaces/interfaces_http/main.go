package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("An Error has ocurred.\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
