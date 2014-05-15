// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi

package cef

/*
#include <stdlib.h>
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"
extern void initialize_life_span_handler(struct _cef_life_span_handler_t* lifeHandler);
*/
import "C"

import (
	"fmt"
)

type LifeSpanHandler interface {
	OnAfterCreated(browser *Browser)
}

type DefaultLifeSpanHandler struct{}

func (l *DefaultLifeSpanHandler) OnAfterCreated(browser *Browser) {
	fmt.Printf("created browser, handled by lifespan %v", browser)
}

var _LifeSpanHandler *C.struct__cef_life_span_handler_t // requires reference counting
var globalLifespanHandler LifeSpanHandler

//export go_OnAfterCreated
func go_OnAfterCreated(self *C.struct__cef_life_span_handler_t, browser *C.cef_browser_t) {
	if globalLifespanHandler != nil {
		globalLifespanHandler.OnAfterCreated(&Browser{browser})
	}
}
func InitializeLifeSpanHandler() *C.struct__cef_life_span_handler_t {
	var handler *C.struct__cef_life_span_handler_t
	handler = (*C.struct__cef_life_span_handler_t)(C.calloc(1, C.sizeof_struct__cef_life_span_handler_t))
	C.initialize_life_span_handler(handler)
	return handler
}

func SetLifespanHandler(handler LifeSpanHandler) {
	globalLifespanHandler = handler
}
