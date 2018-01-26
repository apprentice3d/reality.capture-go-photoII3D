package server

import (
	"log"
	"net/http"
	"github.com/apprentice3d/forge-api-go-client/oauth"
	"github.com/apprentice3d/forge-api-go-client/recap"
	"encoding/json"
)



type ForgeServices struct {
	oauth oauth.AuthApi
	recap recap.ReCapAPI
}


func (service ForgeServices) getToken(writer http.ResponseWriter, request *http.Request) {

	bearer, err := service.oauth.Authenticate("viewables:read")
	if err != nil {
		writer.WriteHeader(http.StatusNotAcceptable)
		writer.Write([]byte(err.Error()))
		return
	}

	encoder := json.NewEncoder(writer)
	encoder.Encode(bearer)
}



func uploadFiles(writer http.ResponseWriter, request *http.Request) {

	request.ParseMultipartForm(32 << 20)
	for idx, file := range request.MultipartForm.File {
		log.Printf("%s => %v\n", idx, file)
	}

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write([]byte("Files received"))

}