package stringsutils_test

import (
	"testing"

	"github.com/mirzakhany/pkg/stringsutils"
	"github.com/stretchr/testify/assert"
)

func TestReplaceMsg(t *testing.T) {

	msg := "Hi NAME , its from SENDER, you won a CODE gift"
	placeHolders := []string{"NAME", "SENDER", "CODE"}
	values := []string{"dear", "foo", "1234"}

	r := stringsutils.ReplaceMsg(msg, placeHolders, values)
	assert.Equal(t, r, "Hi dear , its from foo, you won a 1234 gift")
}
