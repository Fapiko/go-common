package sliceutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {

	assert := assert.New(t)

	testSlice := []string{
		"arnold",
		"face",
		"palmer",
	}

	assert.True(StringInSlice(testSlice, "arnold"))
	assert.True(StringInSlice(testSlice, "face"))
	assert.True(StringInSlice(testSlice, "palmer"))
	assert.False(StringInSlice(testSlice, "awesomesauce"))
	assert.False(StringInSlice(testSlice, "arnol"))

}
