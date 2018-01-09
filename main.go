package main

import (
	"github.com/apprentice3d/forge-api-go-client/recap"
	"os"
	"log"
	"github.com/apprentice3d/forge-api-go-client/oauth"
)

const PORT = ":80"

//const photosceneid = "iF8SC83xFVCBFd9W3hSyISB06c7HnEX12xVNlGeBWTQ"

func main() {
	//server.StartServer(PORT)
	clientId := os.Getenv("FORGE_CLIENT_ID")
	clientSecret := os.Getenv("FORGE_CLIENT_SECRET")

	client := oauth.NewTwoLeggedClient(clientId, clientSecret)
	token, err := client.Authenticate("data:write")
	if err != nil {
		log.Fatalln(err.Error())
	}

	recapApi := recap.NewReCapAPIWithCredentials(clientId, clientSecret)
	photoScene, err := recapApi.CreatePhotoScene("testare", nil)

	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("Created a scene with ID: %s\n", photoScene.ID)

	files_to_upload := []string{"C:\\Temp\\mycollection.png",
		//"image_samples/DSC_1158.JPG",
		//"image_samples/DSC_1159.JPG",
		//"image_samples/DSC_1160.JPG",
		//"image_samples/DSC_1162.JPG",
		//"image_samples/DSC_1163.JPG",
		//"image_samples/DSC_1164.JPG",
		//"image_samples/DSC_1165.JPG",
	}

	someUrl := "https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1160.JPG"

	for _, file := range(files_to_upload) {
		result, err := recap.AddFileToScene(photoScene.ID, file, token.AccessToken)
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Printf("%v\n", result)
		}
	}

	result, err := recap.AddFileToScene(photoScene.ID, someUrl, token.AccessToken)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("%v\n", result)
	}

	resultDelete, err := recap.DeleteScene(photoScene.ID, token.AccessToken)

	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("deleted scene with id  = %s, %s", resultDelete.Resource, resultDelete.Message)
	}
}
