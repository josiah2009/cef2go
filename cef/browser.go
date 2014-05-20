package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"

void ExecuteJavaScript(cef_browser_t* browser, const char* code, const char* script_url, int start_line)
{
    cef_frame_t * frame = browser->get_main_frame(browser);
    cef_string_t * codeCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(code, strlen(code), codeCef);
    cef_string_t * urlVal = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(script_url, strlen(script_url), urlVal);

    frame->execute_java_script(frame, codeCef, urlVal, start_line);

    cef_string_userfree_utf16_free(urlVal);
    cef_string_userfree_utf16_free(codeCef);
}

void LoadURL(cef_browser_t* browser, const char* url)
{
    cef_frame_t * frame = browser->get_main_frame(browser);
    cef_string_t * urlCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(url, strlen(url), urlCef);
    frame->load_url(frame, urlCef);
    cef_string_userfree_utf16_free(urlCef);
}

cef_window_handle_t GetWindowHandle(cef_browser_t* browser)
{
    cef_browser_host_t * host = browser->get_host(browser);
    return host->get_window_handle(host);
}

cef_string_utf8_t * cefStringToUtf8(cef_string_t * source) {
    cef_string_utf8_t * output = cef_string_userfree_utf8_alloc();
    if (source == 0) {
        return output;
    }
    cef_string_to_utf8(source->str, source->length, output);
    return output;
}

cef_string_utf8_t * GetURL(cef_browser_t* browser)
{
    cef_frame_t * frame = browser->get_main_frame(browser);
    cef_string_t * str = frame->get_url(frame);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    // cef_string_userfree_free(str);
    return out;
}

*/
import "C"

import (
	"unsafe"
)

func CreateBrowser(hwnd unsafe.Pointer, browserSettings *BrowserSettings, url string) (browser *Browser) {
	Logger.Println("CreateBrowser, url=", url, "hwnd=", hwnd, "on ui thread?", OnUIThread())

	// Initialize cef_window_info_t structure.
	var windowInfo *C.cef_window_info_t
	windowInfo = (*C.cef_window_info_t)(C.calloc(1, C.sizeof_cef_window_info_t))
	FillWindowInfo(windowInfo, hwnd)
	C.cef_browser_host_create_browser(windowInfo, _ClientHandler, CEFString(url), browserSettings.ToCStruct(), nil)
	Logger.Println("CreateBrowser async", url)
	return globalLifespanHandler.RegisterAndWaitForBrowser(url)
}

type Browser struct {
	cbrowser *C.cef_browser_t
}

func (b *Browser) ExecuteJavaScript(code, url string, startLine int) {
	codeCString := C.CString(code)
	defer C.free(unsafe.Pointer(codeCString))
	urlCString := C.CString(url)
	defer C.free(unsafe.Pointer(urlCString))
	C.ExecuteJavaScript(b.cbrowser, codeCString, urlCString, C.int(startLine))
}

func (b *Browser) LoadURL(url string) {
	urlCString := C.CString(url)
	defer C.free(unsafe.Pointer(urlCString))
	C.LoadURL(b.cbrowser, urlCString)
}

func (b *Browser) GetWindowHandle() C.cef_window_handle_t {
	return C.GetWindowHandle(b.cbrowser)
}

func (b *Browser) GetURL() string {
	cs := C.GetURL(b.cbrowser)
	defer C.cef_string_userfree_utf8_free(cs)
	return C.GoString(cs.str)
}

type BrowserSettings struct {
	///
	// Controls whether file URLs will have access to all URLs. Also configurable
	// using the "allow-universal-access-from-files" command-line switch.
	///
	UniversalAccessFromFileUrls bool

	///
	// Controls whether file URLs will have access to other file URLs. Also
	// configurable using the "allow-access-from-files" command-line switch.
	///
	FileAccessFromFileUrls bool

	///
	// Controls whether web security restrictions (same-origin policy) will be
	// enforced. Disabling this setting is not recommend as it will allow risky
	// security behavior such as cross-site scripting (XSS). Also configurable
	// using the "disable-web-security" command-line switch.
	///
	WebSecurity bool
	///
	// Controls whether WebGL can be used. Note that WebGL requires hardware
	// support and may not work on all systems even when enabled. Also
	// configurable using the "disable-webgl" command-line switch.
	///
	Webgl bool

	///
	// Controls whether content that depends on accelerated compositing can be
	// used. Note that accelerated compositing requires hardware support and may
	// not work on all systems even when enabled. Also configurable using the
	// "disable-accelerated-compositing" command-line switch.
	///
	AcceleratedCompositing bool
}

func (b *BrowserSettings) ToCStruct() (cefBrowserSettings *C.struct__cef_browser_settings_t) {
	// Initialize cef_browser_settings_t structure.
	cefBrowserSettings = (*C.struct__cef_browser_settings_t)(C.calloc(1, C.sizeof_struct__cef_browser_settings_t))
	cefBrowserSettings.size = C.sizeof_struct__cef_browser_settings_t

	cefBrowserSettings.universal_access_from_file_urls = cefStateFromBool(b.UniversalAccessFromFileUrls)
	cefBrowserSettings.file_access_from_file_urls = cefStateFromBool(b.FileAccessFromFileUrls)
	cefBrowserSettings.web_security = cefStateFromBool(b.WebSecurity)
	cefBrowserSettings.webgl = cefStateFromBool(b.Webgl)
	cefBrowserSettings.accelerated_compositing = cefStateFromBool(b.AcceleratedCompositing)
	return cefBrowserSettings
}
