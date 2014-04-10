// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package cef

/*
CEF capi fixes
--------------
1. In cef_string.h:
    this => typedef cef_string_utf16_t cef_string_t;
    to => #define cef_string_t cef_string_utf16_t
2. In cef_export.h:
    #elif defined(COMPILER_GCC)
    #define CEF_EXPORT __attribute__ ((visibility("default")))
    #ifdef OS_WIN
    #define CEF_CALLBACK __stdcall
    #else
    #define CEF_CALLBACK
    #endif
*/

/*
#cgo CFLAGS: -I./../../
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "handlers/cef_app.h"
#include "handlers/cef_client.h"
*/
import "C"
import "unsafe"
import (
	"log"
	"os"
	"runtime"
)

var Logger *log.Logger = log.New(os.Stdout, "[cef] ", log.Lshortfile)

var _MainArgs *C.struct__cef_main_args_t
var _AppHandler *C.cef_app_t               // requires reference counting
var _ClientHandler *C.struct__cef_client_t // requires reference counting

// Sandbox is disabled. Including the "cef_sandbox.lib"
// library results in lots of GCC warnings/errors. It is
// compatible only with VS 2010. It would be required to
// build it using GCC. Add -lcef_sandbox to LDFLAGS.
// capi doesn't expose sandbox functions, you need do add
// these before import "C":
// void* cef_sandbox_info_create();
// void cef_sandbox_info_destroy(void* sandbox_info);
var _SandboxInfo unsafe.Pointer

type Settings struct {
	CachePath           string
	LogSeverity         int
	LogFile             string
	ResourcesDirPath    string
	LocalesDirPath      string
	RemoteDebuggingPort int
}

type BrowserSettings struct {
}

const (
	LOGSEVERITY_DEFAULT      = C.LOGSEVERITY_DEFAULT
	LOGSEVERITY_VERBOSE      = C.LOGSEVERITY_VERBOSE
	LOGSEVERITY_INFO         = C.LOGSEVERITY_INFO
	LOGSEVERITY_WARNING      = C.LOGSEVERITY_WARNING
	LOGSEVERITY_ERROR        = C.LOGSEVERITY_ERROR
	LOGSEVERITY_ERROR_REPORT = C.LOGSEVERITY_ERROR_REPORT
	LOGSEVERITY_DISABLE      = C.LOGSEVERITY_DISABLE
)

func CEFString(original string) (final *C.cef_string_t) {
	final = (*C.cef_string_t)(C.calloc(1, C.sizeof_cef_string_t))
	charString := C.CString(original)
	defer C.free(unsafe.Pointer(charString))
	C.cef_string_from_utf8(charString, C.strlen(charString), final)
	return final
}

func SetLogger(logger *log.Logger) {
	Logger = logger
}

func _InitializeGlobalCStructures() {
	_MainArgs = (*C.struct__cef_main_args_t)(
		C.calloc(1, C.sizeof_struct__cef_main_args_t))

	_AppHandler = (*C.cef_app_t)(
		C.calloc(1, C.sizeof_cef_app_t))
	C.initialize_app_handler(_AppHandler)

	_ClientHandler = (*C.struct__cef_client_t)(
		C.calloc(1, C.sizeof_struct__cef_client_t))
	C.initialize_client_handler(_ClientHandler)
}

func ExecuteProcess(appHandle unsafe.Pointer) int {
	Logger.Println("ExecuteProcess, args=", os.Args)

	_InitializeGlobalCStructures()
	FillMainArgs(_MainArgs, appHandle)

	// Sandbox info needs to be passed to both cef_execute_process()
	// and cef_initialize().
	// OFF: _SandboxInfo = C.cef_sandbox_info_create()

	var exitCode C.int = C.cef_execute_process(_MainArgs, _AppHandler,
		_SandboxInfo)
	if exitCode >= 0 {
		os.Exit(int(exitCode))
	}
	return int(exitCode)
}

func Initialize(settings Settings) int {
	Logger.Println("Initialize")

	if _MainArgs == nil {
		// _MainArgs structure is initialized and filled in ExecuteProcess.
		// If cef_execute_process is not called, and there is a call
		// to cef_initialize, then it would result in creation of infinite
		// number of processes. See Issue 1199 in CEF:
		// https://code.google.com/p/chromiumembedded/issues/detail?id=1199
		Logger.Println("ERROR: missing a call to ExecuteProcess")
		return 0
	}

	// Initialize cef_settings_t structure.
	var cefSettings *C.struct__cef_settings_t
	cefSettings = (*C.struct__cef_settings_t)(
		C.calloc(1, C.sizeof_struct__cef_settings_t))
	cefSettings.size = C.sizeof_struct__cef_settings_t

	// cache_path
	// ----------
	if settings.CachePath != "" {
		Logger.Println("CachePath=", settings.CachePath)
	}
	var cachePath *C.char = C.CString(settings.CachePath)
	defer C.free(unsafe.Pointer(cachePath))
	C.cef_string_from_utf8(cachePath, C.strlen(cachePath),
		&cefSettings.cache_path)

	// log_severity
	// ------------
	cefSettings.log_severity =
		(C.cef_log_severity_t)(C.int(settings.LogSeverity))

	// log_file
	// --------
	if settings.LogFile != "" {
		Logger.Println("LogFile=", settings.LogFile)
	}
	var logFile *C.char = C.CString(settings.LogFile)
	defer C.free(unsafe.Pointer(logFile))
	C.cef_string_from_utf8(logFile, C.strlen(logFile),
		&cefSettings.log_file)

	// resources_dir_path
	// ------------------
	if settings.ResourcesDirPath == "" && runtime.GOOS != "darwin" {
		// Setting this path is required for the tests to run fine.
		cwd, _ := os.Getwd()
		settings.ResourcesDirPath = cwd
	}
	if settings.ResourcesDirPath != "" {
		Logger.Println("ResourcesDirPath=", settings.ResourcesDirPath)
	}
	var resourcesDirPath *C.char = C.CString(settings.ResourcesDirPath)
	defer C.free(unsafe.Pointer(resourcesDirPath))
	C.cef_string_from_utf8(resourcesDirPath, C.strlen(resourcesDirPath),
		&cefSettings.resources_dir_path)

	// locales_dir_path
	// ----------------
	if settings.LocalesDirPath == "" && runtime.GOOS != "darwin" {
		// Setting this path is required for the tests to run fine.
		cwd, _ := os.Getwd()
		settings.LocalesDirPath = cwd + "/locales"
	}
	if settings.LocalesDirPath != "" {
		Logger.Println("LocalesDirPath=", settings.LocalesDirPath)
	}
	var localesDirPath *C.char = C.CString(settings.LocalesDirPath)
	defer C.free(unsafe.Pointer(localesDirPath))
	C.cef_string_from_utf8(localesDirPath, C.strlen(localesDirPath),
		&cefSettings.locales_dir_path)

        cefSettings.remote_debugging_port = C.int(settings.RemoteDebuggingPort)

	// no_sandbox
	// ----------
	cefSettings.no_sandbox = C.int(1)

	ret := C.cef_initialize(_MainArgs, cefSettings, _AppHandler, _SandboxInfo)
	return int(ret)
}

func CreateBrowser(hwnd unsafe.Pointer, browserSettings BrowserSettings, url string) (browser *C.cef_browser_t) {
	Logger.Println("CreateBrowser, url=", url)

	// Initialize cef_window_info_t structure.
	var windowInfo *C.cef_window_info_t
	windowInfo = (*C.cef_window_info_t)(
		C.calloc(1, C.sizeof_cef_window_info_t))
	FillWindowInfo(windowInfo, hwnd)

	// Initialize cef_browser_settings_t structure.
	var cefBrowserSettings *C.struct__cef_browser_settings_t
	cefBrowserSettings = (*C.struct__cef_browser_settings_t)(
		C.calloc(1, C.sizeof_struct__cef_browser_settings_t))
	cefBrowserSettings.size = C.sizeof_struct__cef_browser_settings_t

	return C.cef_browser_host_create_browser_sync(windowInfo, _ClientHandler, CEFString(url), cefBrowserSettings, nil)
}

func RunMessageLoop() {
	Logger.Println("RunMessageLoop")
	C.cef_run_message_loop()
}

func QuitMessageLoop() {
	Logger.Println("QuitMessageLoop")
	C.cef_quit_message_loop()
}

func Shutdown() {
	Logger.Println("Shutdown")
	C.cef_shutdown()
	// OFF: cef_sandbox_info_destroy(_SandboxInfo)
}
