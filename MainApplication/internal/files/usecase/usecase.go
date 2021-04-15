package usecase

import (
	"HttpBigFilesServer/MainApplication/internal/files/model"
	"HttpBigFilesServer/MainApplication/internal/files/repository"
	"HttpBigFilesServer/MainApplication/pkg/logger"
	"io"
	"os"
	"time"
)

type Interface interface {
	Download(id uint64, seeker uint64) (model.File, *os.File, error)
	Upload(file io.ReadCloser, name string, size uint64) (model.File, error)
}

type usecase struct {
	info repository.InterfaceDataBase
	file repository.InterfaceFile
	log  logger.Interface
}

func New(db repository.InterfaceDataBase, sys repository.InterfaceFile, l logger.Interface) Interface {
	return usecase{
		info: db,
		file: sys,
		log:  l,
	}
}

func (u usecase) Download(id uint64, seeker uint64) (model.File, *os.File, error) {
	info, err := u.info.Get(id)
	if err != nil {
		return model.File{}, nil, err
	}

	file, err := u.file.Get(id, seeker)
	if err != nil {
		return model.File{}, nil, err
	}

	return info, file, nil
}

func (u usecase) Upload(file io.ReadCloser, name string, size uint64) (model.File, error) {
	fid, err := u.info.GenID()
	if err != nil {
		return model.File{}, ErrorCouldNotGenID
	}
	path, err := u.file.Save(fid, file, size)
	if err != nil {
		if len(path) > 0 {
			u.file.Remove(path)
		}
		return model.File{}, err
	}
	fileInfo := model.File{
		Id:       fid,
		Name:     name,
		Path:     path,
		Uploaded: uint64(time.Now().Unix()),
		Size:     size,
	}
	err = u.info.Save(fileInfo)
	if err != nil {
		return model.File{}, ErrorCouldNotSaveFileInfo
	}
	return fileInfo, nil
}
