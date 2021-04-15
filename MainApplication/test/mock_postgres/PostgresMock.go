package mock_postgres

import (
	FileModel "HttpBigFilesServer/MainApplication/internal/files/model"
	"HttpBigFilesServer/MainApplication/internal/files/repository"
	"HttpBigFilesServer/MainApplication/internal/files/repository/postgres"
	"HttpBigFilesServer/MainApplication/pkg/logger/logrus"
	model "github.com/go-pg/pg/v9/orm"
	"github.com/stretchr/testify/mock"
	"gitlab.com/slax0rr/go-pg-wrapper/mocks"
	ormmocks "gitlab.com/slax0rr/go-pg-wrapper/mocks/orm"
)

var FileInfoTest =FileModel.File{
	Id: 0,
	Name: "text.txt",
	Path: "static/text.txt",
	Uploaded: 1,
	Size: 12,
}

func MockFile(db *mocks.DB) *ormmocks.Query  {
	query := new(ormmocks.Query)
	mockCall := db.On("Model", mock.AnythingOfType("*model.File")).Return(query)
	mockCall.RunFn = func(args mock.Arguments) {
		User := args[0].(*FileModel.File)
		*User = FileInfoTest
	}
	return query
}

func MockFileDB() (*mocks.DB, repository.InterfaceDataBase) {
	db := new(mocks.DB)
	l := logrus.New()
	r := postgres.New(db, l)
	return db, r
}

type MockResult struct {
}

func (r MockResult) Model() model.Model {
	panic("implement me!")
}
func (r MockResult) RowsAffected() int {
	panic("implement me!")
}
func (r MockResult) RowsReturned() int {
	panic("implement me!")
}
