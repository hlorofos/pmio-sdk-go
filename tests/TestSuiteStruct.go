package tests

import (
	pmio "pmo-sdk-go/pmio"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	client *pmio.ClientApi
}
