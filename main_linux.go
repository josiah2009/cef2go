// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package main

import "C"
import (
	"github.com/op/go-logging"
	"github.com/paperlesspost/cef2go/cef"
	"github.com/paperlesspost/cef2go/gtk"
	"os"
)

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
	settings.ResourcesDirPath = releasePath
	settings.LocalesDirPath = releasePath + "/locales"
	settings.CachePath = cwd + "/webcache"      // Set to empty to disable
	settings.LogSeverity = cef.LOGSEVERITY_INFO // LOGSEVERITY_VERBOSE
	settings.LogFile = cwd + "/debug.log"
	settings.RemoteDebuggingPort = 7000
	cef.Initialize(settings)

	// Create GTK window.
	gtk.Initialize()
	window := gtk.CreateWindow("cef2go example", 1024, 768)
	gtk.ConnectDestroySignal(window, OnDestroyWindow)

	// Create browser.
	browserSettings := &cef.BrowserSettings{}
	url := "file://" + cwd + "/Release/example.html"
	go func() {
		browser := cef.CreateBrowser(window, browserSettings, url, true)
		browser.ExecuteJavaScript("console.log('we outchea');", "sup.js", 1)
	}()
	// CEF loop and shutdown.
	cef.RunMessageLoop()
	cef.Shutdown()
	os.Exit(0)
}

func OnDestroyWindow() {
	cef.QuitMessageLoop()
}
