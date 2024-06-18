package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://amazon.com",
		"https://golang.org",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	// fmt.Println(<-channel)
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-channel)
	// }

	// Repeat
	// for {
	// 	go checkLink(<-channel, channel)
	// }

	// Alterative loop
	// for link := range channel {
	// 	go checkLink(link, channel)
	// }

	// Function Literal === Anonymous function
	for link := range channel {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}()
	}

}

func checkLink(link string, channel chan string) {
	fmt.Println("Start check link:", link)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down. Err:", err)
		// channel <- "Might be down"
		channel <- link
		return
	}

	fmt.Println(link, "is up!!")
	// channel <- "Might be up"
	channel <- link
}
