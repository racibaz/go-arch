package helper

import (
	"reflect"
	"runtime"
	"strings"
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

	/*
		it takes the string
		"github.com/racibaz/go-arch/internal/modules/user/features/signup/v1/adapters/endpoints/http.RegisterUserHandler.Store"
		and gives the func name "Store"
	*/

	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	funcName := parts[len(parts)-1]

	return funcName
}
