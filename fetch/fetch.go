/**
Fecth given URL
Example from the Donovan and Kernighan book
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix("http://", url) {
			url = "http://" + url
		}

		response, err := http.Get(url)

		fmt.Printf("Got status: %s", response.Status)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		if _, err := io.Copy(os.Stdout, response.Body); err != nil {
			fmt.Fprintf(os.Stderr, "Reading %s: %v\n", url, err)
			os.Exit(1)
		}

		response.Body.Close()
	}
}
