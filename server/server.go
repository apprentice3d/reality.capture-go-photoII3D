package server

import (
	"log"
	"net/http"
	"os"

	"github.com/apprentice3d/forge-api-go-client/oauth"
)

var (
	tokenManager *oauth.TwoLeggedApi
)

//StartServer is responsible for setting up and lunching a simple web-server on the specified port
func StartServer(port string) {

	tokenManager = setupForgeOAuth()

	fs := http.FileServer(http.Dir("client/build"))
	http.Handle("/", fs)
	http.HandleFunc("/upload", uploadFiles)

	log.Println("Starting server on port " + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err.Error())
	}
}
func uploadFiles(writer http.ResponseWriter, request *http.Request) {

	request.ParseMultipartForm(32 << 20)
	for idx, file := range request.MultipartForm.File {
		log.Printf("%s => %v\n", idx, file)
	}

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write([]byte("Files received"))

}

func setupForgeOAuth() *oauth.TwoLeggedApi {
	clientID := os.Getenv("FORGE_CLIENT_ID")
	clientSecret := os.Getenv("FORGE_CLIENT_SECRET")

	return oauth.NewTwoLeggedClient(clientID, clientSecret)
}
