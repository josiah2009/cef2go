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

type RenderHandler interface {
	GetRootScreenRect(*C.cef_rect_t) int
	GetViewRect(*C.cef_rect_t) int
	GetScreenPoint(int, int, *int, *int) int
	GetScreenInfo(*C.cef_screen_info_t) int
	OnPopupShow(int)
	OnPopupSize(*C.cef_rect_t)
	OnPaint(C.cef_paint_element_type_t, C.size_t, unsafe.Pointer, unsafe.Pointer, int, int)
	OnCursorChange(C.cef_cursor_handle_t)
	OnScrollOffsetChanged()
}

//export go_RenderHandlerGetRootScreenRect
func go_RenderHandlerGetRootScreenRect(browserId int, rect *C.cef_rect_t) int {
	if b, ok := BrowserById(browserId); ok {
		return b.RenderHandler.GetRootScreenRect(rect)
	}
	return 0
}

//export go_RenderHandlerGetViewRect
func go_RenderHandlerGetViewRect(browserId int, rect *C.cef_rect_t) int {
	if b, ok := BrowserById(browserId); ok {
		return b.RenderHandler.GetViewRect(rect)
	}
	return 0
}

//export go_RenderHandlerGetScreenPoint
func go_RenderHandlerGetScreenPoint(browserId, x, y int, screenX *int, screenY *int) int {
	if b, ok := BrowserById(browserId); ok {
		return b.RenderHandler.GetScreenPoint(x, y, screenX, screenY)
	}
	return 0
}

//export go_RenderHandlerGetScreenInfo
func go_RenderHandlerGetScreenInfo(browserId int, info *C.cef_screen_info_t) int {
	if b, ok := BrowserById(browserId); ok {
		return b.RenderHandler.GetScreenInfo(info)
	}
	return 0
}

//export go_RenderHandlerOnPopupShow
func go_RenderHandlerOnPopupShow(browserId int, show int) {
	if b, ok := BrowserById(browserId); ok {
		b.RenderHandler.OnPopupShow(show)
	}
}

//export go_RenderHandlerOnPopupSize
func go_RenderHandlerOnPopupSize(browserId int, size *C.cef_rect_t) {
	if b, ok := BrowserById(browserId); ok {
		b.RenderHandler.OnPopupSize(size)
	}
}

//export go_RenderHandlerOnPaint
func go_RenderHandlerOnPaint(browserId int, paintType C.cef_paint_element_type_t, dirtyRectsCount C.size_t, dirtyRects unsafe.Pointer, buffer unsafe.Pointer, width, height int) {
	if b, ok := BrowserById(browserId); ok {
		b.RenderHandler.OnPaint(paintType, dirtyRectsCount, dirtyRects, buffer, width, height)
	}
	// gbuffer := C.GoBytes(buffer, C.int(width*height*4))
	// log.Debug("on paint %d %d %dx%d", browserId, len(gbuffer), width, height)
	// bgra_image := &bgra.BGRA{Pix: gbuffer, Stride: 4 * width, Rect: image.Rect(0, 0, 700, 700)}
	// filename := fmt.Sprintf("./buffer_%d.png", browserId)
	// f, err := os.Create(filename)
	// if err != nil {
	// 	log.Error("Could not open file for writing %s", err)
	// 	return
	// }
	// err = png.Encode(f, bgra_image)
	// if err != nil {
	// 	log.Error("Could not encode png %s", err)
	// 	return
	// }
	// log.Debug("Wrote buffer to %s", filename)

}

//export go_RenderHandlerOnCursorChange
func go_RenderHandlerOnCursorChange(browserId int, cursor C.cef_cursor_handle_t) {
	if b, ok := BrowserById(browserId); ok {
		b.RenderHandler.OnCursorChange(cursor)
	}
}

//export go_RenderHandlerOnScrollOffsetChanged
func go_RenderHandlerOnScrollOffsetChanged(browserId int) {
	if b, ok := BrowserById(browserId); ok {
		b.RenderHandler.OnScrollOffsetChanged()
	}
}

type DefaultRenderHandler struct {
	Browser *Browser
}

func (d *DefaultRenderHandler) GetRootScreenRect(rect *C.cef_rect_t) int {
	return 0
}

func (d *DefaultRenderHandler) GetViewRect(rect *C.cef_rect_t) int {
	return 0
}

func (d *DefaultRenderHandler) GetScreenPoint(x, y int, screenX, screenY *int) int {
	return 0
}

func (d *DefaultRenderHandler) GetScreenInfo(info *C.cef_screen_info_t) int {
	return 0
}

func (d *DefaultRenderHandler) OnPopupShow(show int) {
}

func (d *DefaultRenderHandler) OnPopupSize(size *C.cef_rect_t) {
}

func (d *DefaultRenderHandler) OnPaint(paintType C.cef_paint_element_type_t, dirtyRectsCount C.size_t, dirtyRects unsafe.Pointer, buffer unsafe.Pointer, width, height int) {
}

func (d *DefaultRenderHandler) OnCursorChange(cursor C.cef_cursor_handle_t) {
}

func (d *DefaultRenderHandler) OnScrollOffsetChanged() {
}
