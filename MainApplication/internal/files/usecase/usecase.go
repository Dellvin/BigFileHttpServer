package usecase

import (
	"HttpBigFilesServer/MainApplication/internal/files/repository"
	"io"
	"os"
	"sync"
)
var uploadWaiter sync.WaitGroup
type Interface interface {
	Download()
	Upload(file io.ReadCloser)
}

type usecase struct {
	Db repository.Interface
}

func New(db repository.Interface) Interface{
	return usecase{Db: db}
}

func (u usecase)Download(){

}

func (u usecase)Upload(file io.ReadCloser){
	defer file.Close()
	fd, err:=os.Create("hi")
	if err != nil{
		return
	}
	defer fd.Close()
	uploadWaiter.Add(1)
	go io.Copy(fd, file)

	uploadWaiter.Wait()
}