package main

import (
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

type Flags struct {
	dest *string
	addr *string
	id   *int
	seek *int
	chunk *int
}

const Boundary = "DellvinBlackDellvinBlackDellvinBlackDellvinBlack"

func main() {
	f := setupCLArgs()
	if f.addr == nil || f.id == nil {
		return
	}
	addr := prepareURL(f)
	tr := http.DefaultTransport
	client := http.Client{
		Transport: tr,

	}

	resp, err := client.Get(addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	if resp.StatusCode == 200 {
		var fd *os.File
		if f.dest != nil {
			fd, err = os.Create(*f.dest)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		buf := make([]byte, *f.chunk)
		partReader := multipart.NewReader(resp.Body, Boundary)
		for {
			part, err := partReader.NextPart()
			if err == io.EOF {
				break
			}
			var n int
			for {
				n, err = part.Read(buf)
				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Println(err.Error())
					return
				}
				if fd != nil {
					fd.Write(buf[:n])
				} else {
					fmt.Printf(string(buf[:n]))
				}

			}
			if fd != nil {
				fd.Write(buf[:n])
			} else {
				fmt.Printf(string(buf[:n]))
			}
		}
		fmt.Println("Done")
	} else {
		var body = make([]byte, 256)
		resp.Body.Read(body)
		fmt.Println(string(body))
	}

}

func setupCLArgs() Flags {
	f := Flags{}
	f.dest = flag.String("dest", "./static/file.txt", "a destination for downloaded file")
	f.addr = flag.String("addr", "http://localhost:8080", "address of remote server")
	f.id = flag.Int("id", 18, "file id")
	f.seek = flag.Int("seek", 32, "pos to seek in file")
	f.chunk=flag.Int("chunk", 4096, "size of chunk")
	flag.Parse()
	return f
}

func prepareURL(f Flags) string {
	url := *f.addr
	url += "/download/"
	url += strconv.Itoa(*f.id)
	if f.chunk!=nil{
		url += "/" + strconv.Itoa(*f.chunk)
	}
	if f.seek != nil {
		url += "/" + strconv.Itoa(*f.seek)
	}
	return url
}
