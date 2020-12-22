package tests

import (
	"testing"

	"github.com/echenim/globant/utils"

	"github.com/stretchr/testify/assert"
)

var (
	config utils.IConfigurationManager
)

func TestUnixToHumanReasible(t *testing.T) {
	config = utils.NewManager()
	rs := config.UnixToHumanReasible(1608681600)
	assert.NotEmpty(t, rs, "")
}

func TestTimeParse(t *testing.T) {
	config = utils.NewManager()
	rs := config.TimeParse(1608681600)
	assert.NotEmpty(t, rs, "")
}
func TestWindScale(t *testing.T) {
	config = utils.NewManager()
	rs := config.WindScale(6.0)
	assert.Contains(t, rs, "Light Breeze")
}

func TestDirection(t *testing.T) {
	config = utils.NewManager()
	rs := config.Direction(90)
	assert.Contains(t, rs, "east")
}

func TestCloudiness(t *testing.T) {
	config = utils.NewManager()
	rs := config.Cloudiness(12)
	assert.Contains(t, rs, "Mostly Clear")
}

// func TestConfigurationManager(t *testing.T) {
// 	config = utils.NewManager()
// 	rs := config.ConfigurationManager()
// 	assert.Equal(t, rs.ConfigurationManagerList[0].Port, "8080")
// }
