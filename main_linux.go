// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package main

import "C"
import (
	"fmt"
	"github.com/op/go-logging"
	"github.com/paperlesspost/cef2go/cef"
	"os"
)

func main() {
	cwd, _ := os.Getwd()
	logging.SetLevel(logging.DEBUG, "cef")
	var releasePath = os.Getenv("RELEASE_PATH")
	if releasePath == "" {
		releasePath = cwd
	}
	// you need to register to the callback before we fork processes
	cef.RegisterV8Callback("sup", cef.V8Callback(func(args []cef.V8Value) {
		arg0 := cef.V8ValueToInt32(args[0])
		arg1 := cef.V8ValueToInt32(args[1])
		arg2 := cef.V8ValueToBool(args[2])
		arg3 := cef.V8ValueToString(args[3])
		fmt.Printf("Calling V8Callback args: %d %d %v %s\n", arg0, arg1, arg2, arg3)
	}))
	// CEF subprocesses.
	cef.ExecuteProcess(nil)

	// CEF initialize.
	settings := cef.Settings{}
	settings.ResourcesDirPath = releasePath
	settings.LocalesDirPath = releasePath + "/locales"
	settings.CachePath = cwd + "/webcache"         // Set to empty to disable
	settings.LogSeverity = cef.LOGSEVERITY_VERBOSE // LOGSEVERITY_VERBOSE
	settings.LogFile = cwd + "/debug.log"
	settings.RemoteDebuggingPort = 7000
	cef.Initialize(settings)

	// Create GTK window.
	//gtk.Initialize()
	// window := gtk.CreateWindow("cef2go example", 1024, 768)
	// gtk.ConnectDestroySignal(window, OnDestroyWindow)
	//cef.XlibRegisterHandlers()

	// Create browser.
	browserSettings := &cef.BrowserSettings{}
	url := "file://" + cwd + "/Release/example.html"
	go func() {
		browser := cef.CreateBrowser(browserSettings, url, true)
		browser.ExecuteJavaScript("console.log('we outchea'); cef2go.callback('sup', 10, 10, true, 'something');", "sup.js", 1)
	}()
	// CEF loop and shutdown.
	cef.RunMessageLoop()
	cef.Shutdown()
	os.Exit(0)
}

func OnDestroyWindow() {
	cef.QuitMessageLoop()
}
