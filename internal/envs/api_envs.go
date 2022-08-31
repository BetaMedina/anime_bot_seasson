package envs

import (
	"log"
	"os"
)

type ApiEnvs struct {
	API_URL string
}

func LoadApiEnvs() *ApiEnvs {
	API_URL, hasApiUrl := os.LookupEnv("API_URL")
	if !hasApiUrl {
		log.Fatal("Environment Variable API_URL must be defined")
	}
	config := &ApiEnvs{API_URL: API_URL}
	return config
}
