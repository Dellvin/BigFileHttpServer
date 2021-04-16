package tests

import (
	"HttpBigFilesServer/MainApplication/config"
	"HttpBigFilesServer/MainApplication/internal/files/model"
	"HttpBigFilesServer/MainApplication/internal/files/repository/system"
	"HttpBigFilesServer/MainApplication/internal/files/usecase"
	"HttpBigFilesServer/MainApplication/pkg/logger/logrus"
	"HttpBigFilesServer/MainApplication/test/mock_repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDownloadInfoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var l = logrus.New()
	mockFile := mock_repository.NewMockInterfaceFile(ctrl)
	mockInfo := mock_repository.NewMockInterfaceDataBase(ctrl)
	mockInfo.EXPECT().Get(uint64(18)).Return(model.File{}, system.ErrorCreateFile)
	sys := usecase.New(mockInfo, mockFile, l)
	fInfo, file, err := sys.Download(18, 0)
	assert.NotNil(t, err)
	assert.Nil(t, file)
	assert.Equal(t, fInfo, model.File{})
}

func TestDownloadFileError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var l = logrus.New()
	mockFile := mock_repository.NewMockInterfaceFile(ctrl)
	mockInfo := mock_repository.NewMockInterfaceDataBase(ctrl)
	mockInfo.EXPECT().Get(uint64(18)).Return(model.File{}, nil)
	mockFile.EXPECT().Get(uint64(18), uint64(0)).Return(nil, system.ErrorCreateFile)
	sys := usecase.New(mockInfo, mockFile, l)
	fInfo, file, err := sys.Download(18, 0)
	assert.NotNil(t, err)
	assert.Nil(t, file)
	assert.Equal(t, fInfo, model.File{})
}

func TestDownloadSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var l = logrus.New()
	mockFile := mock_repository.NewMockInterfaceFile(ctrl)
	mockInfo := mock_repository.NewMockInterfaceDataBase(ctrl)
	mockInfo.EXPECT().Get(uint64(18)).Return(model.File{Id: 1}, nil)
	mockFile.EXPECT().Get(uint64(18), uint64(0)).Return(&os.File{}, nil)
	sys := usecase.New(mockInfo, mockFile, l)
	fInfo, file, err := sys.Download(18, 0)
	assert.Nil(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, fInfo, model.File{Id: 1})
}

func TestUploadGenIDError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var l = logrus.New()
	mockFile := mock_repository.NewMockInterfaceFile(ctrl)
	mockInfo := mock_repository.NewMockInterfaceDataBase(ctrl)
	mockInfo.EXPECT().GenID().Return(uint64(0), usecase.ErrorCouldNotGenID)
	sys := usecase.New(mockInfo, mockFile, l)
	test := strings.NewReader("shiny!")
	testReader := ioutil.NopCloser(test)
	fInfo, err := sys.Upload(testReader, "file.txt", uint64(12), config.ChunkSize)
	assert.NotNil(t, err)
	assert.Equal(t, fInfo, model.File{})
}

func TestUploadSaveError(t *testing.T) {
	test := strings.NewReader("")
	testReader := ioutil.NopCloser(test)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var l = logrus.New()
	mockFile := mock_repository.NewMockInterfaceFile(ctrl)
	mockInfo := mock_repository.NewMockInterfaceDataBase(ctrl)
	mockInfo.EXPECT().GenID().Return(uint64(0), nil)
	mockFile.EXPECT().Save(uint64(0), testReader, uint64(12), config.ChunkSize).Return("", usecase.ErrorCreateFile)
	sys := usecase.New(mockInfo, mockFile, l)
	fInfo, err := sys.Upload(testReader, "file.txt", uint64(12), config.ChunkSize)
	assert.NotNil(t, err)
	assert.Equal(t, fInfo, model.File{})
}

func TestUploadSaveWithPathError(t *testing.T) {
	test := strings.NewReader("")
	testReader := ioutil.NopCloser(test)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var l = logrus.New()
	mockFile := mock_repository.NewMockInterfaceFile(ctrl)
	mockInfo := mock_repository.NewMockInterfaceDataBase(ctrl)
	mockInfo.EXPECT().GenID().Return(uint64(0), nil)
	mockFile.EXPECT().Save(uint64(0), testReader, uint64(12), config.ChunkSize).Return("path", usecase.ErrorCreateFile)
	mockFile.EXPECT().Remove("path").Return()
	sys := usecase.New(mockInfo, mockFile, l)
	fInfo, err := sys.Upload(testReader, "file.txt", uint64(12), config.ChunkSize)
	assert.NotNil(t, err)
	assert.Equal(t, fInfo, model.File{})
}

func TestUploadSaveInfoError(t *testing.T) {
	test := strings.NewReader("")
	testReader := ioutil.NopCloser(test)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var l = logrus.New()
	mockFile := mock_repository.NewMockInterfaceFile(ctrl)
	mockInfo := mock_repository.NewMockInterfaceDataBase(ctrl)
	mockInfo.EXPECT().GenID().Return(uint64(0), nil)
	mockFile.EXPECT().Save(uint64(0), testReader, uint64(12), config.ChunkSize).Return("path", nil)
	mockInfo.EXPECT().Save(model.File{
		Id:       0,
		Name:     "file.txt",
		Path:     "path",
		Uploaded: uint64(time.Now().Unix()),
		Size:     12,
	}).Return(usecase.ErrorCouldNotSaveFileInfo)
	sys := usecase.New(mockInfo, mockFile, l)
	fInfo, err := sys.Upload(testReader, "file.txt", uint64(12), config.ChunkSize)
	assert.NotNil(t, err)
	assert.Equal(t, fInfo, model.File{})
}

func TestUploadSuccess(t *testing.T) {
	test := strings.NewReader("")
	testReader := ioutil.NopCloser(test)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var l = logrus.New()
	mockFile := mock_repository.NewMockInterfaceFile(ctrl)
	mockInfo := mock_repository.NewMockInterfaceDataBase(ctrl)
	mockInfo.EXPECT().GenID().Return(uint64(0), nil)
	mockFile.EXPECT().Save(uint64(0), testReader, uint64(12), config.ChunkSize).Return("path", nil)
	mockInfo.EXPECT().Save(model.File{
		Id:       0,
		Name:     "file.txt",
		Path:     "path",
		Uploaded: uint64(time.Now().Unix()),
		Size:     12,
	}).Return(nil)
	sys := usecase.New(mockInfo, mockFile, l)
	fInfo, err := sys.Upload(testReader, "file.txt", uint64(12), config.ChunkSize)
	assert.Nil(t, err)
	assert.Equal(t, fInfo, model.File{
		Id:       0,
		Name:     "file.txt",
		Path:     "path",
		Uploaded: uint64(time.Now().Unix()),
		Size:     12,
	})
}
