// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cefcapi

#pragma once

#include "handlers/cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_browser_process_handler_capi.h"
#include "include/capi/cef_render_process_handler_capi.h"
#include "include/capi/cef_v8_capi.h"

// ----------------------------------------------------------------------------
// cef_app_t
// ----------------------------------------------------------------------------

///
// Implement this structure to provide handler implementations. Methods will be
// called by the process and/or thread indicated.
///

///
// Provides an opportunity to view and/or modify command-line arguments before
// processing by CEF and Chromium. The |process_type| value will be NULL for
// the browser process. Do not keep a reference to the cef_command_line_t
// object passed to this function. The CefSettings.command_line_args_disabled
// value can be used to start with an NULL command-line object. Any values
// specified in CefSettings that equate to command-line arguments will be set
// before this function is called. Be cautious when using this function to
// modify command-line arguments for non-browser processes as this may result
// in undefined behavior including crashes.
///
void CEF_CALLBACK on_before_command_line_processing(
        struct _cef_app_t* self, const cef_string_t* process_type,
        struct _cef_command_line_t* command_line) {
    DEBUG_CALLBACK("on_before_command_line_processing\n");
}

///
// Provides an opportunity to register custom schemes. Do not keep a reference
// to the |registrar| object. This function is called on the main thread for
// each process and the registered schemes should be the same across all
// processes.
///
void CEF_CALLBACK on_register_custom_schemes(
        struct _cef_app_t* self,
        struct _cef_scheme_registrar_t* registrar) {
    DEBUG_CALLBACK("on_register_custom_schemes\n");
}

void CEF_CALLBACK on_context_initialized(
      struct _cef_browser_process_handler_t* self) {
    DEBUG_CALLBACK("on_context_initialized!\n");
}

///
// Return the handler for resource bundle events. If
// CefSettings.pack_loading_disabled is true (1) a handler must be returned.
// If no handler is returned resources will be loaded from pack files. This
// function is called by the browser and render processes on multiple threads.
///
struct _cef_resource_bundle_handler_t*
        CEF_CALLBACK get_resource_bundle_handler(struct _cef_app_t* self) {
    DEBUG_CALLBACK("get_resource_bundle_handler\n");
    return NULL;
}

///
// Return the handler for functionality specific to the browser process. This
// function is called on multiple threads in the browser process.
///
struct _cef_browser_process_handler_t* 
        CEF_CALLBACK get_browser_process_handler(struct _cef_app_t* self) {
    DEBUG_CALLBACK("get_browser_process_handler\n");
    return NULL;
}

int CEF_CALLBACK cef_v8handler_execute(struct _cef_v8handler_t* self,
      const cef_string_t* name, struct _cef_v8value_t* object,
      size_t argumentsCount, struct _cef_v8value_t* const* arguments,
      struct _cef_v8value_t** retval, cef_string_t* exception) {
    DEBUG_CALLBACK("v8handler->execute\n");
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

///
// Return the handler for functionality specific to the render process. This
// function is called on the render process main thread.
///
struct _cef_render_process_handler_t*
        CEF_CALLBACK get_render_process_handler(struct _cef_app_t* self) {
    DEBUG_CALLBACK("get_render_process_handler\n");
    cef_render_process_handler_t* renderProcessHandler = (cef_render_process_handler_t*)calloc(1, sizeof(cef_render_process_handler_t));
    renderProcessHandler->base.size = sizeof(cef_render_process_handler_t);
    initialize_cef_base((cef_base_t*) renderProcessHandler);
    renderProcessHandler->on_web_kit_initialized = cef_render_process_handler_t_on_webkit_initialized;
    return renderProcessHandler;
}

void initialize_app_handler(cef_app_t* app) {
    DEBUG_CALLBACK("initialize_app_handler\n");
    app->base.size = sizeof(cef_app_t);
    initialize_cef_base((cef_base_t*)app);
    // callbacks
    app->on_before_command_line_processing = on_before_command_line_processing;
    app->on_register_custom_schemes = on_register_custom_schemes;
    DEBUG_CALLBACK("set get_resource_bundle_handler\n");
    app->get_resource_bundle_handler = get_resource_bundle_handler;
    DEBUG_CALLBACK("set get_browser_process_handler\n");
    app->get_browser_process_handler = get_browser_process_handler;
    DEBUG_CALLBACK("set render_browser_process_handler\n");
    app->get_render_process_handler = get_render_process_handler;
}
