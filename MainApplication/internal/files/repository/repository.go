package repository

import (
	"HttpBigFilesServer/MainApplication/internal/files/model"
	"errors"
	"io"
	"os"
)

//go:generate mockgen -source=./DataBaseRequests.go -destination=../../../test/mock_LetterRepository/LetterRepositoryMock.go
var GetFileError = errors.New("Could not get file")

type InterfaceDataBase interface {
	Get(uint64) (model.File, error)
	Save(model.File) error
	GenID() (uint64, error)
}

type InterfaceFile interface {
	Get(uint64, uint64) (*os.File, error)
	Save(uint64, io.ReadCloser, uint64) (string, error)
	Remove(string)
}
