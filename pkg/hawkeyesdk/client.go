package hawkeyesdk

type ClientSettings struct {
	AuthToken string
	DevEnv    bool
	BaseUrl   string
}

const (
	BASE_URL     = "https://hawkeye.g2it.co/api"
	DEV_BASE_URL = "https://qa.hawkeye.g2it.co/api"
)

func NewHawkeyeClient(authToken string, devEnv bool) ClientSettings {
	var url string
	if devEnv {
		url = DEV_BASE_URL
	} else {
		url = BASE_URL
	}
	client := ClientSettings{
		AuthToken: authToken,
		DevEnv:    devEnv,
		BaseUrl:   url,
	}
	return client
}
