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

type LifeSpanHandler struct {
	browser chan *Browser
}

func (l *LifeSpanHandler) RegisterAndWaitForBrowser(url string) (browser *Browser) {
	return <-l.browser
}

func (l *LifeSpanHandler) OnAfterCreated(browser *Browser) {
	url := browser.GetURL()
	Logger.Printf("created browser, handled by lifespan %v, url %s\n", browser, url)
	l.browser <- browser
}

var _LifeSpanHandler *C.struct__cef_life_span_handler_t // requires reference counting
var globalLifespanHandler *LifeSpanHandler
