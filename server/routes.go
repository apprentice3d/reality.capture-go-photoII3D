package server

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/apprentice3d/forge-api-go-client/recap"
	"strconv"
	"io/ioutil"
	"encoding/binary"
	"fmt"
)

func (service ForgeServices) getToken(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	bearer, err := service.Authenticate("viewables:read")
	if err != nil {
		writer.WriteHeader(http.StatusNotAcceptable)
		encoder.Encode(FrontendReport{"NACK", err.Error()})
		return
	}

	err = encoder.Encode(bearer)
	log.Println("ERROR: could not encode bearer: ", err.Error())
}

func (service ForgeServices) createScene(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(FrontendReport{"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var sceneCreationRequest CreateSceneRequest
	err := decoder.Decode(&sceneCreationRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", "Could not parse the body request"})
		return
	}

	log.Printf("createSceneRequest: %v => ", sceneCreationRequest)

	photoScene, err := service.CreatePhotoScene(sceneCreationRequest.SceneName,
		sceneCreationRequest.OutputFormats,
		sceneCreationRequest.SceneType)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", err.Error()})
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
		encoder.Encode(FrontendReport{"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var imageUploadRequest ImageSendRequest
	err := decoder.Decode(&imageUploadRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", "Could not parse the body request"})
		return
	}

	log.Printf("imageSendRequest for %d images => ", len(imageUploadRequest.ImageList))

	uploadStatus := make(map[string]string)

	for _, link := range imageUploadRequest.ImageList {
		_, err := service.AddFileToSceneUsingLink(imageUploadRequest.SceneID, link.ImageURI)
		var result string
		if err != nil {

			result = fmt.Sprintf("[FAIL] with error: %s", err.Error())
		} else {
			result = fmt.Sprintf("[SUCCESS]")
		}

		uploadStatus[link.ImageURI] = result
		log.Printf("[%s] => %s", link.ImageURI, uploadStatus[link.ImageURI])
	}

	if err != nil {

		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", err.Error()})
		return
	}

	log.Printf("Done imageSendRequest.\n")

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder.Encode(struct {
		Result string                    `json:"result"`
		Reply  map[string]string `json:"reply"`
	}{"ACK", uploadStatus})

	return
}


func (service ForgeServices) uploadFiles(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "*")
	encoder := json.NewEncoder(writer)

	request.ParseMultipartForm(32 << 20)
	file, _, err := request.FormFile("file")
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", "Could not get the file from request"})
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", "Could not parse the body request"})
		return
	}
	defer request.Body.Close()

	if binary.Size(data) == 0 {
		log.Printf("Frontend is just checking the server availability")
		writer.WriteHeader(http.StatusOK)
		return
	}

	sceneID := request.Header.Get("sceneid")

	log.Printf("imageUploadRequest of size %v for sceneID= %s\n", binary.Size(data), sceneID)

	result, err := service.AddFileToSceneUsingData(sceneID, data)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", err.Error()})
		return
	}

	log.Printf("Done uploadingFileRequest for sceneID=%s", sceneID)

	encoder.Encode(struct {
		SceneID string                   `json:"scene_id"`
		Data    recap.FileUploadingReply `json:"data"`
	}{sceneID,
		result,
	})

	return

}

func (service ForgeServices) startProcessing(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(FrontendReport{"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var sceneStartRequest SceneIDContent
	err := decoder.Decode(&sceneStartRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", "Could not parse the body request"})
		return
	}

	log.Printf("sceneStartRequest: %v => ", sceneStartRequest)

	result, err := service.StartSceneProcessing(sceneStartRequest.SceneID)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", err.Error()})
		return
	}

	log.Printf("Done sceneStartRequest. Started for ID=%s", result.PhotoScene.ID)

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder.Encode(struct {
		Result  string `json:"result"`
		SceneID string `json:"scene_id"`
	}{"ACK", result.PhotoScene.ID})

	return
}

func (service ForgeServices) checkProgress(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(FrontendReport{"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var sceneProgressRequest SceneIDContent
	err := decoder.Decode(&sceneProgressRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", "Could not parse the body request"})
		return
	}

	progress, err := service.GetSceneProgress(sceneProgressRequest.SceneID)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", err.Error()})
		return
	}

	value, err := strconv.Atoi(progress.PhotoScene.Progress)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", err.Error()})
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
		encoder.Encode(FrontendReport{"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var resultRequest ResultRequest
	err := decoder.Decode(&resultRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", "Could not parse the body request"})
		return
	}

	log.Printf("resultRequest: %v\n", resultRequest)

	result, err := service.GetSceneResults(resultRequest.SceneID, resultRequest.Format)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(FrontendReport{"NACK", err.Error()})
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


