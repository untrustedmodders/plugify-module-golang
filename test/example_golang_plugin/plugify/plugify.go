package plugify

import "C"

type PluginStartCallback func()
type PluginEndCallback func()

type Plugify struct {
	fnPluginStartCallback PluginStartCallback
	fnPluginEndCallback   PluginEndCallback
}

var plugify Plugify = Plugify{
	fnPluginStartCallback: func() {},
	fnPluginEndCallback:   func() {},
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
