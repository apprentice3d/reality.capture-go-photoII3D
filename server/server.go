package server

import (
	"log"
	"net/http"
	"os"

	"github.com/apprentice3d/forge-api-go-client/oauth"
)


//StartServer is responsible for setting up and lunching a simple web-server on the specified port
func StartServer(port string) {

	service := ForgeServices{
		oauth:setupForgeOAuth(),
	}

	fs := http.FileServer(http.Dir("client/build"))
	http.Handle("/", fs)

	// routes
	http.HandleFunc("/uploadLocalImages", uploadFiles)
	http.HandleFunc("/gettoken", service.getToken)






	log.Println("Starting server on port " + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err.Error())
	}
}


func setupForgeOAuth() oauth.AuthApi {
	clientID := os.Getenv("FORGE_CLIENT_ID")
	clientSecret := os.Getenv("FORGE_CLIENT_SECRET")

	return oauth.NewTwoLeggedClient(clientID, clientSecret)
}
