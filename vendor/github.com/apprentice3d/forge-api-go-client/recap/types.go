package recap

import (
	"github.com/apprentice3d/forge-api-go-client/oauth"
)


type ReCapAPI struct {
	oauth.AuthApi
	ReCapPath string
}


type MetaData struct {
	Name   string
	Values string
}


type PhotoScene struct {
	ID       string     `json:"photosceneid"`
	Name     string     `json:"name,omitempty"`
	Files    []string   `json:",omitempty"`
	Formats  []string   `json:",omitempty"`
	Metadata []MetaData `json:",omitempty"`
}



type SceneCreationReply struct {
	Usage      string     `json:",omitempty"`
	Resource   string     `json:",omitempty"`
	PhotoScene PhotoScene `json:"Photoscene"`
}

type SceneDeletionReply struct {
	Usage    string `json:",omitempty"`
	Resource string `json:",omitempty"`
	Message  string `json:"msg"`
}

type FileUploadingReply struct {
	Usage    string `json:",omitempty"`
	Resource string `json:",omitempty"`
	Files    *struct {
		File struct {
			FileName string `json:"filename"`
			FileID   string `json:"fileid"`
			FileSize string `json:"filesize"`
			Message  string `json:"msg"`
		} `json:"file"`
	} `json:"Files"`
}

type LinksUploadingReply struct {
	Usage    string `json:",omitempty"`
	Resource string `json:",omitempty"`
	Files    *struct {
		File []struct {
			FileName string `json:"filename"`
			FileID   string `json:"fileid"`
			FileSize string `json:"filesize"`
			Message  string `json:"msg"`
		} `json:"file"`
	} `json:"Files"`
	Error    *struct {
		Code    string `json:"code"`
		Message string `json:"msg"`
	} `json:"Error"`
}

type SceneStartProcessingReply struct {
	Message    string     `json:"msg"`
	PhotoScene PhotoScene `json:"Photoscene"`
}

type SceneProgressReply struct {
	Usage      string `json:",omitempty"`
	Resource   string `json:",omitempty"`
	PhotoScene struct {
		ID       string `json:"photosceneid"`
		Message  string `json:"progressmsg"`
		Progress string `json:"progress"`
	} `json:"Photoscene"`
}

/*
{
    "Usage": "1.1370530128479",
    "Resource": "\/photoscene\/NzNyWObtb3uu8Cw2yFawHzW4yP9cNXsultMYlZZQdYs\/progress",
    "Photoscene": {
        "photosceneid": "NzNyWObtb3uu8Cw2yFawHzW4yP9cNXsultMYlZZQdYs",
        "progressmsg": "Created",
        "progress": "0"
    }
}

*/

type SceneResultReply struct {
	PhotoScene struct {
		ID        string `json:"photosceneid"`
		Message   string `json:"progressmsg"`
		Progress  string `json:"progress"`
		SceneLink string `json:"scenelink"`
		FileSize  string `json:"filesize"`
	} `json:"Photoscene"`
}

// ErrorMessage represents a struct corresponding to successfully received task, but failed due to some reasons.
//
// 	Frequently the operation succeeded with returning code 200, meaning that the task was
// 	received successfully, but failed to execute due to reasons specified in message
// 	(g.e. uploading a file by specifying an wrong link: POST request is successful,
// 	but internally it failed to download the file because of the wrongly provided link)
type ErrorMessage struct {
	Usage    string `json:",omitempty"`
	Resource string `json:",omitempty"`
	Error    *struct {
		Code    string `json:"code"`
		Message string `json:"msg"`
	} `json:"Error"`
}
