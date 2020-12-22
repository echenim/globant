package midleware

import (
	"time"

	"github.com/echenim/globant/models"
	"github.com/patrickmn/go-cache"
)

var (
	c = cache.New(2*time.Minute, 10*time.Minute)
)

func addToCach(key string, data *models.SixDaysForecastSetResp) {
	c.Set(key, &data, cache.DefaultExpiration)
}

func get(key string) models.SixDaysForecastSetResp {
	data, found := c.Get(key)
	if !found {

	}

	rsObj := models.SixDaysForecastSetResp{}
	rsObj, ok := data.(models.SixDaysForecastSetResp)

	if ok {
		return rsObj
	}

	return rsObj
}
