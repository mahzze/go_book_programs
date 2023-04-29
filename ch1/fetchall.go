package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	file, err := os.OpenFile("/tmp/fetchallresults", os.O_WRONLY, 600)
	if err != nil {
		file, err = os.Create("/tmp/fetchallresults")
		if err != nil {
			os.Exit(1)
		}
	}
	for range os.Args[1:] {
		file.WriteString(<-ch + "\n")
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, response.Body)
	response.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.fs %7d %s", secs, nbytes, url)
}
