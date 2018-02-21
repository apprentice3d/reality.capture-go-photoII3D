package recap

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func TestGeneralUseCase(t *testing.T) {

	// prepare the credentials
	clientID := os.Getenv("FORGE_CLIENT_ID")
	clientSecret := os.Getenv("FORGE_CLIENT_SECRET")

	recapAPI := NewReCapAPIWithCredentials(clientID, clientSecret)

	//get a token
	token, err := recapAPI.Authenticate("data:write data:read")
	if err != nil {
		t.Fatalf("Coud not authorize with provided credentials: %s\n",
			err.Error())
	}

	path := recapAPI.Host + recapAPI.ReCapPath

	var photoSceneID string
	formats := []string{"rcm", "obj"}

	t.Run("Create a scene", func(t *testing.T) {
		photoScene, err := CreatePhotoScene(path, "test_scene", formats, "object", token.AccessToken)

		if err != nil {
			t.Fatalf("Failed to create a photoscene: %s\n", err.Error())
		}
		photoSceneID = photoScene.ID
		//t.Logf("Successfully created a scene with id=%s\n", photoSceneID)

	})

	t.Run("Load remotely located files to created PhotoScene", func(t *testing.T) {
		filesToUpload := []string{
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1158.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1159.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1160.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1162.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1163.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1164.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1165.JPG",
		}


			result, err := AddFileToSceneUsingLinks(path, photoSceneID, filesToUpload, token.AccessToken)
			if err != nil {
				t.Fatalf("Could not upload files: %s\n",  err.Error())
			}
			if result.Files == nil {
				t.Fatalf("Could not upload files: %s\n", err.Error())
			}

	})

	t.Run("Start the PhotoScene processing", func(t *testing.T) {

		sceneID, err := StartSceneProcessing(path, photoSceneID, token.AccessToken)
		if err != nil {
			t.Fatalf("Failed to start PhotoScene processing: %s\n", err.Error())
		}
		if sceneID != photoSceneID {
			t.Fatalf("The received scene id (%s) is not matching the expected (%s)\n",
				sceneID,
				photoSceneID)
		}
	})

	t.Run("Check PhotoScene progress", func(t *testing.T) {

		progress, err := GetSceneProgress(path, photoSceneID, token.AccessToken)
		if err != nil {
			t.Fatalf("Failed to check PhotoScene progress: %s\n", err.Error())
		}
		if progress.PhotoScene.ID != photoSceneID {
			t.Fatalf("The received scene id (%s) is not matching the expected (%s)\n",
				progress.PhotoScene.ID,
				photoSceneID)
		}

	})

	t.Run("Check the finished PhotoScene", func(t *testing.T) {

		//check and make sure the processing is complete
		var progressResult SceneProgressReply
		for {
			if progressResult, err = GetSceneProgress(path, photoSceneID, token.AccessToken); err != nil {
				t.Errorf("Failed to get the PhotoScene progress: %s\n", err.Error())
			}

			ratio, err := strconv.ParseFloat(progressResult.PhotoScene.Progress, 64)

			if err != nil {
				t.Errorf("Failed to parse progress results: %s\n", err.Error())
			}

			if ratio == float64(100.0) {
				break
			}
			time.Sleep(5 * time.Second)
		}

		result, err := GetScene(path, photoSceneID, token.AccessToken, formats[0])
		if err != nil {
			t.Fatalf("Failed to get the PhotoScene results: %s\n", err.Error())
		}
		if result.PhotoScene.ID != photoSceneID {
			t.Fatalf("The received scene id (%s) is not matching the expected (%s)\n",
				result.PhotoScene.ID,
				photoSceneID)
		}
		if len(result.PhotoScene.SceneLink) == 0 {
			t.Errorf("The received result link (%s) is empty\n",
				result.PhotoScene.SceneLink)
		}

		t.Logf(result.PhotoScene.SceneLink)
	})

	t.Run("Deleting the created PhotoScene", func(t *testing.T) {
		_, err := DeleteScene(path, photoSceneID, token.AccessToken)

		if err != nil {
			t.Fatalf("Failed to delete the photoscene: %s\n", err.Error())
		}
	})

}

func TestReCapFreeFunctions(t *testing.T) {

	// prepare the credentials
	clientID := os.Getenv("FORGE_CLIENT_ID")
	clientSecret := os.Getenv("FORGE_CLIENT_SECRET")
	recapAPI := NewReCapAPIWithCredentials(clientID, clientSecret)

	//get a token
	token, err := recapAPI.Authenticate("data:write data:read")
	if err != nil {
		t.Fatalf("Coud not authorize with provided credentials: %s\n",
			err.Error())
	}

	path := recapAPI.Host + recapAPI.ReCapPath

	var photosceneID string

	if err != nil {
		t.Fatalf("Coud not authorize with provided credentials: %s\n",
			err.Error())
	}

	t.Run("Create a scene using a free function", func(t *testing.T) {
		photoScene, err := CreatePhotoScene(path, "testare", nil, "object", token.AccessToken)

		if err != nil {
			t.Fatalf("Failed to create a photoscene: %s\n", err.Error())
		}
		photosceneID = photoScene.ID
	})

	t.Run("Check fail on create a scene with empty name", func(t *testing.T) {
		_, err := CreatePhotoScene(path, "", nil, "object", token.AccessToken)

		if err == nil {
			t.Fatalf("Should fail creating a scene with empty name\n")
		}
	})

	t.Run("Check loading a remote file to created PhotoScene", func(t *testing.T) {
		filesToUpload := []string{
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1158.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1159.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1160.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1162.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1163.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1164.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1165.JPG",
		}


			result, err := AddFileToSceneUsingLinks(path, photosceneID, filesToUpload, token.AccessToken)
			if err != nil {
				t.Fatalf("Could not upload files: %s\n",  err.Error())
			}
			if result.Files == nil {
				t.Fatalf("Could not upload files: %s\n",  err.Error())
			}

	})

	t.Run("Check loading a wrong remote file to created photoscene", func(t *testing.T) {
		fileToUpload := []string{"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_2000.JPG"}

		result, err := AddFileToSceneUsingLinks(path, photosceneID, fileToUpload, token.AccessToken)
		if err == nil {
			t.Fatalf("Uploading a wrong file %s should fail, but it is not!\n", fileToUpload)
		}
		if result.Files != nil {
			t.Fatalf("Uploading a wrong file %s should return an empty content, but it is not!\n", fileToUpload)
		}

	})

	t.Run("Check deleting a photoscene", func(t *testing.T) {
		_, err := DeleteScene(path, photosceneID, token.AccessToken)

		if err != nil {
			t.Fatalf("Failed to delete the photoscene: %s\n", err.Error())
		}
	})
}


func TestAddFileToSceneUsingLinks(t *testing.T) {
	// prepare the credentials
	clientID := os.Getenv("FORGE_CLIENT_ID")
	clientSecret := os.Getenv("FORGE_CLIENT_SECRET")
	recapAPI := NewReCapAPIWithCredentials(clientID, clientSecret)

	//get a token
	token, err := recapAPI.Authenticate("data:write data:read")
	if err != nil {
		t.Fatalf("Coud not authorize with provided credentials: %s\n",
			err.Error())
	}

	path := recapAPI.Host + recapAPI.ReCapPath

	var photosceneID string

	if err != nil {
		t.Fatalf("Coud not authorize with provided credentials: %s\n",
			err.Error())
	}

	t.Run("Create a scene using a free function", func(t *testing.T) {
		photoScene, err := CreatePhotoScene(path, "testare", nil, "object", token.AccessToken)

		if err != nil {
			t.Fatalf("Failed to create a photoscene: %s\n", err.Error())
		}
		photosceneID = photoScene.ID
	})

	t.Run("Check loading a remote file to created PhotoScene", func(t *testing.T) {
		filesToUpload := []string{
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1158.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1159.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1160.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1162.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1163.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1164.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1165.JPG",
		}


		result, err := AddFileToSceneUsingLinks(path, photosceneID, filesToUpload, token.AccessToken)
		if err != nil {
			t.Fatalf("Could not upload files: %s\n",  err.Error())
		}
		if result.Files == nil {
			t.Fatalf("Could not upload files: %s\n",  err.Error())
		}

		t.Logf("[SUCCESS]: Received result: %+v\n", result)

	})

	t.Run("Check loading a remote file to with an non-existing file", func(t *testing.T) {
		filesToUpload := []string{
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1158.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1159.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1160.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_2000.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1163.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1164.JPG",
			"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_1165.JPG",
		}


		result, err := AddFileToSceneUsingLinks(path, photosceneID, filesToUpload, token.AccessToken)
		if err == nil {
			t.Fatalf("Should fail as it contains a wrong file: %s\n",  err.Error())
		}
		if result.Error == nil {
			t.Fatalf("Should contain an error, but it does not: %s\n",  err.Error())
		}

		t.Logf("[SUCCESS]: Received an error as expected: %v\n", result.Error)

	})

	t.Run("Check loading a wrong remote file to created photoscene", func(t *testing.T) {
		fileToUpload := []string{"https://s3.amazonaws.com/adsk-recap-public/forge/lion/DSC_2000.JPG"}

		result, err := AddFileToSceneUsingLinks(path, photosceneID, fileToUpload, token.AccessToken)
		if err == nil {
			t.Logf("[SHOULD FAIL]: Received result: %+v\n", result)
			t.Fatalf("Uploading a wrong file %s should fail, but it is not!\n", fileToUpload)
		}
		if result.Files != nil {
			t.Fatalf("Uploading a wrong file %s should return an empty content, but it is not!\n", fileToUpload)
		}

	})

	t.Run("Check deleting a photoscene", func(t *testing.T) {
		_, err := DeleteScene(path, photosceneID, token.AccessToken)

		if err != nil {
			t.Fatalf("Failed to delete the photoscene: %s\n", err.Error())
		}
	})
}
