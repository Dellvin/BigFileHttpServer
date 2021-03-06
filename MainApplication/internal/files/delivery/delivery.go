package delivery

import (
	"HttpBigFilesServer/MainApplication/config"
	"HttpBigFilesServer/MainApplication/internal/files/usecase"
	"HttpBigFilesServer/MainApplication/pkg"
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
	Uc  usecase.Interface
	Log logger.Interface
}

func New(usecase usecase.Interface, l logger.Interface) Interface {
	return delivery{
		Uc:  usecase,
		Log: l,
	}
}

func (d delivery) Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		d.Log.WarningStr("expected GET method instead of " + r.Method)
		w.WriteHeader(405)
		return
	}
	vars := mux.Vars(r)
	fIDstr := vars["Id"]
	fID, err := strconv.Atoi(fIDstr)
	if err != nil {
		d.Log.WarningStr("expected INT value: " + err.Error())
		w.WriteHeader(405)
		return
	}
	seekerStr := vars["From"]
	seeker, err := strconv.Atoi(seekerStr)
	if err != nil {
		d.Log.WarningStr("expected INT value: " + err.Error())
		w.WriteHeader(405)
		return
	}

	chunkSize:=vars["Chunk"]
	chunk:=config.ChunkSize
	if len(chunkSize)>0{
		chunk, err=strconv.Atoi(chunkSize)
		if err!=nil || chunk>config.MaxChunk{
			d.Log.WarningStr("expected INT value in Chunk-size: " + err.Error())
			w.WriteHeader(405)
			return
		}
	}

	info, fd, err := d.Uc.Download(uint64(fID), uint64(seeker))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer fd.Close()

	setHeadersForDownload(w, info)

	pR, pW := io.Pipe()
	go MultiPart(info.Name, pW, fd, d.Log, chunk)
	io.Copy(w, pR)
	_ = pR.Close()
}

func (d delivery) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		d.Log.WarningStr("expected POST method instead of " + r.Method)
		w.WriteHeader(405)
		return
	}
	fsize, err := strconv.Atoi(r.Header.Get("File-size"))
	if err != nil {
		d.Log.WarningStr("expected INT value: " + err.Error())
		w.WriteHeader(405)
		return
	}
	chunkSize:=r.Header.Get("Chunk-size")
	chunk:=config.ChunkSize
	if len(chunkSize)>0{
		chunk, err=strconv.Atoi(chunkSize)
		if err!=nil || chunk>config.MaxChunk{
			d.Log.WarningStr("expected INT value in Header Chunk-size: " + err.Error())
			w.WriteHeader(405)
			return
		}
	}
	info, err := d.Uc.Upload(r.Body, r.Header.Get("File-name"), uint64(fsize), chunk)
	if err != nil {
		w.WriteHeader(pkg.HandleDownLoadError(err))
		return
	}
	w.Write(pkg.GetOkDownloadResponse(info))
}
