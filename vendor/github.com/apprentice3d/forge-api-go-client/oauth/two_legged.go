package oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	cfg "github.com/apprentice3d/forge-api-go-client/config"
)

func NewTwoLeggedClient(clientId string, clientSecret string) AuthApi {
	configuration := cfg.NewConfiguration()
	configuration.ClientID = clientId
	configuration.ClientSecret = clientSecret
	return AuthApi{
		configuration,
		"/authentication/v1",
	}
}

//func NewTwoLeggedClientWithBasePath(clientId string, clientSecret string, basePath string) *TwoLeggedApi {
//	client := NewTwoLeggedClient(clientId, clientSecret)
//	client.BasePath = basePath
//	return client
//}

/**
 * POST authenticate
 * Get a two-legged access token by providing the needed scope.
 *
 * @param scope Space-separated list of required scopes
 * @return *Bearer and error
 */
func (a AuthApi) Authenticate(scope string) (*Bearer, error) {

	task := http.Client{}

	body := url.Values{}
	body.Add("client_id", a.Configuration.ClientID)
	body.Add("client_secret", a.Configuration.ClientSecret)
	body.Add("grant_type", "client_credentials")
	body.Add("scope", scope)

	req, err := http.NewRequest("POST",
		a.Host+a.AuthPath+"/authenticate",
		bytes.NewBufferString(body.Encode()),
	)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
	token := new(Bearer)
	err = json.Unmarshal(content, &token)

	return token, err
}

func (a AuthApi) GetConfig() cfg.Configuration {
	return a.Configuration
}
