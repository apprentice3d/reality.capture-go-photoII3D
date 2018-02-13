package recap

import (
	"errors"

	"github.com/apprentice3d/forge-api-go-client/oauth"
)

func NewReCapAPIWithCredentials(ClientID string, ClientSecret string) ReCapAPI {
	recapAPI := ReCapAPI{}
	recapAPI.BasePath = "/photo-to-3d/v1"
	return ReCapAPI{
		oauth.NewTwoLeggedClient(ClientID, ClientSecret),
		"/photo-to-3d/v1",
	}
}

// CreatePhotoScene is used to prepare a scene with a given name, expected output formats and sceneType
// 	name - should not be empty
// 	formats - should be of type rcm, rcs, obj, ortho or report
// 	sceneType - should be either "aerial" or "object"
func (api ReCapAPI) CreatePhotoScene(name string, formats []string, sceneType string) (scene PhotoScene, err error) {

	bearer, err := api.Authenticate("data:write")
	if err != nil {
		return
	}
	path := api.Host + api.ReCapPath
	scene, err = CreatePhotoScene(path, name, formats, sceneType, bearer.AccessToken)

	return
}

func (api ReCapAPI) AddFileToSceneUsingLinks(sceneID string, links []string) (uploads LinksUploadingReply, err error) {
	bearer, err := api.Authenticate("data:write")
	if err != nil {
		return
	}
	path := api.Host + api.ReCapPath

	uploads, err = AddFileToSceneUsingLinks(path, sceneID, links, bearer.AccessToken)
	return
}


func (api ReCapAPI) AddFilesToSceneUsingData(sceneID string, data []byte) (uploads FileUploadingReply, err error) {
	bearer, err := api.Authenticate("data:write")
	if err != nil {
		return
	}
	path := api.Host + api.ReCapPath

	uploads, err = AddFileToSceneUsingFileData(path, sceneID, data, bearer.AccessToken)

	return
}

func (api ReCapAPI) StartSceneProcessing(sceneID string) (result string, err error) {
	bearer, err := api.Authenticate("data:write")
	if err != nil {
		return
	}
	path := api.Host + api.ReCapPath
	result, err = StartSceneProcessing(path, sceneID, bearer.AccessToken)
	return
}

func (api ReCapAPI) GetSceneProgress(sceneID string) (progress SceneProgressReply, err error) {
	bearer, err := api.Authenticate("data:read")
	if err != nil {
		return
	}
	path := api.Host + api.ReCapPath
	progress, err = GetSceneProgress(path, sceneID, bearer.AccessToken)
	return
}

func (api ReCapAPI) GetSceneResults(sceneID string, format string) (result SceneResultReply, err error) {
	bearer, err := api.Authenticate("data:read")
	if err != nil {
		return
	}
	path := api.Host + api.ReCapPath
	result, err = GetScene(path, sceneID, bearer.AccessToken, format)
	return
}

func (api ReCapAPI) CancelSceneProcessing(scene PhotoScene) (sceneID string, err error) {
	err = errors.New("method not implemented")
	return
}

func (api ReCapAPI) DeleteScene(scene PhotoScene) (sceneID string, err error) {
	bearer, err := api.Authenticate("data:write")
	if err != nil {
		return
	}
	path := api.Host + api.ReCapPath
	_, err = DeleteScene(path, scene.ID, bearer.AccessToken)
	sceneID = scene.ID
	return
}
