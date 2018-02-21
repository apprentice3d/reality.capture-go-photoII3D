package recap

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func createPhotoScene(path string, name string, formats []string, sceneType string, token string) (scene PhotoScene, err error) {

	if sceneType != "object" && sceneType != "aerial" {
		err = errors.New("the scene type is not supported. Expecting 'object' or 'aerial', got " + sceneType)
		return
	}
	task := http.Client{}

	body := url.Values{}
	body.Add("scenename", name)
	body.Add("format", strings.Join(formats, " "))
	body.Add("scenetype", sceneType)

	req, err := http.NewRequest("POST",
		path+"/photoscene",
		bytes.NewBufferString(body.Encode()),
	)

	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	response, err := task.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		content, _ := ioutil.ReadAll(response.Body)
		err = errors.New("[" + strconv.Itoa(response.StatusCode) + "] " + string(content))
		return
	}

	decoder := json.NewDecoder(response.Body)

	sceneCreationReply := SceneCreationReply{}
	err = decoder.Decode(&sceneCreationReply)

	if err != nil {
		err = errors.New("[JSON DECODING ERROR] " + err.Error())
		return
	}

	// This check is necessary, as there are cases when server returns status OK, but contains an error message
	if bodyError := sceneCreationReply.Error; bodyError != nil {
		err = errors.New("[" + bodyError.Code + "] " + bodyError.Message)
		return
	}

	scene = sceneCreationReply.PhotoScene

	return
}

func addFileToSceneUsingLink(path string, photoSceneID string, link string, token string) (result FileUploadingReply, err error) {

	task := http.Client{}

	params := `photosceneid=` + photoSceneID + `&type=image`
	params += `&file[0]=` + link

	body := strings.NewReader(params)

	req, err := http.NewRequest("POST",
		path+"/file",
		body,
	)
	if err != nil {
		log.Println("could not prepare the request to send links: ", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)
	response, err := task.Do(req)
	if err != nil {
		log.Println("could not send image links: ", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		content, _ := ioutil.ReadAll(response.Body)
		err = errors.New("[" + strconv.Itoa(response.StatusCode) + "] " + string(content))
		return
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&result)

	if err != nil {
		err = errors.New("[JSON DECODING ERROR] " + err.Error())
		return
	}

	// This check is necessary, as there are cases when server returns status OK, but contains an error message
	if bodyError := result.Error; bodyError != nil {
		err = errors.New("[" + bodyError.Code + "] " + bodyError.Message)
		return
	}

	return
}

func addFileToSceneUsingFileData(path string, photoSceneID string, data []byte, token string) (result FileUploadingReply, err error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile("file[0]", "datafile")
	if err != nil {
		log.Println(err.Error())
		return
	}

	dataContent := bytes.NewReader(data);
	if _, err = io.Copy(formFile, dataContent); err != nil {
		log.Println(err.Error())
		return
	}

	writer.WriteField("photosceneid", photoSceneID)
	writer.WriteField("type", "image")
	writer.Close()

	task := http.Client{}

	req, err := http.NewRequest("POST",
		path+"/file",
		body)

	if err != nil {
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := task.Do(req)

	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		content, _ := ioutil.ReadAll(response.Body)
		err = errors.New("[" + strconv.Itoa(response.StatusCode) + "] " + string(content))
		return
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&result)

	if err != nil {
		err = errors.New("[JSON DECODING ERROR] " + err.Error())
		return
	}

	// This check is necessary, as there are cases when server returns status OK, but contains an error message
	if bodyError := result.Error; bodyError != nil {
		err = errors.New("[" + bodyError.Code + "] " + bodyError.Message)
		return
	}

	return
}

func startSceneProcessing(path string, photoSceneID string, token string) (result SceneStartProcessingReply, err error) {
	task := http.Client{}

	req, err := http.NewRequest("POST",
		path+"/photoscene/"+photoSceneID,
		nil,
	)

	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	response, err := task.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		content, _ := ioutil.ReadAll(response.Body)
		err = errors.New("[" + strconv.Itoa(response.StatusCode) + "] " + string(content))
		return
	}


	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&result)

	if err != nil {
		err = errors.New("[JSON DECODING ERROR] " + err.Error())
		return
	}

	// This check is necessary, as there are cases when server returns status OK, but contains an error message
	if bodyError := result.Error; bodyError != nil {
		err = errors.New("[" + bodyError.Code + "] " + bodyError.Message)
		return
	}

	return
}

func getSceneProgress(path string, photoSceneID string, token string) (result SceneProgressReply, err error) {
	task := http.Client{}

	req, err := http.NewRequest("GET",
		path+"/photoscene/"+photoSceneID+"/progress",
		nil,
	)

	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	response, err := task.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		content, _ := ioutil.ReadAll(response.Body)
		err = errors.New("[" + strconv.Itoa(response.StatusCode) + "] " + string(content))
		return
	}


	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&result)

	if err != nil {
		err = errors.New("[JSON DECODING ERROR] " + err.Error())
		return
	}

	// This check is necessary, as there are cases when server returns status OK, but contains an error message
	if bodyError := result.Error; bodyError != nil {
		err = errors.New("[" + bodyError.Code + "] " + bodyError.Message)
		return
	}

	return
}

func getSceneResult(path string, photoSceneID string, token string, format string) (result SceneResultReply, err error) {
	task := http.Client{}

	body := strings.NewReader("format=" + format)

	req, err := http.NewRequest("GET",
		path+"/photoscene/"+photoSceneID,
		body,
	)

	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	response, err := task.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		content, _ := ioutil.ReadAll(response.Body)
		err = errors.New("[" + strconv.Itoa(response.StatusCode) + "] " + string(content))
		return
	}


	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&result)

	if err != nil {
		err = errors.New("[JSON DECODING ERROR] " + err.Error())
		return
	}

	// This check is necessary, as there are cases when server returns status OK, but contains an error message
	if bodyError := result.Error; bodyError != nil {
		err = errors.New("[" + bodyError.Code + "] " + bodyError.Message)
		return
	}

	return
}

func cancelSceneProcessing(path string, photoSceneID string, token string) (result SceneCancelReply, err error) {
	task := http.Client{}

	req, err := http.NewRequest("POST",
		path+"/photoscene/"+photoSceneID + "/cancel",
		nil,
	)

	if err != nil {
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	response, err := task.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		content, _ := ioutil.ReadAll(response.Body)
		err = errors.New("[" + strconv.Itoa(response.StatusCode) + "] " + string(content))
		return
	}


	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(&result)

	if err != nil {
		err = errors.New("[JSON DECODING ERROR] " + err.Error())
		return
	}

	// This check is necessary, as there are cases when server returns status OK, but contains an error message
	if bodyError := result.Error; bodyError != nil {
		err = errors.New("[" + bodyError.Code + "] " + bodyError.Message)
		return
	}

	return

}

func deleteScene(path string, photoSceneID string, token string) (result SceneDeletionReply, err error) {
	task := http.Client{}

	req, err := http.NewRequest("DELETE",
		path+"/photoscene/"+photoSceneID,
		nil,
	)

	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	response, err := task.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		content, _ := ioutil.ReadAll(response.Body)
		err = errors.New("[" + strconv.Itoa(response.StatusCode) + "] " + string(content))
		return
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&result)

	if err != nil {
		err = errors.New("[JSON DECODING ERROR] " + err.Error())
		return
	}

	// This check is necessary, as there are cases when server returns status OK, but contains an error message
	if bodyError := result.Error; bodyError != nil {
		err = errors.New("[" + bodyError.Code + "] " + bodyError.Message)
		return
	}

	return
}
