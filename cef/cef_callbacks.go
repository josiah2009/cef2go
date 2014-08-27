package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"
*/
import "C"

//export go_OnAfterCreated
func go_OnAfterCreated(self *C.struct__cef_life_span_handler_t, browser *C.cef_browser_t) {
	if globalLifespanHandler != nil {
		globalLifespanHandler.OnAfterCreated(&Browser{browser})
	}
}

//export go_Log
func go_Log(str *C.char) {
	// Logger.Print(C.GoString(str))
}
