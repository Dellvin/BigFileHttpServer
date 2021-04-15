package postgres

import (
	"HttpBigFilesServer/MainApplication/internal/files/model"
	"HttpBigFilesServer/MainApplication/internal/files/repository"
	"HttpBigFilesServer/MainApplication/pkg"
	"HttpBigFilesServer/MainApplication/pkg/logger"
	pgwrapper "gitlab.com/slax0rr/go-pg-wrapper"
)

type dataBase struct {
	DB  pgwrapper.DB
	log logger.Interface
}

func New(db pgwrapper.DB, l logger.Interface) repository.InterfaceDataBase {
	return dataBase{
		DB:  db,
		log: l,
	}
}

func (db dataBase) Get(id uint64) (model.File, error) {
	var file model.File
	exist := db.DB.Model(&file).Where("id=?", id).Select()
	if exist != nil {
		return model.File{}, repository.GetFileError
	}
	return file, nil
}

func (db dataBase) Save(file model.File) error {
	_, err := db.DB.Model(&file).Insert()
	if err != nil {
		db.log.Error(err)
		return err
	}
	return nil
}

func (db dataBase) GenID() (uint64, error) {
	for {
		id, err := pkg.GenId()
		if err != nil {
			db.log.Error(err)
			return 0, err
		}
		file := model.File{Id: id}
		exist := db.DB.Model(file).Where("id=?", id).Select()
		if exist != nil { //TODO make a normal error handler
			return id, nil
		}
	}
}
