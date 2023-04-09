package commonpages

import (
	"fmt"
	"strings"

	"github.com/yunerou/oauth2-client/singleton"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type oauth2Config struct {
	Name              string  `mapstructure:"Name"`
	ClientID          string  `mapstructure:"ClientID"`
	ClientSecret      string  `mapstructure:"ClientSecret"`
	RedirectURL       string  `mapstructure:"RedirectURL"`
	WellknownEndpoint *string `mapstructure:"WellknownEndpoint"`
	AuthURL           *string `mapstructure:"AuthURL"`
	TokenURL          *string `mapstructure:"TokenURL"`
}

func (d *commonPages) registerOauth2Provider() {
	d.oauth2Provider = map[string]*oauth2.Config{}
	viperConfigs := []oauth2Config{}
	err := singleton.GetViper().UnmarshalKey("OAUTH2", &viperConfigs)
	if err != nil {
		panic("config file of OAUTH2 is wrong format")
	}
	fmt.Println("@@@@@@")
	fmt.Println(viperConfigs)
	for _, vConfig := range viperConfigs {
		ep, sc := configEndpointAndScope(
			vConfig.WellknownEndpoint,
			vConfig.AuthURL,
			vConfig.TokenURL,
		)
		oauth2Conf := &oauth2.Config{
			ClientID:     vConfig.ClientID,
			ClientSecret: vConfig.ClientSecret,
			Endpoint:     ep,
			RedirectURL:  vConfig.RedirectURL,
			Scopes:       sc,
		}
		d.oauth2Provider[vConfig.Name] = oauth2Conf
	}
}

func configEndpointAndScope(
	wellknownEndpoint, authURL, tokenURL *string,
) (endpoint oauth2.Endpoint, scope []string) {
	if wellknownEndpoint != nil {
		switch x := strings.ToLower(*wellknownEndpoint); x {
		case "github":
			return github.Endpoint, []string{"user", "repo"}
		case "google":
			return google.Endpoint, []string{"https://www.googleapis.com/auth/userinfo.email"}
		default:
			panic(fmt.Sprintf("well-known oauth2 provider not support %s", x))
		}
	}
	if authURL != nil && tokenURL != nil {
		return oauth2.Endpoint{
			AuthURL:  *authURL,
			TokenURL: *tokenURL,
		}, []string{"all"}
	}
	panic("must provide endpoint for Oauth by wellknownEndpoint or authURL & tokenURL ")

}
