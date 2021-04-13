package main

import (
"net/http"
"net/url"
"os"
"fmt"
)



func main() {
	tr := http.DefaultTransport

	client := &http.Client{
		Transport: tr,
		Timeout:   0,
	}
	r, err:=os.Open("book.txt")
	if err!=nil{
		return
	}
	defer r.Close()
	req := &http.Request{
		Method: "POST",
		URL: &url.URL{
			Scheme: "http",
			Host:   "localhost:8000",
			Path:   "/upload",
		},
		ProtoMajor: 1,
		ProtoMinor: 1,
		ContentLength: -1,
		Body: r,
	}
	fmt.Printf("Doing request\n")
	_, err = client.Do(req)
	fmt.Printf("Done request. Err: %v\n", err)
}
