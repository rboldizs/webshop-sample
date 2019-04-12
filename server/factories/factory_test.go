package factories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactory(t *testing.T) {

	fact := GetFactory()

	assert.NotNil(t, fact)
	assert.NotNil(t, fact.GetConfigurationObj())

}
