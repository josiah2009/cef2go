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
	"unsafe"
)

type LifeSpanHandler struct {
	browsers map[unsafe.Pointer]chan *Browser
        windowInfos []*C.cef_window_info_t
}

func (l *LifeSpanHandler) RegisterAndWaitForBrowser(hwnd unsafe.Pointer, windowInfo *C.cef_window_info_t) (browser *Browser) {
	l.browsers[hwnd] = make(chan *Browser)
	return <-l.browsers[hwnd]
}

func (l *LifeSpanHandler) OnAfterCreated(browser *Browser) {
	widgetHandle := browser.GetWindowHandle()
	Logger.Printf("created browser, handled by lifespan %v, %d\n", browser, widgetHandle)
        var hwnd unsafe.Pointer
        for i, w := range l.windowInfos {
          Logger.Printf("Window Info %d: widget %d parent %d", i, w.widget, w.parent_widget)
          if w.widget == widgetHandle {
              hwnd = unsafe.Pointer(w.parent_widget)
	      Logger.Printf("Found for widgetHandle %v %v\n", widgetHandle, hwnd)
              break
          }
        }
        if hwnd == nil {
		Logger.Printf("Failed finding parent for hwnd %v\n", widgetHandle)
        }
	b, ok := l.browsers[hwnd]
	if ok == true {
		b <- browser
	} else {
		Logger.Printf("Failed looking up browser at hwnd %v\n", hwnd)
	}
}

var _LifeSpanHandler *C.struct__cef_life_span_handler_t // requires reference counting
var globalLifespanHandler *LifeSpanHandler
