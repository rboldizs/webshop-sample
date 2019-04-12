package factories

import (
	"sync"

	endpoint_api "git.appagile.io/consumption/consumption-go-common/api"
	"git.appagile.io/services-paas/bb/vSweets/shop-backend-rest/rest/endpoints"
	"git.appagile.io/services-paas/bb/vSweets/shop-backend-rest/server/config"
)

var fact Factory
var factOnce sync.Once

//Factory interface is used to expose and create basic objects
type Factory interface {
	GetConfigurationObj() config.Configuration
	GetStatusEPObj() endpoint_api.Endpoint
}

type factory struct {
}

//GetFactory creates a singleton BeanFactory
func GetFactory() Factory {
	factOnce.Do(func() {
		fact = NewFactory()
	})

	return fact
}

// NewFactory creates a new BeanFactory
func NewFactory() Factory {
	return &factory{}

}

var configObj config.Configuration
var confOnce sync.Once

//GetConfigurationObj returns a singleton config object
func (ft *factory) GetConfigurationObj() config.Configuration {
	confOnce.Do(func() {
		configObj = config.GetConfig()
	})
	return configObj
}

var statusEPObj endpoint_api.Endpoint
var statusOnce sync.Once

//GetStatusEPObj returns a singleton status object
func (ft *factory) GetStatusEPObj() endpoint_api.Endpoint {
	statusOnce.Do(func() {
		statusEPObj = &endpoints.StatusEP{}
	})
	return statusEPObj
}
