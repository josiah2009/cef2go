// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi

package cef

/*
#include <stdlib.h>
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"
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
	fmt.Printf("created browser, handled by lifespan %v %v", browser, browser.GetWindowHandle())
}

var _LifeSpanHandler *C.struct__cef_life_span_handler_t // requires reference counting
var globalLifespanHandler LifeSpanHandler

func SetLifespanHandler(handler LifeSpanHandler) {
	globalLifespanHandler = handler
}
