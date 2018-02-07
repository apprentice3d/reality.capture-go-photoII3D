package oauth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	cfg "github.com/apprentice3d/forge-api-go-client/config"
)

func NewInformationalApi() InformationalApi {
	configuration := cfg.NewConfiguration()
	return InformationalApi{
		configuration,
	}
}

//func NewInformationalApiWithBasePath(basePath string) InformationalApi {
//	configuration := cfg.NewConfiguration()
//	configuration.BasePath = basePath
//
//	return InformationalApi{
//		Configuration: configuration,
//	}
//}

/**
 * GET users/@me
 * GET users/@me
 *
 * @return *UserProfile
 */
func (a InformationalApi) AboutMe(token Bearer) (*UserProfile, error) {

	requestPath := a.Host + a.BasePath + "/userprofile/v1/users/@me"

	task := http.Client{}

	req, err := http.NewRequest("GET",
		requestPath,
		nil,
	)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	response, err := task.Do(req)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		err = errors.New("[" + strconv.Itoa(response.StatusCode) + "] " + string(content))
		return nil, err
	}
	information := new(UserProfile)
	err = json.Unmarshal(content, &information)

	return information, err
}
