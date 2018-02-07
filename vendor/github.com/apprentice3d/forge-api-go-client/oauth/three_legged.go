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

func NewThreeLeggedApi(redirectUri string) AuthApi {
	configuration := cfg.NewConfiguration()
	configuration.RedirectURI = redirectUri
	return AuthApi{
		configuration,
		"/authentication/v1",
	}
}

//func NewThreeLeggedApiWithBasePath(redirectUri string, basePath string) *ThreeLeggedApi {
//	configuration := cfg.NewConfiguration()
//	configuration.BasePath = basePath
//	configuration.RedirectURI = redirectUri
//	return &ThreeLeggedApi{
//		configuration,
//	}
//}

/**
 * GET authorize
 * This is the browser URL to redirect an end user to in order to acquire the user’s consent for your app to access the specified resources. Note: You do not call this URL directly in your server code. See the Get a 3-Legged Token tutorial for more information on how to use this endpoint.
 *
 * @param clientId Client ID of the app
 * @param responseType Must be code
 * @param RedirectUri URL-encoded callback URL that the end user will be redirected to after completing the authorization flow Note: This must match the pattern of the callback URL field of the app’s registration in the My Apps section. The pattern may include wildcards after the hostname, allowing different redirect_uri values to be specified in different parts of your app.
 * @param scope Space-separated list of required scopes Note: A URL-encoded space is* &#x60;&#x60;%20&#x60;&#x60;. See the* &#x60;Scopes &lt;/en/docs/oauth/v2/overview/scopes&gt;&#x60; *page for more information on when scopes are required.
 * @param state A URL-encoded payload containing arbitrary data that the authentication flow will pass back verbatim in a state query parameter to the callback URL
 * @return void
 */
func (a AuthApi) Authorize(scope string, state string) (string, error) {

	request, err := http.NewRequest("GET",
		a.Host+a.AuthPath+"/authorize",
		nil,
	)

	if err != nil {
		return "", err
	}

	query := request.URL.Query()
	query.Add("client_id", a.Configuration.ClientID)
	query.Add("response_type", "code")
	query.Add("redirect_uri", a.Configuration.RedirectURI)
	query.Add("scope", scope)
	query.Add("state", state)

	request.URL.RawQuery = query.Encode()

	return request.URL.String(), nil
}

/**
 * POST gettoken
 * Exchange an authorization code extracted from a GET authorize callback for a three-legged access token.
 *
 * @param clientId Client ID of the app
 * @param clientSecret Client secret of the app
 * @param grantType Must be &#x60;&#x60;authorization_code&#x60;&#x60;
 * @param code The autorization code captured from the code query parameter when the GET authorize redirected back to the callback URL
 * @param redirectUri URL-encoded callback URL that the end user will be redirected to after completing the authorization flow Note: This must match the pattern of the callback URL field of the app’s registration in the My Apps section. The pattern may include wildcards after the hostname, allowing different redirect_uri values to be specified in different parts of your app.
 * @return *Bearer
 */
func (a AuthApi) Gettoken(code string) (*Bearer, error) {

	task := http.Client{}

	body := url.Values{}
	body.Add("client_id", a.Configuration.ClientID)
	body.Add("client_secret", a.Configuration.ClientSecret)
	body.Add("grant_type", "authorization_code")
	body.Add("code", code)
	body.Add("redirect_uri", a.Configuration.RedirectURI)

	req, err := http.NewRequest("POST",
		a.Host+a.AuthPath+"/gettoken",
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

/**
 * POST refreshtoken
 * Acquire a new access token by using the refresh token provided by the POST gettoken endpoint. See the Field Guide for more information about refresh tokens.
 * @param refreshToken The refresh token used to acquire a new access token
 * @param scope Space-separated list of required scopes If this parameter is omitted, the returned access token will have the same scopes as the original access token. If this parameter is specified, it must represent a subset of the scopes present in the original access token. Note: A URL-encoded space is %20.
 * @return *Bearer error
 */
func (a AuthApi) RefreshToken(refreshToken string, scope string) (*Bearer, error) {

	task := http.Client{}

	body := url.Values{}
	body.Add("client_id", a.Configuration.ClientID)
	body.Add("client_secret", a.Configuration.ClientSecret)
	body.Add("grant_type", "refresh_token")
	body.Add("refresh_token", refreshToken)
	body.Add("scope", scope)

	req, err := http.NewRequest("POST",
		a.Host+a.AuthPath+"/refreshtoken",
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
