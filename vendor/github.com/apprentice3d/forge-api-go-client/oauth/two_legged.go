package oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// TwoLeggedAuth struct holds data necessary for making requests in 2-legged context
type TwoLeggedAuth struct {
	AuthData
}

// TwoLeggedAuthenticator interface defines the method necessary to qualify as 2-legged authenticator
type TwoLeggedAuthenticator interface {
	Authenticate(scope string) (Bearer, error)
}

// NewTwoLeggedClient returns a 2-legged authenticator with default host and authPath
func NewTwoLeggedClient(clientID, clientSecret string) TwoLeggedAuth {
	return TwoLeggedAuth{
		AuthData{
			clientID,
			clientSecret,
			"https://developer.api.autodesk.com",
			"/authentication/v1",
		},
	}
}

// Authenticate allows getting a token with a given scope
func (a TwoLeggedAuth) Authenticate(scope string) (bearer Bearer, err error) {

	task := http.Client{}

	body := url.Values{}
	body.Add("client_id", a.ClientID)
	body.Add("client_secret", a.ClientSecret)
	body.Add("grant_type", "client_credentials")
	body.Add("scope", scope)

	req, err := http.NewRequest("POST",
		a.Host+a.AuthPath+"/authenticate",
		bytes.NewBufferString(body.Encode()),
	)

	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
	err = decoder.Decode(&bearer)

	return
}
