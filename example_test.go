package assert_test

import (
	"fmt"
	"testing"

	"code.gopub.tech/assert"
)

type myErr struct {
	msg string
}

func (e *myErr) Error() string { return e.msg }

func nilErr() *myErr { return nil }

func TestAssert(t *testing.T) {
	assert.True(t, 1 == 1) //lint:ignore SA4000 always true
	assert.False(t, false)
	assert.Equal(t, 1, 1)
	assert.NotEqual(t, 1, 0)
	assert.DeepEqual(t, 1, 1)

	var err error
	assert.Nil(t, err)
	err = fmt.Errorf("err")
	assert.NotNil(t, err)

	err = nilErr()
	assert.True(t, err != nil) //lint:ignore SA4023 always true
	assert.Nil(t, err)

	assert.Panic(t, func() {
		_ = err.Error()
	}, "err is nil")

	t.Run("assert fail", func(t *testing.T) {
		t.Skip()
		assert.True(t, false, "assert fail")
		assert.False(t, true, "assert fail")
		assert.Equal(t, 1, 2)
		assert.NotEqual(t, 1, 1)
		assert.DeepEqual(t, 1, 2)
		assert.Nil(t, 1)
		assert.NotNil(t, nil)
		assert.Panic(t, func() {})
	})
}
