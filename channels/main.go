package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://mercadolivre.com.br",
		"http://facebook.com",
		"http://twitter.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%v is not responding properly at time: %v \n", link, time.Now())
		c <- link
		return
	}
	fmt.Printf("%v is working fine. OK! \n", link)
	c <- link

}
