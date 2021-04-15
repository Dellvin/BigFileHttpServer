package system

import (
	"HttpBigFilesServer/MainApplication/config"
	"HttpBigFilesServer/MainApplication/internal/files/repository"
	"HttpBigFilesServer/MainApplication/pkg/logger"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

type system struct {
	path string
	log  logger.Interface
}

func New(path string, l logger.Interface) repository.InterfaceFile {
	return system{
		path: path,
		log: l,
	}
}

func (s system) Get(id uint64, seeker uint64) (*os.File, error){
	path:=s.path+strconv.Itoa(int(id))
	file, err:=os.Open(path)
	if err!=nil{
		s.log.WarningStr("Could not open file: "+path)
		return nil, ErrorOpening
	}
	_, err=file.Seek(int64(seeker), 0)
	if err!=nil{
		s.log.WarningStr("Could not seek file: "+path)
		return nil, ErrorSeeking
	}
	return file, nil
}

func (s system) Save(id uint64, file io.ReadCloser, size uint64) (string, error){
	path:=s.path+strconv.Itoa(int(id))
	fd, err := os.Create(path)

	if err != nil {
		s.log.WarningStr("Could not create file: "+path)
		return "", ErrorCreateFile
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
				s.log.Error(err)
				return path, ErrorLoading
			}
			total+=uint64(n)
			_, err=fd.Write(buf[:n])
			if err!=nil{
				s.log.Error(err)
				return path,ErrorWriteFile
			}
		}
	}
	if total != size {
		s.log.ErrorStr("length does not match")
		return path,ErrorSizesDoesNotMatch
	}
	return path, nil
}

func (s system)Remove(path string){
	err:=os.Remove(path)
	if err!=nil{
		s.log.Error(err)
	}
}