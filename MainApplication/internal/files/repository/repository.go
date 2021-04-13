package repository

import (
	"HttpBigFilesServer/MainApplication/internal/files/model"
	"errors"
)

var GetFileError = errors.New("Could not get file")

type Interface interface {
	GetFileInfo(int64) (model.File, error)
	SetFileInfo(model.File) error
}