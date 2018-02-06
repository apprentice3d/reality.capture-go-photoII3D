package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"math/rand"
)

func (service ForgeServices) getToken(writer http.ResponseWriter, request *http.Request) {

	bearer, err := service.oauth.Authenticate("viewables:read")
	if err != nil {
		writer.WriteHeader(http.StatusNotAcceptable)
		_ , err = writer.Write([]byte(err.Error()))
		return
	}

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(bearer)
	log.Println("ERROR: could not encode bearer: ", err.Error())
}

func (service ForgeServices) createScene(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(struct {
			Result      string `json:"result"`
			Description string `json:"description"`
		}{"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var sceneCreationRequest CreateSceneRequest
	err := decoder.Decode(&sceneCreationRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(struct {
			Result      string `json:"result"`
			Description string `json:"description"`
		}{"NACK", "Could not parse the body request"})
		return
	}

	//TODO: implement sending create scene request

	log.Printf("createSceneRequest: %v\n", sceneCreationRequest)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder.Encode(struct {
		Result  string `json:"result"`
		SceneID string `json:"scene_id"`
	}{"ACK", "placeholder for status id"})

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

func (service ForgeServices) sendFiles(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(struct {
			Result      string `json:"result"`
			Description string `json:"description"`
		}{"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var imageUploadRequest ImageSendRequest
	err := decoder.Decode(&imageUploadRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(struct {
			Result      string `json:"result"`
			Description string `json:"description"`
		}{"NACK", "Could not parse the body request"})
		return
	}

	//TODO: implement sending image upload request

	log.Printf("imageSendRequest: %v\n", imageUploadRequest)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder.Encode(struct {
		Result string `json:"result"`
	}{"ACK"})

	return
}

func (service ForgeServices) startProcessing(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(struct {
			Result      string `json:"result"`
			Description string `json:"description"`
		}{"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var sceneStartRequest SceneIDContent
	err := decoder.Decode(&sceneStartRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(struct {
			Result      string `json:"result"`
			Description string `json:"description"`
		}{"NACK", "Could not parse the body request"})
		return
	}

	//TODO: implement starting the image processing

	log.Printf("sceneStartRequest: %v\n", sceneStartRequest)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder.Encode(struct {
		Result string `json:"result"`
	}{"ACK"})

	return
}

func (service ForgeServices) checkProgress(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(struct {
			Result      string `json:"result"`
			Description string `json:"description"`
		}{"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var sceneProgressRequest SceneIDContent
	err := decoder.Decode(&sceneProgressRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(struct {
			Result      string `json:"result"`
			Description string `json:"description"`
		}{"NACK", "Could not parse the body request"})
		return
	}

	//TODO: implement starting the image processing
	progressMimick := rand.Intn(120)

	log.Printf("sceneProgressRequest: %v, returning progress = %d\n", sceneProgressRequest, progressMimick)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder.Encode(struct {
		SceneID  string `json:"scene_id"`
		Progress int    `json:"progress"`
	}{sceneProgressRequest.SceneID, progressMimick})

	return
}

func (service ForgeServices) getResult(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(struct {
			Result      string `json:"result"`
			Description string `json:"description"`
		}{"NACK", "Expecting POST request"})
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var resultRequest ResultRequest
	err := decoder.Decode(&resultRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(struct {
			Result      string `json:"result"`
			Description string `json:"description"`
		}{"NACK", "Could not parse the body request"})
		return
	}

	//TODO: implement starting the image processing

	log.Printf("resultRequest: %v\n", resultRequest)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder.Encode(struct {
		SceneID string `json:"scene_id"`
		Format  string `json:"format"`
		Data    string `json:"data"`
	}{resultRequest.SceneID,
		resultRequest.Format,
		"This is placeholder for result"})

	return
}
