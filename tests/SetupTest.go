package tests

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	pmio "pmo-sdk-go/pmio"
)

func (suite *ClientTestSuite) SetupTest() {
	cfg := OauthConfig{}
	configFile, err := os.Open("../env.json")
	defer configFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(configFile).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	api := pmio.NewClientApi()
	api.Configuration.BasePath = "https://" + cfg.Host + "/api/v1"
	api.Configuration.Host = cfg.Host
	api.Configuration.APIKey["Authorization"] = cfg.Key
	api.Configuration.APIKeyPrefix["Authorization"] = "Bearer"
	api.Configuration.AccessToken = cfg.Key
	suite.client = api
	rand.Seed(time.Now().UnixNano())
}
