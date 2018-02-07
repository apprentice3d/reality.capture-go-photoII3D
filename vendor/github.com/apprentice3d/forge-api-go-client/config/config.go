// Package config contains the Configuration definition
// along with default Config construction function.
//	Customizing the BasePath, Host and RedirectURI fields might be useful for cases like mocking, proxies etc
package config

// Configuration struct is responsible for aggregating data
// used by different Forge API in context of Go Forge SDK
type Configuration struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	BasePath     string `json:"basePath,omitempty"`
	Host         string `json:"host,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
}

//NewConfiguration returns a default configuration with:
//	- Host = "https://developer.api.autodesk.com"
//	- BaseBath = "/"
//	- RedirectURI = "http://localhost:3000/callback"
func NewConfiguration() Configuration {
	cfg := Configuration{
		Host:        "https://developer.api.autodesk.com",
		BasePath:    "/",
		RedirectURI: "http://localhost:3000/callback",
	}

	return cfg
}
