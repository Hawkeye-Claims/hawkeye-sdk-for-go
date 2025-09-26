package hawkeyesdk

import "net/http"

type ClientSettings struct {
	AuthToken  string
	DevEnv     bool
	BaseUrl    string
	HTTPClient *http.Client
}

type Environment string

type Option func(*ClientSettings)

const (
	PROD Environment = "prod"
	DEV  Environment = "dev"
)

func WithEnvironment(env Environment) Option {
	return func(c *ClientSettings) {
		if env == DEV {
			c.DevEnv = true
			c.BaseUrl = "https://qa.hawkeye.g2it.co/api"
		} else {
			c.DevEnv = false
			c.BaseUrl = "https://hawkeye.g2it.co/api"
		}
	}
}

func NewHawkeyeClient(authToken string, opts ...Option) *ClientSettings {
	var url string
	url = "https://hawkeye.g2it.co/api"

	httpClient := &http.Client{}
	client := ClientSettings{
		AuthToken:  authToken,
		BaseUrl:    url,
		HTTPClient: httpClient,
	}

	for _, opt := range opts {
		opt(&client)
	}

	return &client
}
