package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

const Boundary = "DellvinBlackDellvinBlackDellvinBlackDellvinBlack"

func main() {
	tr := http.DefaultTransport
	client := http.Client{
		Transport: tr,
	}
	resp, err :=client.Get("http://localhost:8080/download/4")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	if resp.StatusCode==200{
		buf := make([]byte, 4096)
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
				} else if err!=nil{
					fmt.Println(err.Error())
					return
				}
				fmt.Printf(string(buf[:n]))
			}
			fmt.Printf(string(buf[:n]))
		}
		fmt.Println("Done")
	} else{
		var body = make([]byte, 256)
		resp.Body.Read(body)
		fmt.Println(string(body))
	}

}