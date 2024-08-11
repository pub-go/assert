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
	assert.Eq(t, 1, 1)
	assert.Ne(t, 1, 0)
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
	assert.Panic(t, func() {
		assert.True(t, false, "will panic")
	})
	assert.Panic(t, func() {
		assert.False(t, true)
	})
	assert.Panic(t, func() {
		assert.Panic(t, func() {})
	})
	assert.Panic(t, func() {
		assert.Nil(t, 1)
	})
	assert.Panic(t, func() {
		assert.NotNil(t, err)
	})
	assert.Panic(t, func() {
		assert.Eq(t, 1, 2)
	})
	assert.Panic(t, func() {
		assert.DeepEqual(t, 1, 2)
	})
	assert.Panic(t, func() {
		assert.Ne(t, 1, 1)
	})
}
