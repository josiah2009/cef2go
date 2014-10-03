package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"
*/
import "C"

import "unsafe"

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

//export go_OnConsoleMessage
func go_OnConsoleMessage(browser *C.cef_browser_t, message *C.cef_string_t, source *C.cef_string_t, line int) {
	consoleHandler(CEFToGoString(message), CEFToGoString(source), line)
}

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
          return callback(); 
        } 
      })();
    `
    C.cef_register_extension(CEFString("v8/cef2go"), CEFString(extCode), handler);
}

//export go_V8HandlerExecute
func go_V8HandlerExecute(name *C.cef_string_t, object *C.cef_v8value_t, argsCount C.size_t, args unsafe.Pointer, retval **C.cef_v8value_t, exception *C.cef_string_t) int {
	log.Debug("JS EXECUTION: %s", CEFToGoString(name))
	return 0
}

