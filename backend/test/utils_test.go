package test

import (
	"testing"

	"github.com/iacopoghilardi/alarm-tool/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsEmailValid(t *testing.T) {
	assert.True(t, utils.IsEmailValid("test@test.com"))
	assert.False(t, utils.IsEmailValid("test"))
}
