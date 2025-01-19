package assert

import (
	"fmt"
	"reflect"
	"testing"
)

func True(t *testing.T, cond bool, messages ...any) {
	if !cond {
		t.Helper()
		t.Error(toMsg(messages, "Expected true, got false"))
	}
}

func False(t *testing.T, cond bool, messages ...any) {
	if cond {
		t.Helper()
		t.Error(toMsg(messages, "Expected false, got true"))
	}
}

func Equal(t *testing.T, left, right any, messages ...any) {
	if left != right {
		t.Helper()
		t.Error(toMsg(messages, "Expected equal, left=%#v, right=%#v", left, right))
	}
}

func DeepEqual(t *testing.T, left, right any, messages ...any) {
	if !reflect.DeepEqual(left, right) {
		t.Helper()
		t.Error(toMsg(messages, "Expected deep equal, left=%#v, right=%#v", left, right))
	}
}

func NotEqual(t *testing.T, left, right any, messages ...any) {
	if left == right {
		t.Helper()
		t.Error(toMsg(messages, "Expected not equal, left=%#v, right=%#v", left, right))
	}
}

func Nil(t *testing.T, x any, messages ...any) {
	if !isNil(x) {
		t.Helper()
		t.Error(toMsg(messages, "Expected nil, got %#v", x))
	}
}

func NotNil(t *testing.T, x any, messages ...any) {
	if isNil(x) {
		t.Helper()
		t.Error(toMsg(messages, "Expected not nil, got %#v", x))
	}
}

func Panic(t *testing.T, fn func(), messages ...any) {
	defer func() {
		t.Helper()
		x := recover()
		if x == nil {
			t.Error(toMsg(messages, "Expected panic, but not"))
		}
	}()
	fn()
}

func toMsg(args []any, fallback ...any) string {
	if s := toStr(args); s != "" {
		return s
	}
	return toStr(fallback)
}

func toStr(args []any) string {
	if len(args) > 0 {
		if template, ok := args[0].(string); ok {
			return fmt.Sprintf(template, args[1:]...)
		}
	}
	return ""
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
