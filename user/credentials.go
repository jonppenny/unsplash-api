package user

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetUserCredentials(config *string) (string, error) {
	viper.SetConfigFile(*config)
	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	clientId := fmt.Sprintf("%s", viper.Get("client_id"))

	return clientId, nil
}
