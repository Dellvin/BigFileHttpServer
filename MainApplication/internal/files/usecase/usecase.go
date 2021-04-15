package usecase

import (
	"HttpBigFilesServer/MainApplication/config"
	"HttpBigFilesServer/MainApplication/internal/files/model"
	"HttpBigFilesServer/MainApplication/internal/files/repository"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

type Interface interface {
	Download()
	Upload(file io.ReadCloser, name string, size uint64) error
}

type usecase struct {
	rep repository.Interface
}

func New(db repository.Interface) Interface {
	return usecase{rep: db}
}

func (u usecase) Download() {

}

func (u usecase) Upload(file io.ReadCloser, name string, size uint64) error {
	fd, err := os.Create(name)
	if err != nil {
		fmt.Println("Creating")
		return ErrorCreateFile
	}
	defer fd.Close()
	buf := make([]byte, 4096)
	var total uint64 = 0
	partReader := multipart.NewReader(file, config.Boundary)
	for {
		part, err := partReader.NextPart()
		if err != nil {
			break
		}
		var n int
		for {
			n, err = part.Read(buf)
			if err == io.EOF || err == io.ErrUnexpectedEOF{
				break
			} else if err !=nil {
				return ErrorLoading
			}
			total+=uint64(n)
			_, err=fd.Write(buf[:n])
			if err!=nil{
				fmt.Println("Writing")
				return ErrorWriteFile
			}
		}
	}
	if total != size {
		return ErrorSizesDoesNotMatch
	}
	fid, err:=u.rep.IsIdExist()
	if err!=nil{
		return ErrorCouldNotGenID
	}
	err=u.rep.SetFileInfo(model.File{
		Id: fid,
		Name: name,
		Path: config.Path +name,
		Uploaded: uint64(time.Now().Unix()),
		Size: size,
		})
	if err!=nil{
		return ErrorCouldNotSaveFileInfo
	}
	return nil
}

