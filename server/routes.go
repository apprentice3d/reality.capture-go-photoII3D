package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/apprentice3d/forge-api-go-client/recap"
	"strconv"
)

func (service ForgeServices) getToken(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	bearer, err := service.Authenticate("viewables:read")
	if err != nil {
		writer.WriteHeader(http.StatusNotAcceptable)
		encoder.Encode(FrontendReport {"NACK", err.Error()})
		return
	}

	err = encoder.Encode(bearer)
	log.Println("ERROR: could not encode bearer: ", err.Error())
}

func (service ForgeServices) createScene(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(FrontendReport {"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var sceneCreationRequest CreateSceneRequest
	err := decoder.Decode(&sceneCreationRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport {"NACK", "Could not parse the body request"})
		return
	}

	log.Printf("createSceneRequest: %v => ", sceneCreationRequest)

	photoScene, err:= service.CreatePhotoScene(sceneCreationRequest.SceneName,
		sceneCreationRequest.OutputFormats,
		sceneCreationRequest.SceneType)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport {"NACK", err.Error()})
		return
	}

	log.Printf("Done createSceneRequest. SceneID=%s\n", photoScene.ID)

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder.Encode(struct {
		Result  string `json:"result"`
		SceneID string `json:"scene_id"`
	}{"ACK", photoScene.ID})

	return
}


func (service ForgeServices) sendFiles(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(FrontendReport {"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var imageUploadRequest ImageSendRequest
	err := decoder.Decode(&imageUploadRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport {"NACK", "Could not parse the body request"})
		return
	}

	imageLinks := []string{}

	for _, link := range imageUploadRequest.ImageList {
		imageLinks = append(imageLinks, link.ImageURI)
	}

	log.Printf("imageSendRequest for %d images => ", len(imageLinks))

	uploadResult, err:= service.AddFileToSceneUsingLinks(imageUploadRequest.SceneID, imageLinks)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport {"NACK", err.Error()})
		return
	}

	log.Printf("Done imageSendRequest.\n")

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder.Encode(struct {
		Result string `json:"result"`
		Reply recap.LinksUploadingReply `json:"reply"`
	}{"ACK", uploadResult})

	return
}

func (service ForgeServices) startProcessing(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(FrontendReport {"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var sceneStartRequest SceneIDContent
	err := decoder.Decode(&sceneStartRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport {"NACK", "Could not parse the body request"})
		return
	}

	log.Printf("sceneStartRequest: %v => ", sceneStartRequest)

	result, err := service.StartSceneProcessing(sceneStartRequest.SceneID)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport {"NACK", err.Error()})
		return
	}

	log.Printf("Done sceneStartRequest. Started for ID=%s", result)

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder.Encode(struct {
		Result string `json:"result"`
		SceneID string `json:"scene_id"`
	}{"ACK", result})

	return
}

func (service ForgeServices) checkProgress(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(FrontendReport {"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var sceneProgressRequest SceneIDContent
	err := decoder.Decode(&sceneProgressRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport {"NACK", "Could not parse the body request"})
		return
	}

	progress, err := service.GetSceneProgress(sceneProgressRequest.SceneID)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport {"NACK", err.Error()})
		return
	}


	value, err := strconv.Atoi(progress.PhotoScene.Progress)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport {"NACK", err.Error()})
		return
	}

	log.Printf("sceneProgressRequest: %v => returning progress = %d\n", sceneProgressRequest, value)

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder.Encode(struct {
		SceneID  string `json:"scene_id"`
		Progress int    `json:"progress"`
	}{sceneProgressRequest.SceneID, value})

	return
}

func (service ForgeServices) getResult(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(FrontendReport {"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var resultRequest ResultRequest
	err := decoder.Decode(&resultRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport {"NACK", "Could not parse the body request"})
		return
	}

	sampleSceneID := "qQwB6HacvPCXnb8VXO0PQDsqPzyGw8JPxCB79XmjcPs"
	log.Println(sampleSceneID)

	log.Printf("resultRequest: %v\n", resultRequest)

	result, err := service.GetSceneResults(resultRequest.SceneID, resultRequest.Format)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport {"NACK", err.Error()})
		return
	}

	log.Printf("Done resultRequest. Returning link %s\n", result.PhotoScene.SceneLink)

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder.Encode(struct {
		SceneID string `json:"scene_id"`
		Format  string `json:"format"`
		Link    string `json:"link"`
	}{resultRequest.SceneID,
		resultRequest.Format,
		result.PhotoScene.SceneLink})

	return
}





func (service ForgeServices) uploadFiles(writer http.ResponseWriter, request *http.Request) {

	file, err := os.Create("localfile.jpg")

	if err != nil {
		log.Fatal(err.Error())
	}

	size, err := io.Copy(file, request.Body)

	//data, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Received file")
	log.Println(size)
	file.Close()

	//err := request.ParseMultipartForm(32 << 20)
	//if err != nil {
	//	log.Println(err.Error())
	//	writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "*")
	writer.Write([]byte("Files received"))

	//request.ParseMultipartForm(32 << 20)
	//for idx, file := range request.MultipartForm.File {
	//	log.Printf("%s => %v\n", idx, file)
	//}
	//
	//writer.Header().Set("Access-Control-Allow-Origin", "*")
	//writer.Write([]byte("Files received"))

}