package plugify

/*
#cgo LDFLAGS: -L${SRCDIR}/libplugify.a
#include <plugify.h>
*/
import "C"
import "unsafe"

type PluginStartCallback func()
type PluginEndCallback func()

type Plugify struct {
	fnPluginStartCallback PluginStartCallback
	fnPluginEndCallback   PluginEndCallback
}

var plugify Plugify = Plugify {
	fnPluginStartCallback: func() {},
	fnPluginEndCallback:   func() {},
}

const kApiVersion = 1

//export Plugify_Init
func Plugify_Init(api []uintptr, version int32) int32 {
	if version < kApiVersion {
		return kApiVersion
	}
	C.Plugify_SetMethodPtr(unsafe.Pointer(api[0]));
	return 0
}

//export Plugify_PluginStart
func Plugify_PluginStart() {
	plugify.fnPluginStartCallback()
}

func OnPluginStart(fn PluginStartCallback) {
	plugify.fnPluginStartCallback = fn
}

//export Plugify_PluginEnd
func Plugify_PluginEnd() {
	plugify.fnPluginEndCallback()
}

func OnPluginEnd(fn PluginEndCallback) {
	plugify.fnPluginEndCallback = fn
}

type Vector2 struct {
    x float32
    y float32
}
type Vector3 struct {
    x float32
    y float32
    z float32
}
type Vector4 struct {
    x float32
    y float32
    z float32
    w float32
}