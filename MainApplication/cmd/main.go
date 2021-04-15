package main

import (
	"HttpBigFilesServer/MainApplication/config"
	"HttpBigFilesServer/MainApplication/internal/files/delivery"
	"HttpBigFilesServer/MainApplication/internal/files/repository/postgres"
	"HttpBigFilesServer/MainApplication/internal/files/usecase"
	"HttpBigFilesServer/MainApplication/internal/postgresSetup"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func main(){
	db := postgresSetup.DataBase{}
	_, err:=db.Init(config.DbUser, config.DbPassword, config.DbDB)
	if err !=nil{
		print(err.Error())
	}
	re:= postgres.New(db.DB)
	uc:=usecase.New(re)
	de:=delivery.New(uc)

	router := mux.NewRouter()
	router.HandleFunc("/upload", de.Upload)
	router.HandleFunc("/download/{Id:.*}", de.Download)

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
	fmt.Println("starting Main at ", config.Port)
	err = server.ListenAndServe()
	if err!=nil{
		fmt.Println(err.Error())
	}
}
