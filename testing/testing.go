package testing

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Assert will log the given message if condition is false.
func Assert(condition bool, t testing.TB, msg string, v ...interface{}) {
	AssertUp(condition, t, 1, msg, v...)
}

// AssertUp is like assert, but used inside helper functions, to ensure that
// the file and line number reported by failures corresponds to one or more
// levels up the stack.
func AssertUp(condition bool, t testing.TB, caller int, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(caller + 1)
		v = append([]interface{}{filepath.Base(file), line}, v...)
		fmt.Printf("%s:%d: "+msg+"\n", v...)
		t.FailNow()
	}
}

// Equals tests that the two values are equal according to reflect.DeepEqual.
func Equals(exp, act interface{}, t testing.TB) {
	EqualsUp(exp, act, t, 1)
}

// EqualsUp is like equals, but used inside helper functions, to ensure that the
// file and line number reported by failures corresponds to one or more levels
// up the stack.
func EqualsUp(exp, act interface{}, t testing.TB, caller int) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(caller + 1)
		fmt.Printf("%s:%d: exp: %v (%T), got: %v (%T)\n",
			filepath.Base(file), line, exp, exp, act, act)
		t.FailNow()
	}
}

// IsNil reports a failure if the given value is not nil.  Note that values
// which cannot be nil will always fail this check.
func IsNil(obtained interface{}, t testing.TB) {
	IsNilUp(obtained, t, 1)
}

// IsNilUp is like isNil, but used inside helper functions, to ensure that the
// file and line number reported by failures corresponds to one or more levels
// up the stack.
func IsNilUp(obtained interface{}, t testing.TB, caller int) {
	if !_isNil(obtained) {
		_, file, line, _ := runtime.Caller(caller + 1)
		fmt.Printf("%s:%d: expected nil, got: %v\n", filepath.Base(file), line, obtained)
		t.FailNow()
	}
}

// NotNil reports a failure if the given value is nil.
func NotNil(obtained interface{}, t testing.TB) {
	NotNilUp(obtained, t, 1)
}

// NotNilUp is like notNil, but used inside helper functions, to ensure that the
// file and line number reported by failures corresponds to one or more levels
// up the stack.
func NotNilUp(obtained interface{}, t testing.TB, caller int) {
	if _isNil(obtained) {
		_, file, line, _ := runtime.Caller(caller + 1)
		fmt.Printf("%s:%d: expected non-nil, got: %v\n", filepath.Base(file), line, obtained)
		t.FailNow()
	}
}

// _isNil is a helper function for isNil and notNil, and should not be used
// directly.
func _isNil(obtained interface{}) bool {
	if obtained == nil {
		return true
	}

	switch v := reflect.ValueOf(obtained); v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}

	return false
}
