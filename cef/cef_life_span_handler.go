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
	"unsafe"
)

type LifeSpanHandler struct {
	browsers map[unsafe.Pointer]chan *Browser
}

func (l *LifeSpanHandler) RegisterAndWaitForBrowser(hwnd unsafe.Pointer) (browser *Browser) {
	l.browsers[hwnd] = make(chan *Browser)
	return <-l.browsers[hwnd]
}

func (l *LifeSpanHandler) OnAfterCreated(browser *Browser) {
	hwnd := unsafe.Pointer(browser.GetWindowHandle())
	fmt.Printf("created browser, handled by lifespan %v %v", browser, hwnd)
	b, ok := l.browsers[hwnd]
	if ok == true {
		b <- browser
	} else {
		fmt.Printf("Failed looking up browser at hwnd %v", hwnd)
	}
}

var _LifeSpanHandler *C.struct__cef_life_span_handler_t // requires reference counting
var globalLifespanHandler *LifeSpanHandler
