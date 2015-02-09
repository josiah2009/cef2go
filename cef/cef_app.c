#include <string.h>
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_browser_process_handler_capi.h"
#include "include/capi/cef_render_process_handler_capi.h"
#include "include/capi/cef_v8_capi.h"
#include "cef_base.h"

// ----------------------------------------------------------------------------
// cef_app_t
// ----------------------------------------------------------------------------


int CEF_CALLBACK cef_v8handler_execute(struct _cef_v8handler_t* self,
      const cef_string_t* name, struct _cef_v8value_t* object,
      size_t argumentsCount, struct _cef_v8value_t* const* arguments,
      struct _cef_v8value_t** retval, cef_string_t* exception) {
    //DEBUG_CALLBACK("v8handler->execute\n");
    return go_V8HandlerExecute(name, object, argumentsCount, arguments, retval, exception);
}

// Set up the javascript cef extensions
void CEF_CALLBACK cef_render_process_handler_t_on_webkit_initialized(struct _cef_render_process_handler_t* self) {
    cef_v8handler_t* goV8Handler = (cef_v8handler_t*)calloc(1, sizeof(cef_v8handler_t));
    goV8Handler->base.size = sizeof(cef_v8handler_t);
    initialize_cef_base((cef_base_t*) goV8Handler);
    goV8Handler->execute = cef_v8handler_execute;
    go_RenderProcessHandlerOnWebKitInitialized(goV8Handler);
}

// Set up the context initialized callback
void CEF_CALLBACK cef_browser_process_handler_t_on_context_initialized(struct _cef_browser_process_handler_t* self) {
    go_BrowserProcessHandlerOnContextInitialized();
}

///
// Return the handler for functionality specific to the render process. This
// function is called on the render process main thread.
///
struct _cef_render_process_handler_t* CEF_CALLBACK get_render_process_handler(struct _cef_app_t* self) {
    //DEBUG_POINTER("get_render_process_handler", self);
    cef_render_process_handler_t* renderProcessHandler = (cef_render_process_handler_t*)calloc(1, sizeof(cef_render_process_handler_t));
    renderProcessHandler->base.size = sizeof(cef_render_process_handler_t);
    initialize_cef_base((cef_base_t*) renderProcessHandler);
    renderProcessHandler->on_web_kit_initialized = cef_render_process_handler_t_on_webkit_initialized;
    return renderProcessHandler;
}


///
// Return the handler for functionality specific to the browser process. This
// function is called on multiple threads in the browser process.
///
struct _cef_browser_process_handler_t* CEF_CALLBACK get_browser_process_handler(struct _cef_app_t* self) {
    cef_browser_process_handler_t* browserProcessHandler = (cef_browser_process_handler_t*)calloc(1, sizeof(cef_browser_process_handler_t));
    browserProcessHandler->base.size = sizeof(cef_browser_process_handler_t);
    initialize_cef_base((cef_base_t*) browserProcessHandler);
    browserProcessHandler->on_context_initialized = cef_browser_process_handler_t_on_context_initialized;
    return browserProcessHandler;
}

void initialize_app_handler(cef_app_t* app) {
    // DEBUG_POINTER("initialize_app_handler", app);
    app->base.size = sizeof(cef_app_t);
    initialize_cef_base((cef_base_t*)app);
    // callbacks
    app->get_render_process_handler = get_render_process_handler;
    app->get_browser_process_handler = get_browser_process_handler;
}
