package delivery

import (
	"HttpBigFilesServer/MainApplication/config"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func MultiPart(filename string, pW *io.PipeWriter, fd *os.File) {
	defer pW.Close()
	multipartW := multipart.NewWriter(pW)
	defer multipartW.Close()
	err:=multipartW.SetBoundary(config.Boundary)
	if err!=nil{//TODO normal function
		fmt.Println(err.Error())
		return
	}
	partW, err := multipartW.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	connector := io.TeeReader(fd, partW)
	buf := make([]byte, 4096)
	for {
		/* fd -> connector -> partW -> multipartW -> pW -> pR */
		n, err := connector.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("The error reading from connector: %v", err)
		}
		buf = buf[:n]
	}
}

