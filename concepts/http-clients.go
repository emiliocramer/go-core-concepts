package concepts

import (
	"bufio"
	"fmt"
	"net/http"
)

func HttpClients() {

	// Declaring a response
	// http.Get returns *Response and an error
	// *Response has a Header, Body...
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic("couldn't establish connection")
	}
	defer resp.Body.Close()

	// Check you've established a connection
	fmt.Println("Status: ", resp.Status)

	// Create a new scanner to read the response's body (type reader)
	// Iterate through the first 5 lines of the body and s.Scan()
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
