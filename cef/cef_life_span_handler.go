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
	browsers map[string]chan *Browser
}

func (l *LifeSpanHandler) RegisterAndWaitForBrowser(url string) (browser *Browser) {
	l.browsers[url] = make(chan *Browser)
	return <-l.browsers[url]
}

func (l *LifeSpanHandler) OnAfterCreated(browser *Browser) {
	url := browser.GetURL()
	Logger.Printf("created browser, handled by lifespan %v, url %s\n", browser, url)
	b, ok := l.browsers[url]
	if ok == true {
		b <- browser
	} else {
		Logger.Printf("Failed looking up browser at url %s\n", url)
	}
}

var _LifeSpanHandler *C.struct__cef_life_span_handler_t // requires reference counting
var globalLifespanHandler *LifeSpanHandler
