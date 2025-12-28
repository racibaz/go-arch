package helper

import (
	"reflect"
	"runtime"
)

// StructName returns the name of the struct type
func StructName(v any) string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

// CurrentFuncName returns the name of the calling function
func CurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
