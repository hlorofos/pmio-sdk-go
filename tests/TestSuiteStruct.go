package tests

import (
	"github.com/ProcessMaker/pmio-sdk-go"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	client *pmio.Client
}
