package main

import (
	"HttpBigFilesServer/MainApplication/config"
	"HttpBigFilesServer/MainApplication/internal/files/delivery"
	"HttpBigFilesServer/MainApplication/internal/files/repository/postgres"
	"HttpBigFilesServer/MainApplication/internal/files/repository/system"
	"HttpBigFilesServer/MainApplication/internal/files/usecase"
	"HttpBigFilesServer/MainApplication/internal/postgresSetup"
	"HttpBigFilesServer/MainApplication/pkg/logger"
	"flag"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"strconv"
)

type Flags struct {
	port *int
}

func setup(l logger.Interface) (http.Server, error) {
	cliArgSetup()
	db, err := initDataBase()
	if err != nil {
		return http.Server{}, err
	}

	router := initRouters(db, l)

	handler := cors.New(cors.Options{
		AllowedOrigins:   config.AllowedOriginsCORS,
		AllowedHeaders:   config.AllowedHeadersCORS,
		AllowedMethods:   config.AllowedMethodsCORS,
		AllowCredentials: true,
	}).Handler(router)

	server := http.Server{
		Addr:         config.Port,
		Handler:      handler,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
	return server, nil
}

func initDataBase() (postgresSetup.DataBase, error) {
	db := postgresSetup.DataBase{}
	_, err := db.Init(config.DbUser, config.DbPassword, config.DbDB)
	if err != nil {
		return postgresSetup.DataBase{}, err
	}
	return db, nil
}

func initRouters(db postgresSetup.DataBase, logger logger.Interface) *mux.Router {
	re := postgres.New(db.DB, logger)
	fi := system.New(config.Path, logger)
	uc := usecase.New(re, fi, logger)
	de := delivery.New(uc, logger)

	router := mux.NewRouter()
	router.HandleFunc("/upload", de.Upload)
	router.HandleFunc("/download/{Id}", de.Download)
	router.HandleFunc("/download/{Id}/{From}", de.Download)
	return router
}

func cliArgSetup() {
	f := Flags{}
	f.port = flag.Int("port", 8080, "port to start server")
	flag.Parse()
	if f.port != nil {
		config.Port = ":" + strconv.Itoa(*f.port)
	}
}
