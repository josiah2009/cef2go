// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi


#include <stdlib.h>
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"


void CEF_CALLBACK cef_life_span_handler_t_on_after_created(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    printf("on_after_created\n");
    go_OnAfterCreated(self, browser);
}

void initialize_life_span_handler(struct _cef_life_span_handler_t* lifeHandler) {
    lifeHandler->base.size = sizeof(cef_life_span_handler_t);
    // callbacks
    lifeHandler->on_after_created = cef_life_span_handler_t_on_after_created;
}
