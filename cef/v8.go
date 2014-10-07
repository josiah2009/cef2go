package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"
#include "cef_helpers.h"
*/
import "C"

import (
	"reflect"
	"unsafe"
)

type V8Value *C.cef_v8value_t
type V8Callback func([]V8Value)

var V8Callbacks map[string]V8Callback

//export go_RenderProcessHandlerOnWebKitInitialized
func go_RenderProcessHandlerOnWebKitInitialized(handler *C.cef_v8handler_t) {
	log.Debug("go_RenderProcessHandlerOnWebKitInitialized")
	extCode := `
      var cef2go;
      if (!cef2go) { 
        cef2go = {}; 
      } 
      (function() { 
        cef2go.callback = function() { 
          native function callback(); 
          return callback.apply(this, arguments); 
        } 
      })();
    `
	C.cef_register_extension(CEFString("v8/cef2go"), CEFString(extCode), handler)
}

//export go_V8HandlerExecute
func go_V8HandlerExecute(name *C.cef_string_t, object *C.cef_v8value_t, argsCount C.size_t, args **C.cef_v8value_t, retval **C.cef_v8value_t, exception *C.cef_string_t) int {
	argsN := int(argsCount)
	if argsN < 1 {
		return 0
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(args)),
		Len:  argsN,
		Cap:  argsN,
	}
	arguments := *(*[]V8Value)(unsafe.Pointer(&hdr))
	log.Debug("Args: %v", arguments)
	callbackNameValue, arguments := arguments[0], arguments[1:]
	callbackName := V8ValueToString(callbackNameValue)
	log.Debug("callbackName %s %v", callbackName, V8Callbacks)
	if cb, ok := V8Callbacks[callbackName]; ok {
		cb(arguments)
		return 1
	}
	return 0
}

func RegisterV8Callback(name string, callback V8Callback) {
	if V8Callbacks == nil {
		V8Callbacks = make(map[string]V8Callback)
	}
	V8Callbacks[name] = callback
}

func V8ValueToInt32(v V8Value) int32 {
	return int32(C.v8ValueToInt32((*C.cef_v8value_t)(v)))
}

func V8ValueToString(v V8Value) string {
	return CEFToGoString(C.v8ValueToString((*C.cef_v8value_t)(v)))
}

func V8ValueToBool(v V8Value) bool {
	return int(C.v8ValueToBool((*C.cef_v8value_t)(v))) == 1
}
