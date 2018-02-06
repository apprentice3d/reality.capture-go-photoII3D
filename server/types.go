package server

import (
	"github.com/apprentice3d/forge-api-go-client/oauth"
	"github.com/apprentice3d/forge-api-go-client/recap"
)

// ForgeServices holds all references necessary to access Forge services
type ForgeServices struct {
	recap.ReCapAPI
}

// CreateSceneRequest reflects structure of data received from client for creating a PhotoScene
type CreateSceneRequest struct {
	SceneName     string   `json:"scene_name"`
	OutputFormats []string `json:"output_formats"`
	SceneType     string   `json:"scene_type"`
}

// ImageSendRequest reflects structure of data received from client sending remote located files to PhotoScene
type ImageSendRequest struct {
	SceneID   string             `json:"scene_id"`
	ImageList []ImageSendContent `json:"image_list"`
}

// ImageSendContent reflects the structure of data associated to each remote located image
// that has to be send to Photoscene
type ImageSendContent struct {
	ID       int16  `json:"id"`
	ImageURI string `json:"image_uri"`
}

// SceneIDContent is useful when the body request from client contains a reference to scene id in that scope
type SceneIDContent struct {
	SceneID string `json:"scene_id"`
}

// ResultRequest reflects the request for the specified format associated to given scene id
type ResultRequest struct {
	SceneID string `json:"scene_id"`
	Format  string `json:"format"`
}

// FrontendReport is used to prepare a JSON response to frontend
type FrontendReport struct {
	Result      string `json:"result"`
	Description string `json:"description"`
}