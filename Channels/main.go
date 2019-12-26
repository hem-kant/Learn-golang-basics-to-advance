package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://golang.org",
		"http://www.stackoverflow.com",
	}

	// creating a new channel
	c := make(chan string)
	for _, link := range links {
		//checkLink(link)
		// using GO to make concurrent/routine calls with out waiitng
		go checkLink(link, c)
	}
	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}
	// ********* OR ************
	//for {
	//	go checkLink(<-c, c)
	//}
	// ********* OR ************
	for l := range c {
		// go function literal = anonymous or lambada funtions
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // passing argument to linteral
		// 5 sec sleep
		//go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link // check agian
		return
	}
	fmt.Println(link, "Is up!!")
	c <- link // check agian
}
