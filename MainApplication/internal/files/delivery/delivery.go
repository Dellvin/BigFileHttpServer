package delivery

import (
	"HttpBigFilesServer/MainApplication/errors/handler"
	"HttpBigFilesServer/MainApplication/internal/files/usecase"
	"HttpBigFilesServer/MainApplication/pkg/logger"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type Interface interface {
	Download(w http.ResponseWriter, r *http.Request)
	Upload(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	Uc       usecase.Interface
	Log		logger.Interface
}

func New(usecase usecase.Interface, l logger.Interface) Interface {
	return delivery{
		Uc: usecase,
		Log: l,
	}
}


func (d delivery)Download(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		d.Log.WarningStr("expected GET method instead of "+r.Method)
		w.WriteHeader(405)
		return
	}
	vars := mux.Vars(r)
	fIDstr := vars["Id"]
	fID, err:=strconv.Atoi(fIDstr)
	if err!=nil{
		d.Log.WarningStr("expected INT value: "+err.Error())
		w.WriteHeader(405)
		return
	}

	info, fd, err:=d.Uc.Download(uint64(fID))
	if err!=nil{
		w.WriteHeader(500)
		return
	}
	defer fd.Close()

	setHeadersForDownload(w, info)

	pR, pW := io.Pipe()
	go MultiPart(info.Name, pW, fd, d.Log)
	io.Copy(w, pR)
	_ = pR.Close()
}

func (d delivery)Upload(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		d.Log.WarningStr("expected POST method instead of "+r.Method)
		w.WriteHeader(405)
		return
	}
	fsize, err:=strconv.Atoi(r.Header.Get("File-size"))
	if err!=nil{
		d.Log.WarningStr("expected INT value: "+err.Error())
		w.WriteHeader(405)
		return
	}
	info, err:=d.Uc.Upload(r.Body, r.Header.Get("File-name"), uint64(fsize))
	if err!=nil{
		w.WriteHeader(handler.HandleDownLoadError(err))
		return
	}
	w.Write(handler.GetOkDownloadResponse(info))
}