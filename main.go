package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now();
	userName := fetchUser(); // 100ms

	responseChannel := make(chan any, 2);
	
	waitGroup := &sync.WaitGroup{}
	
	waitGroup.Add(2);
	go fetchUserLikes(userName, responseChannel, waitGroup); //150ms
	go fetchUserMatches(userName, responseChannel, waitGroup); //100ms
	
	waitGroup.Wait(); //blocks until we have 2 responses

	close(responseChannel);

	for response := range responseChannel {
		fmt.Println("response: ", response);
	}
	


	fmt.Println("took us: ", time.Since(start));
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100);

	return "Hey, you made it here! Hire me now";
}

func fetchUserLikes(userName string, responseChannel chan<- any, wg *sync.WaitGroup )  {
	defer wg.Done();
	
	time.Sleep(time.Millisecond * 150);
	
	responseChannel <- 42;
}

func fetchUserMatches(userName string, responseChannel chan<- any, wg *sync.WaitGroup ) {
	defer wg.Done();
	
	time.Sleep(time.Millisecond * 100);
	responseChannel <- "Daria"
}