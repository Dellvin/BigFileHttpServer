package tests

import (
	"HttpBigFilesServer/MainApplication/pkg"
	"log"
	"testing"
)

func TestRandom(t *testing.T) {
	_, err := pkg.GenId()
	if err != nil {
		log.Fatalf(err.Error())
	}
	_, err = pkg.GenId()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
