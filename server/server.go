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
		oauth: setupForgeOAuth(),
	}

	//serving static files
	fs := http.FileServer(http.Dir("client/build"))
	http.Handle("/", fs)

	// routes
	http.HandleFunc("/create_scene", service.createScene)
	http.HandleFunc("/upload_remote_images", service.sendFiles)
	http.HandleFunc("/upload_local_images", service.uploadFiles)
	http.HandleFunc("/start_process", service.startProcessing)
	http.HandleFunc("/check_progress", service.checkProgress)
	http.HandleFunc("/get_result", service.getResult)
	http.HandleFunc("/get_token", service.getToken)

	log.Println("Starting server on port " + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err.Error())
	}
}

func setupForgeOAuth() oauth.AuthApi {
	clientID := os.Getenv("FORGE_CLIENT_ID")
	clientSecret := os.Getenv("FORGE_CLIENT_SECRET")

	if len(clientID) == 0 || len(clientSecret) == 0 {
		log.Fatal("The FORGE_CLIENT_ID and FORGE_CLIENT_SECRET env vars are not set. \nExiting ...")
	}

	return oauth.NewTwoLeggedClient(clientID, clientSecret)
}
