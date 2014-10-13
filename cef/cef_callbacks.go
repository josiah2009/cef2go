package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"
*/
import "C"

import (
	"unsafe"
)

//export go_OnAfterCreated
func go_OnAfterCreated(self *C.struct__cef_life_span_handler_t, browserId int, browser *C.cef_browser_t) {
	if globalLifespanHandler != nil {
		globalLifespanHandler.OnAfterCreated(&Browser{browserId, browser, nil})
	}
}

//export go_Log
func go_Log(str *C.char) {
	log.Debug(C.GoString(str))
}

//export go_LogPointer
func go_LogPointer(str *C.char, p unsafe.Pointer) {
	log.Debug(C.GoString(str)+" %p", p)
}

//export go_OnConsoleMessage
func go_OnConsoleMessage(browser *C.cef_browser_t, message *C.cef_string_t, source *C.cef_string_t, line int) {
	consoleHandler(CEFToGoString(message), CEFToGoString(source), line)
}
