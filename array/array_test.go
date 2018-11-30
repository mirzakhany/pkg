package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringInArray(t *testing.T) {

	testArr := []string{"THIS", "IS", "a", "test", "string"}
	isIn := StringInArray("test", testArr...)
	assert.True(t, isIn)

	isIn = StringInArray("Not", testArr...)
	assert.False(t, isIn)
}
