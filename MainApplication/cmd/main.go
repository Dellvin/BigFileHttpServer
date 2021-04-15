package main

import (
	"HttpBigFilesServer/MainApplication/config"
	"HttpBigFilesServer/MainApplication/pkg/logger/logrus"
)

func main() {
	log := logrus.New()
	server, err := setup(log)
	if err != nil {
		log.Panic(err)
		return
	}

	log.InfoStr("starting Main at " + config.Port)
	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
