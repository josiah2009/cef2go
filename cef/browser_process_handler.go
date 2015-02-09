package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"
*/
import "C"

//export go_BrowserProcessHandlerOnContextInitialized
func go_BrowserProcessHandlerOnContextInitialized() {
	contextInitialized <- 1
	log.Debug("go_BrowserProcessHandlerOnContextInitialized")
}
