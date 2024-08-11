package assert

import (
	"reflect"
	"strings"
	"testing"
)

func True(t *testing.T, cond bool, messages ...string) {
	t.Helper()
	if !cond {
		panic(msg(messages...))
	}
}

func False(t *testing.T, cond bool, messages ...string) {
	t.Helper()
	if cond {
		panic(msg(messages...))
	}
}

func Panic(t *testing.T, fn func(), messages ...string) {
	defer func() {
		x := recover()
		if x == nil {
			panic(msg(messages...))
		}
	}()
	fn()
}

func Nil(t *testing.T, x any, messages ...string) {
	t.Helper()
	if !isNil(x) {
		panic(msg(messages...))
	}
}

func NotNil(t *testing.T, x any, messages ...string) {
	t.Helper()
	if isNil(x) {
		panic(msg(messages...))
	}
}

func isNil(x any) bool {
	if x == nil {
		return true
	}
	rv := reflect.ValueOf(x)
	switch rv.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Pointer,
		reflect.Slice:
		return rv.IsNil()
	}
	return x == nil
}

func Eq(t *testing.T, left, right any, messages ...string) {
	t.Helper()
	if left != right {
		panic(msg(messages...))
	}
}

func DeepEqual(t *testing.T, left, right any, messages ...string) {
	t.Helper()
	if !reflect.DeepEqual(left, right) {
		panic(msg(messages...))
	}
}

func Ne(t *testing.T, left, right any, messages ...string) {
	t.Helper()
	if left == right {
		panic(msg(messages...))
	}
}

func msg(arg ...string) string {
	if len(arg) > 0 {
		return strings.Join(arg, ". ")
	}
	return "assert failed."
}
