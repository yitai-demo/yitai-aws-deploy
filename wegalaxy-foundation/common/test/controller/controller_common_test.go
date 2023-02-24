package controller

import (
	"testing"

	"github.com/degalaxy/wegalaxy-service/wegalaxy-foundation/common/controller"
	"github.com/stretchr/testify/assert"
)

func TestTrimAndFirstToLowerCase(t *testing.T) {
	assert.Equal(t, "", controller.TrimAndFirstToLowerCase(""))
	assert.Equal(t, "", controller.TrimAndFirstToLowerCase("   "))
	assert.Equal(t, "a", controller.TrimAndFirstToLowerCase("   a"))
	assert.Equal(t, "a", controller.TrimAndFirstToLowerCase("A   "))
	assert.Equal(t, "abcDef", controller.TrimAndFirstToLowerCase(" abcDef   "))
	assert.Equal(t, "aBCDef", controller.TrimAndFirstToLowerCase(" ABCDef   "))
}
