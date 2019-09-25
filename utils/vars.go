package utils

import (
	"fmt"
	"github.com/beegons/config"
)

func GetOrionURL() string {
	config := config.New()
	return fmt.Sprintf("http://%s:%d/v2", config.Fiware.OrionCB_Host, config.Fiware.OrionCB_Port)
}

func GetCygnusURL() string {
	config := config.New()
	return fmt.Sprintf("http://%s:%d", config.Fiware.Cygnus_Host, config.Fiware.Cygnus_Port)
}

func GetDBURL() string {
	config := config.New()
	return fmt.Sprintf("mongodb://%s:%d/", config.Database.Host, config.Database.Port)
}

func GetAppDBName() string {
	config := config.New()
	return config.Database.Database
}

func GetCygnusDBName() string {
	config := config.New()
	return config.Fiware.Cygnus_Database
}

func GetFlinkURL() string {
	config := config.New()
	return fmt.Sprintf("http://%s:%d", config.Flink.Host, config.Flink.Port)
}

func GetAlertURL() string {
	config := config.New()
	return fmt.Sprintf("http://%s:%d", config.Server.Hostname, config.Server.Port)
}
