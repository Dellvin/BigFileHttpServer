package main

import (
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const boundary = "DellvinBlackDellvinBlackDellvinBlackDellvinBlack"

type Flags struct {
	src  *string
	addr *string
}

func main() {
	f := setupCLArgs()
	if f.addr == nil {
		return
	}
	fname := getFileName(*f.src)
	tr := http.DefaultTransport

	client := &http.Client{
		Transport: tr,
		//Timeout:   10 * time.Second,
	}

	fmt.Println("Set up pipe")
	pR, pW := io.Pipe()
	fd, err := os.Open(*f.src)
	if err != nil {
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
		partW, err0 := multipartW.CreateFormFile("file", *f.src)
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
			Host:   *f.addr,
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
		fname,
	)
	req.Header.Set("File-size", strconv.Itoa(int(fi.Size())))
	fmt.Printf("Doing request\n")
	time.Sleep(time.Second)
	_, err = client.Do(req)
	pR.Close()
	fmt.Printf("Done request. Err: %v\n", err)
}

func setupCLArgs() Flags {
	f := Flags{}
	f.src = flag.String("src", "./static/file.txt", "the source for uploaded file")
	f.addr = flag.String("addr", "localhost:8080", "address of remote server")

	flag.Parse()
	return f
}

func prepareURL(f Flags) string {
	url := *f.addr
	url += "/upload"
	return url
}

func getFileName(path string) string {
	splited := strings.Split(path, "/")
	return splited[len(splited)-1]
}
