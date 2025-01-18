package auth

import (
	"context"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var (
	v               = viper.New()
	providerConfigs = map[string]*IDP{}
)

func init() {
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_", "/", "_"))
	v.AutomaticEnv()
	v.AddConfigPath(func() (ret string) {
		if ret = os.Getenv("APP_CONFIG_PATH"); ret == "" {
			ret = "/etc/ums/"
		}
		return
	}())
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Sub("idps").Unmarshal(&providerConfigs); err != nil {
		panic(err)
	}

	if err := updateIDPProviders(providerConfigs); err != nil {
		panic(err)
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Configuration file changed")
	})
	v.WatchConfig()
}

type IDP struct {
	Endpoint              string   `mapstructure:"endpoint"`
	ClientID              string   `mapstructure:"client_id"`
	ClientSecret          string   `mapstructure:"client_secret"`
	RedirectURI           string   `mapstructure:"redirect_uri"`
	PostLogoutRedirectURI string   `mapstructure:"post_logout_redirect_uri"`
	Scope                 []string `mapstructure:"scopes"`

	provider        *oidc.Provider        `mapstructure:"-"`
	oauth2Conf      *oauth2.Config        `mapstructure:"-"`
	oidcConf        *oidc.Config          `mapstructure:"-"`
	idTokenVerifier *oidc.IDTokenVerifier `mapstructure:"-"`
}

func updateIDPProviders(idps map[string]*IDP) error {
	for key, idpConfig := range idps {
		provider, err := oidc.NewProvider(context.Background(), idpConfig.Endpoint)
		if err != nil {
			return err
		}
		idps[key].provider = provider
		idps[key].oidcConf = &oidc.Config{
			ClientID: idpConfig.ClientID,
		}
		idpConfig.idTokenVerifier = provider.Verifier(idps[key].oidcConf)
		idps[key].oauth2Conf = &oauth2.Config{
			ClientID:     idpConfig.ClientID,
			ClientSecret: idpConfig.ClientSecret,
			Endpoint:     provider.Endpoint(),
			RedirectURL:  idpConfig.RedirectURI,
			Scopes: func(additionalScopes []string) []string {
				defaultScopes := []string{oidc.ScopeOpenID, "email", "profile"}
				combinedScopes := defaultScopes
				for _, scope := range additionalScopes {
					if slices.Index[[]string](combinedScopes, scope) >= 0 {
						combinedScopes = append(combinedScopes, scope)
					}
				}

				return combinedScopes
			}(idpConfig.Scope),
		}
	}
	return nil
}

func getIDP(idp string) *IDP {
	return providerConfigs[idp]
}
