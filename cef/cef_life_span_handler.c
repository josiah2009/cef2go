// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi


#include <stdlib.h>
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"


int CEF_CALLBACK cef_life_span_handler_t_on_before_popup(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser,
        struct _cef_frame_t* frame,
        const cef_string_t* target_url,
        const cef_string_t* target_frame_name,
        const struct _cef_popup_features_t* popupFeatures,
        struct _cef_window_info_t* windowInfo,
        struct _cef_client_t** client,
        struct _cef_browser_settings_t* settings,
        int* no_javascript_access) {
    printf("OnBeforePopup\n");
    return 0;
}


void CEF_CALLBACK cef_life_span_handler_t_on_after_created(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    go_OnAfterCreated(self, browser);
}

int CEF_CALLBACK cef_life_span_handler_t_run_modal(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    printf("RunModal\n");
    return 0;
}

int CEF_CALLBACK cef_life_span_handler_t_do_close(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    printf("DoClose\n");
    return 0;
}

void CEF_CALLBACK cef_life_span_handler_t_on_before_close(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    printf("BeforeClose\n");
}


void initialize_life_span_handler(struct _cef_life_span_handler_t* lifeHandler) {
    printf("initialize_life_span_handler\n");
    lifeHandler->base.size = sizeof(cef_life_span_handler_t);
    initialize_cef_base((cef_base_t*) lifeHandler);
    // callbacks
    lifeHandler->on_before_popup = cef_life_span_handler_t_on_before_popup;
    lifeHandler->on_after_created = cef_life_span_handler_t_on_after_created;
    lifeHandler->run_modal = cef_life_span_handler_t_run_modal;
    lifeHandler->do_close = cef_life_span_handler_t_do_close;
    lifeHandler->on_before_close = cef_life_span_handler_t_on_before_close;
}
