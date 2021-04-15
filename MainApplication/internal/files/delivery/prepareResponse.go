package delivery

import (
	"HttpBigFilesServer/MainApplication/config"
	"HttpBigFilesServer/MainApplication/internal/files/model"
	"HttpBigFilesServer/MainApplication/pkg/logger"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func MultiPart(filename string, pW *io.PipeWriter, fd *os.File, l logger.Interface) {
	defer pW.Close()
	multipartW := multipart.NewWriter(pW)
	defer multipartW.Close()
	err := multipartW.SetBoundary(config.Boundary)
	if err != nil { //TODO normal function
		return
	}
	partW, err := multipartW.CreateFormFile("file", filename)
	if err != nil {
		l.Error(err)
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
			l.ErrorStr("The error reading from connector: " + err.Error())
		}
		buf = buf[:n]
	}
}

func setHeadersForDownload(w http.ResponseWriter, file model.File) {
	w.Header().Add("ProtoMajor", "1")
	w.Header().Add("ProtoMinor", "1")
	w.Header().Add("ContentLength", "-1")
	w.Header().Add("File-name", file.Name)
	w.Header().Add("File-size", strconv.Itoa(int(file.Size)))
	w.Header().Add("File-uploaded", strconv.Itoa(int(file.Uploaded)))
	w.Header().Add("File-ID", strconv.Itoa(int(file.Id)))
}
