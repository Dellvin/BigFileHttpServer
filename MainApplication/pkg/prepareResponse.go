package pkg

import (
	"HttpBigFilesServer/MainApplication/internal/files/model"
	"HttpBigFilesServer/MainApplication/internal/files/usecase"
	"encoding/json"
)

func HandleDownLoadError(err error) int {
	if err == usecase.ErrorSizesDoesNotMatch ||
		err == usecase.ErrorCreateFile ||
		err == usecase.ErrorWriteFile ||
		err == usecase.ErrorLoading ||
		err == usecase.ErrorCouldNotGenID ||
		err == usecase.ErrorCouldNotSaveFileInfo {
		return 500
	}
	return 400
}

func GetOkDownloadResponse(file model.File) []byte {
	ans, _ := json.Marshal(file)
	return ans
}
