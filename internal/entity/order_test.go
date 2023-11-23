package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIdIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is missing")
}
