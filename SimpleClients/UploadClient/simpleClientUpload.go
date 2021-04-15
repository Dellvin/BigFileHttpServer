package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const filename="book.txt"

const boundary = "DellvinBlackDellvinBlackDellvinBlackDellvinBlack"

func main() {
	tr := http.DefaultTransport

	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	fmt.Println("Set up pipe")
	pR, pW := io.Pipe()
	fd, err := os.Open(filename)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	defer fd.Close()
	go func() {
		defer pW.Close()
		// Set up multipart body for reading
		multipartW := multipart.NewWriter(pW)
		//defer multipartW.Close()
		fmt.Println("Set up multipart writer")
		multipartW.SetBoundary(boundary)
		fmt.Println("Set up boundary")
		partW, err0 := multipartW.CreateFormFile("file", filename)
		fmt.Println("Set up part writer")
		if err0 != nil {
			panic("Something is amiss creating a part")
		}

		if err != nil {
			fmt.Println("ERROR WHILE OPENING FILE")
			return
		}
		connector := io.TeeReader(fd, partW)
		buf := make([]byte, 4096)
		for {
			/* stdin -> connector -> partW -> multipartW -> pW -> pR */
			_, err := connector.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("The error reading from connector: %v", err)
			}
		}

	}()

	// Send http request chunk encoding the multipart message
	var req = &http.Request{
		Method: "POST",
		URL: &url.URL{
			Scheme: "http",
			Host:   "localhost:8080",
			Path:   "/upload",
		},
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: -1,
		Body:          pR,
	}
	fi, err := fd.Stat()
	req.Header = make(http.Header)
	req.Header.Set(
		"File-name",
		"book.txt",
	)
	req.Header.Set("File-size", strconv.Itoa( int(fi.Size())))
	fmt.Printf("Doing request\n")
	_, err = client.Do(req)
	pR.Close()
	fmt.Printf("Done request. Err: %v\n", err)
}
