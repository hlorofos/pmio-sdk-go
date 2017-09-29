package tests

import (
	"testing"
	"log"
	"encoding/json"
	"os"
	pmio "pmo-sdk-go/pmio"
	"math/rand"
	"strconv"
	"time"
	"github.com/stretchr/testify/suite"
//	"github.com/stretchr/testify/assert"
)

type ClientTestSuite struct {
	suite.Suite
	client *pmio.ClientApi
}


func (suite *ClientTestSuite) SetupTest() {
	cfg :=  OauthConfig{}
	configFile, err := os.Open("../env.json")
	defer configFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	err = json.NewDecoder(configFile).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg.Host)

	api := pmio.NewClientApi()
	api.Configuration.BasePath = "https://" + cfg.Host + "/api/v1"
	api.Configuration.Host = cfg.Host 
	api.Configuration.APIKey["Authorization"] = cfg.Key
	api.Configuration.APIKeyPrefix["Authorization"] = "Bearer"
	api.Configuration.AccessToken = cfg.Key
	suite.client = api
}

func (suite *ClientTestSuite) TestAddUserClientApi() {
	//api := setUpUsersApi()
	rand.Seed(time.Now().UnixNano())
	rnd := strconv.Itoa(rand.Int())
	userAtt := pmio.UserAttributes{}
	userAtt.Username = "user"+rnd
	userAtt.Password = "pass+"+rnd
	userAtt.Firstname = "Name"+rnd
	userAtt.Lastname = "Last"+rnd
	userAtt.Email = userAtt.Firstname + "@" + userAtt.Lastname + ".com"

	user := pmio.User{}
	user.Attributes = userAtt
	user.Type_ = "user"

	userCreateItem := pmio.UserCreateItem{Data:user}

	userItem, _, err := suite.client.AddUser(userCreateItem)

	if suite.Nil(err) {
		suite.Equal(userAtt.Username, userItem.Data.Attributes.Username)
	}
}

func TestMPIOClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
