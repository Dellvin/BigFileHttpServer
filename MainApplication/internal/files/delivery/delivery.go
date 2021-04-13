package delivery

import (
	"HttpBigFilesServer/MainApplication/errors"
	"HttpBigFilesServer/MainApplication/internal/files/usecase"
	"net/http"
)

type Interface interface {
	Download(w http.ResponseWriter, r *http.Request)
	Upload(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	Uc       usecase.Interface
}

func New(usecase usecase.Interface) Interface {
	return delivery{Uc: usecase}
}


func (d delivery)Download(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		_, _ = w.Write(errors.NotPost())
		return
	}

}

func (d delivery)Upload(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		_, _ = w.Write(errors.NotGet())
		return
	}
	r.Body
}