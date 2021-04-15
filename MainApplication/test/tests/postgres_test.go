package tests

import (
	mock "HttpBigFilesServer/MainApplication/test/mock_postgres"
	"github.com/stretchr/testify/assert"
	"testing"

)

func TestGet(t *testing.T) {
	db, r := mock.MockFileDB()
	query := mock.MockFile(db)

	query.On("Where", "id=?", mock.FileInfoTest.Id).Return(query)
	query.On("Select").Return(nil)

	answerCorrect, err := r.Get(mock.FileInfoTest.Id)
	assert.Nil(t, err)
	assert.Equal(t, mock.FileInfoTest, answerCorrect)
}

func TestSave(t *testing.T) {
	db, r := mock.MockFileDB()
	query := mock.MockFile(db)
	mockRes:= mock.MockResult{}
	query.On("Insert").Return(mockRes, nil)

	err := r.Save(mock.FileInfoTest)
	assert.Nil(t, err)
}