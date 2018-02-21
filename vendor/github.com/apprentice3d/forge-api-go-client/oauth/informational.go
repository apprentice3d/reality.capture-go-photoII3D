package oauth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// UserProfile reflects the response received when query the profile of an authorizing end user in a 3-legged context
type UserProfile struct {
	UserID    string `json:"userId"`    // The backend user ID of the profile
	UserName  string `json:"userName"`  // The username chosen by the user
	EmailID   string `json:"emailId"`   // The user’s email address
	FirstName string `json:"firstName"` // The user’s first name
	LastName  string `json:"lastName"`  // The user’s last name
	// true if the user’s email address has been verified false if the user’s email address has not been verified
	EmailVerified bool `json:"emailVerified"`
	// true if the user has enabled two-factor authentication false if the user has not enabled two-factor authentication
	Var2FaEnabled bool `json:"2FaEnabled"`
	// A flat JSON object of attribute-value pairs in which the attributes specify available profile image sizes in the
	// format sizeX<pixels> (where <pixels> is an integer that represents both height and width in pixels of square
	// profile images) and the values are URLs for downloading the images via HTTP
	ProfileImages interface{} `json:"profileImages"`
}

// Information struct is holding the host and path used when making queries
// for profile of an authorizing end user in a 3-legged context
type Information struct {
	Host        string `json:"host,omitempty"`
	ProfilePath string `json:"profile_path"`
}

// NewInformationQuerier returns an Informational API accessor with default host and profilePath
func NewInformationQuerier() Information {
	return Information{
		"https://developer.api.autodesk.com",
		"/userprofile/v1/users/@me",
	}
}

//AboutMe is used to get the profile of an authorizing end user, given the token obtained via 3-legged OAuth flow
func (a Information) AboutMe(token string) (profile UserProfile, err error) {

	requestPath := a.Host + a.ProfilePath
	task := http.Client{}

	req, err := http.NewRequest("GET",
		requestPath,
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
	err = decoder.Decode(&profile)

	return
}
