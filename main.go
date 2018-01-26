package main

import (

	"github.com/apprentice3d/forge-photoII3D/server"
)

const port = ":80"

func main() {
	server.StartServer(port)
	//clientID := os.Getenv("FORGE_CLIENT_ID")
	//clientSecret := os.Getenv("FORGE_CLIENT_SECRET")
	//
	//client := oauth.NewTwoLeggedClient(clientID, clientSecret)
	//token, err := client.Authenticate("data:write")
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//
	//recapAPI := recap.NewReCapAPIWithCredentials(clientID, clientSecret)
	//photoScene, err := recapAPI.CreatePhotoScene("testare", nil)
	//
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//
	//log.Printf("Created a scene with ID: %s\n", photoScene.ID)
	//
	//filesToUpload := []string{"C:\\Temp\\mycollection.png"} //"image_samples/DSC_1158.JPG",
	////"image_samples/DSC_1159.JPG",
	////"image_samples/DSC_1160.JPG",
	////"image_samples/DSC_1162.JPG",
	////"image_samples/DSC_1163.JPG",
	////"image_samples/DSC_1164.JPG",
	////"image_samples/DSC_1165.JPG",
	//
	//someURL := "https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1160.JPG"
	//
	//for _, file := range filesToUpload {
	//	result, err := recap.AddFileToScene(photoScene.ID, file, token.AccessToken)
	//	if err != nil {
	//		log.Println(err.Error())
	//	} else {
	//		log.Printf("%v\n", result)
	//	}
	//}
	//
	//result, err := recap.AddFileToScene(photoScene.ID, someURL, token.AccessToken)
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Printf("%v\n", result)
	//}
	//
	//resultDelete, err := recap.DeleteScene(photoScene.ID, token.AccessToken)
	//
	//if err != nil {
	//	log.Fatalln(err.Error())
	//} else {
	//	log.Printf("deleted scene with id  = %s, %s", resultDelete.Resource, resultDelete.Message)
	//}
}
