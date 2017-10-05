package tests

import (
	"math/rand"
	"strconv"

	pmio "github.com/ProcessMaker/pmio-sdk-go"
)

//Test case for AddProcess
func (suite *ClientTestSuite) TestAddProcess() {
	rnd := strconv.Itoa(rand.Int())

	processAtt := pmio.ProcessAttributes{}
	processAtt.Name = "ProcessName" + rnd
	processAtt.Status = "ACTIVE"
	processAtt.Type_ = "NORMAL"

	process := pmio.Process{}
	process.Attributes = processAtt

	processCreateItem := pmio.ProcessCreateItem{Data: process}

	procItem, _, err := suite.client.AddProcess(processCreateItem)
	suite.Require().NoError(err)
	suite.NotEmpty(procItem.Data.Id)
	suite.Equal(processAtt.Name, procItem.Data.Attributes.Name)
}
