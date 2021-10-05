package reflect

import (
	"reflect"
	"runtime"
	"strings"
)

func FNName(fn interface{}) string {
	if fn == nil {
		return ""
	}
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

func TrimFN(name string) string {
	i := strings.LastIndex(name, ".")
	return name[i+1:]
}
