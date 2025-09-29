package hawkeyesdk

import "net/http"

type ClientSettings struct {
	AuthToken  string
	DevEnv     bool
	BaseUrl    string
	HTTPClient *http.Client

	// Services
	Claims    *ClaimsService
	DocFiles  *DocFilesService
	LogTrails *LogTrailsService
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
	client := ClientSettings{
		AuthToken: authToken,
		BaseUrl:   "https://hawkeye.g2it.co/api",
	}

	for _, opt := range opts {
		opt(&client)
	}

	client.ensureHTTPClient()
	client.initServices()

	return &client
}

func (cfg *ClientSettings) ensureHTTPClient() {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{}
	}
}

func (cfg *ClientSettings) initServices() {
	cfg.Claims = NewClaimsService(cfg)
	cfg.DocFiles = NewDocFilesService(cfg)
	cfg.LogTrails = NewLogTrailsService(cfg)
}
