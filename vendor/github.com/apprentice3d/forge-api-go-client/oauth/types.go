package oauth

import cfg "github.com/apprentice3d/forge-api-go-client/config"

type Bearer struct {
	// Will always be Bearer
	TokenType string `json:"token_type"`

	// Access token expiration time (in seconds)
	ExpiresIn int32 `json:"expires_in"`

	// The access token
	AccessToken string `json:"access_token"`

	// The refresh token
	RefreshToken string `json:"refresh_token,omitempty"`
}

type TwoLeggedAuthenticator interface {
	Authenticate(scope string) (*Bearer, error)
	GetConfig() cfg.Configuration
}

type ThreeLeggedAuthenticator interface {
	Authorize(scope string, state string) (string, error)
	Gettoken(code string) (*Bearer, error)
	RefreshToken(refreshToken string, scope string) (*Bearer, error)
}

type AuthApi struct {
	cfg.Configuration
	AuthPath string
}

type UserProfile struct {
	UserId    string `json:"userId"`    // The backend user ID of the profile
	UserName  string `json:"userName"`  // The username chosen by the user
	EmailId   string `json:"emailId"`   // The user’s email address
	FirstName string `json:"firstName"` // The user’s first name
	LastName  string `json:"lastName"`  // The user’s last name
	// true if the user’s email address has been verified false if the user’s email address has not been verified
	EmailVerified bool `json:"emailVerified"`
	// true if the user has enabled two-factor authentication false if the user has not enabled two-factor authentication
	Var2FaEnabled bool `json:"2FaEnabled"`
	// A flat JSON object of attribute-value pairs in which the attributes specify available profile image sizes in the
	// format sizeX<pixels> (where <pixels> is an integer that represents both height and width in pixels of square profile images) and the values are URLs for downloading the images via HTTP
	ProfileImages interface{} `json:"profileImages"`
}

type InformationalApi struct {
	cfg.Configuration
}
