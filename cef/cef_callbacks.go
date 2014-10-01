package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"
*/
import "C"

import (
	"fmt"
	"github.com/errnoh/utils/bgra"
	"image"
	"image/png"
	"os"
	"unsafe"
)

//export go_OnAfterCreated
func go_OnAfterCreated(self *C.struct__cef_life_span_handler_t, browser *C.cef_browser_t) {
	if globalLifespanHandler != nil {
		globalLifespanHandler.OnAfterCreated(&Browser{browser})
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

//export go_RenderHandlerOnPaint
func go_RenderHandlerOnPaint(browserId int, buffer unsafe.Pointer, width, height int) {
	gbuffer := C.GoBytes(buffer, C.int(width*height*4))
	log.Debug("on paint %d %d %dx%d", browserId, len(gbuffer), width, height)
	bgra_image := &bgra.BGRA{Pix: gbuffer, Stride: 32, Rect: image.Rect(0, 0, 700, 700)}
	filename := fmt.Sprintf("./buffer_%d.png", browserId)
	f, err := os.Create(filename)
	if err != nil {
		log.Error("Could not open file for writing %s", err)
		return
	}
	err = png.Encode(f, bgra_image)
	if err != nil {
		log.Error("Could not encode png %s", err)
		return
	}
	log.Debug("Wrote buffer to %s", filename)

}
