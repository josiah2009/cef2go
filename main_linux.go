// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package main

//#include <string.h>
import "C"
import (
	"fmt"
	"github.com/josiah2009/cef2go/cef"
	"github.com/op/go-logging"
	"os"
	"time"
	"unsafe"
)

const browserWidth = 1280
const browserHeight = 720

type OffscreenRenderHandler struct {
}

func (this *OffscreenRenderHandler) GetRootScreenRect(rect *cef.CefRect) int {
	rect.SetDimensions(0, 0, browserWidth, browserHeight)
	return 1
}

func (this *OffscreenRenderHandler) GetViewRect(rect *cef.CefRect) int {
	rect.SetDimensions(0, 0, browserWidth, browserHeight)
	return 1
}

func (this *OffscreenRenderHandler) GetScreenPoint(x, y int, screenX, screenY *int) int {
	return 0
}

func (this *OffscreenRenderHandler) GetScreenInfo(info *cef.CefScreenInfo) int {
	return 0
}

func (this *OffscreenRenderHandler) OnPopupShow(show int) {
}

func (this *OffscreenRenderHandler) OnPopupSize(size *cef.CefRect) {
}

func (this *OffscreenRenderHandler) OnPaint(paintType cef.CefPaintElementType, dirtyRectsCount int, dirtyRects unsafe.Pointer, buffer unsafe.Pointer, width, height int) {
}

func (this *OffscreenRenderHandler) OnCursorChange(cursor cef.CefCursorHandle, ctype cef.CefCursorType, custom_cursor_info cef.CefCursorInfo) {
}

func (this *OffscreenRenderHandler) OnScrollOffsetChanged(x, y float64) {
}

func main() {
	cwd, _ := os.Getwd()
	logging.SetLevel(logging.DEBUG, "cef")
	var releasePath = os.Getenv("RELEASE_PATH")
	if releasePath == "" {
		releasePath = cwd
	}

	// CEF subprocesses.
	cef.ExecuteProcess(nil)

	// CEF initialize.
	settings := cef.Settings{}
	settings.SingleProcess = 0
	settings.ResourcesDirPath = releasePath
	settings.LocalesDirPath = releasePath + "/locales"
	settings.CachePath = cwd + "/webcache"      // Set to empty to disable
	settings.LogSeverity = cef.LOGSEVERITY_INFO // LOGSEVERITY_VERBOSE
	settings.LogFile = cwd + "/debug.log"
	settings.RemoteDebuggingPort = 7000
	init := cef.Initialize(settings)
	fmt.Printf("Initialized: %d", init)
	cef.XlibRegisterHandlers()
	time.Sleep(2500 * time.Millisecond)
	// Create browser.
	browserSettings := &cef.BrowserSettings{}
	url := "file://" + cwd + "/Release/example.html"
	go func() {
		browser := cef.CreateBrowser(browserSettings, url, true)
		browser.RenderHandler = &OffscreenRenderHandler{}
	}()
	// CEF loop and shutdown.
	cef.RunMessageLoop()
	cef.Shutdown()
	os.Exit(1)
}
