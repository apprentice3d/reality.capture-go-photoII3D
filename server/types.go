package server


// CreateSceneRequest reflects data send by client, upon which to create requests to ReCap API
type CreateSceneRequest struct {
	SceneName string	`json:"scene_name"`
	OutputFormats []string	`json:"output_formats"`
	SceneType string	`json:"scene_type"`
}

type ImageSendRequest struct {
	SceneID string `json:"scene_id"`
	ImageList []ImageSendContent	`json:"image_list"`
}

type ImageSendContent struct {
	ID int16	`json:"id"`
	ImageURI string	`json:"image_uri"`
}


type SceneIDContent struct {
	SceneID string `json:"scene_id"`
}

type ResultRequest struct {
	SceneID string `json:"scene_id"`
	Format string	`json:"format"`
}