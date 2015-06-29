// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package main

import "C"
import (
	"fmt"
	"github.com/mmatey/cef2go/cef"
	"github.com/op/go-logging"
	"os"
	//"time"
)

func main() {
	//quit := make(chan int)
	cwd, _ := os.Getwd()
	logging.SetLevel(logging.DEBUG, "cef")
	var releasePath = os.Getenv("RELEASE_PATH")
	if releasePath == "" {
		releasePath = cwd
	}
	// you need to register to the callback before we fork processes
	cef.RegisterV8Callback("loaded", cef.V8Callback(func(args []*cef.V8Value) {
		arg0 := args[0].ToString()
		arg1 := args[1].ToString()
		fmt.Printf("Calling V8Callback Loaded args:  %s %s\n", arg0, arg1)
	}))

	cef.RegisterV8Callback("sup", cef.V8Callback(func(args []*cef.V8Value) {
		arg0 := args[0].ToInt32()
		arg1 := args[1].ToFloat32()
		arg2 := args[2].ToBool()
		arg3 := args[3].ToString()
		fmt.Printf("Calling V8Callback args: %d %f %v %s\n", arg0, arg1, arg2, arg3)
	}))
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
	//time.Sleep(2500 * time.Millisecond)
	// Create browser.
	browserSettings := &cef.BrowserSettings{}
	url := "http://www.retailmenot.com"
	go func() {
		browser := cef.CreateBrowser(browserSettings, "about:blank", false)
		//lifeSpan := cef.LifeSpanHandler{browser}
		//browser2, err := cef.LifeSpanHandler.RegisterAndWaitForBrowser()
		/*		if err != nil {
				fmt.Errorf(err.Error())
			}*/
		// cef.WaitForContextInitialized()
		loadPage(url, browser)
		browser.ExecuteJavaScript("console.log('loaded'); cef2go.callback('sup', window.screen.availWidth, 13.5, true, window.document.location.href);", "sup.js", 1)

		/*		browser.LoadURL("http://retailmenot.com")
				// cef.WaitForContextInitialized()
				browser.ExecuteJavaScript("console.log('loaded'); cef2go.callback('sup', window.screen.availWidth, 13.5, true, window.document.location.href);", "sup.js", 1)*/
	}()
	// CEF loop and shutdown.
	cef.RunMessageLoop()
	cef.Shutdown()
	os.Exit(1)
}

func loadPage(url string, browser *cef.Browser) {
	pageId := "abc"
	browser.LoadURL(url)
	browser.ExecuteJavaScript(fmt.Sprintf("console.log(document.body.length);console.log();if(document.body.length && document.body.length > 0){console.log('TEEEEST1b');cef2go.callback('loaded', '%s', 'a ' + document.body.length)}else{console.log('TEEEST1a');document.addEventListener('DOMContentLoaded', function(){console.log('TEEEEST2');cef2go.callback('loaded', '%s', 'a ' + window.document.title)}, true)};", pageId, pageId), "sup2.js", 1)
}
