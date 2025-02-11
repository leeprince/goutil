package assert_test

import (
	"testing"

	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/testutil/assert"
)

func TestAssertions_Chain(t *testing.T) {
	// err := "error message"
	err := errorx.Raw("error message")

	as := assert.New(t).
		NotNil(err).
		Err(err).
		ErrMsg(err, "error message")

	assert.True(t, as.IsOk())
	assert.False(t, as.IsFail())
}
