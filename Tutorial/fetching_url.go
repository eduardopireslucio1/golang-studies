// Fetch prints the content found at a URL.
package main 

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)

// Exercise 1.7: The function call io.Copy(dst, src) reads from src and writes to dst. 
// Use it instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring 
// a buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.

//Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL if it 
//is missing. You might want to use strings.HasPrefix.

//Exercise 1.9: Modify fetch to also print the HTTP status code, found in resp.Status.
func fetchUrl() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("HTTP Status Code:", resp.Status)
		_, err = io.Copy(os.Stderr, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func main() {
	fetchUrl()
}