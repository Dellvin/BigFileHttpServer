package delivery

import (
	"HttpBigFilesServer/MainApplication/errors"
	"HttpBigFilesServer/MainApplication/internal/files/usecase"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)
var Wg =sync.WaitGroup{}

type Interface interface {
	Download(w http.ResponseWriter, r *http.Request)
	Upload(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	Uc       usecase.Interface
}

func New(usecase usecase.Interface) Interface {
	return delivery{Uc: usecase}
}


func (d delivery)Download(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	filename := vars["Id"]
	if r.Method != http.MethodGet{
		_, _ = w.Write(errors.NotGet())
		return
	}
	pR, pW := io.Pipe()
	fd, err := os.Open(filename)
	if err!=nil{
		return//TODO
	}
	defer fd.Close()
	w.Header().Add("ProtoMajor", "1")
	w.Header().Add("ProtoMinor", "1")
	w.Header().Add("ContentLength", "-1")
	w.Header().Add("File-name", filename)
	w.Header().Add("File-size", filename)
	w.Header().Add("File-uploaded", filename)
	go MultiPart(filename, pW, fd)
	_, err = io.Copy(w, pR)
	if err!=nil{
		//TODO
	}
	_ = pR.Close()
}

func (d delivery)Upload(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		_, _ = w.Write(errors.NotPost())
		return
	}
	fsize, err:=strconv.Atoi(r.Header.Get("File-size"))
	if err!=nil{
		//TODO
	}
	err=d.Uc.Upload(r.Body, "kek.txt", uint64(fsize))
	if err!=nil{

	}


	fmt.Printf("TRANSMISSION COMPLETE")
}