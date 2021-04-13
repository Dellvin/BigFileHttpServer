package errors

import "HttpBigFilesServer/MainApplication/internal/files/model"

type Response struct {
	Code        uint16
	Description string
	sid         string
	User        model.File
}