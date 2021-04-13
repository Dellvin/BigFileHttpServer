package postgres

import (
	"HttpBigFilesServer/MainApplication/internal/files/model"
	"HttpBigFilesServer/MainApplication/internal/files/repository"
	pgwrapper "gitlab.com/slax0rr/go-pg-wrapper"
)

type dataBase struct {
	DB pgwrapper.DB
}

func New(db pgwrapper.DB) repository.Interface {
	return dataBase{DB: db}
}


func (db dataBase)GetFileInfo(id int64) (model.File, error){
	var file model.File
	exist:=db.DB.Model(&file).Where("id=?", id).Select()
	if exist !=nil{
		return model.File{}, repository.GetFileError
	}
	return file, nil
}

func (db dataBase)SetFileInfo(file model.File) error{
	_, err:=db.DB.Model(&file).Insert()
	if err != nil{
		return err
	}
	return nil
}

