package tests

import (
	"testing"

	"github.com/echenim/globant/midleware"
	"github.com/stretchr/testify/assert"
)

var (
	miwaz midleware.IApi
)

func TestGetByCityAndCountry(t *testing.T) {
	miwaz = midleware.NewForeCastAPI()
	rs := miwaz.GetByCityAndCountry("Lagos", "ng", "1508a9a4840a5574c822d70ca2132032")
	assert.NotNil(t, rs, nil)
}

func TestGetByCityAndCountryAndDay(t *testing.T) {
	miwaz = midleware.NewForeCastAPI()
	rs := miwaz.GetByCityAndCountryAndDay("Lagos", "ng", "1508a9a4840a5574c822d70ca2132032")
	assert.NotNil(t, rs, nil)
}
