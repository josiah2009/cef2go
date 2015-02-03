#include <string.h>
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"
#include "include/capi/cef_render_handler_capi.h"
#include "cef_base.h"

typedef struct cef_go_client {
    
  cef_display_handler_t *display_handler;
  cef_life_span_handler_t *life_span_handler;
  cef_render_handler_t *render_handler;

} cef_go_client;

cef_go_client * go_client;

int CEF_CALLBACK cef_display_handler_t_on_console_message(
      struct _cef_display_handler_t* self,
      struct _cef_browser_t* browser, const cef_string_t* message,
      const cef_string_t* source, int line) {
    go_OnConsoleMessage(browser, message, source, line);
    return 1;
}

void CEF_CALLBACK cef_life_span_handler_t_on_after_created(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    //DEBUG_CALLBACK("client->LifeSpanHandler->on_after_created\n");
    go_OnAfterCreated(self, browser->get_identifier(browser), browser);
}


int CEF_CALLBACK cef_render_handler_t_get_root_screen_rect(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, cef_rect_t* rect) {
      //DEBUG_CALLBACK("render_handler->get_root_screen_rect");
      return go_RenderHandlerGetRootScreenRect(browser->get_identifier(browser), rect);
}

int CEF_CALLBACK cef_render_handler_t_get_view_rect(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, cef_rect_t* rect) {
      //DEBUG_CALLBACK("render_handler->get_view_rect");
      return go_RenderHandlerGetViewRect(browser->get_identifier(browser), rect);
}

int CEF_CALLBACK cef_render_handler_t_get_screen_point(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, int viewX, int viewY, int* screenX, int* screenY) {
      //DEBUG_CALLBACK("render_handler->get_screen_point");
      return go_RenderHandlerGetScreenPoint(browser->get_identifier(browser), viewX, viewY, screenX, screenY);
}

int CEF_CALLBACK cef_render_handler_t_get_screen_info(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, struct _cef_screen_info_t* info) {
      //DEBUG_CALLBACK("render_handler->get_screen_info");
      return go_RenderHandlerGetScreenInfo(browser->get_identifier(browser), info);
}

void CEF_CALLBACK cef_render_handler_t_on_popup_show(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, int show) {
      //DEBUG_CALLBACK("render_handler->on_popup_show");
      go_RenderHandlerOnPopupShow(browser->get_identifier(browser), show);
}

void CEF_CALLBACK cef_render_handler_t_on_popup_size(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, const cef_rect_t* rect) {
      //DEBUG_CALLBACK("render_handler->on_popup_size");
      go_RenderHandlerOnPopupSize(browser->get_identifier(browser), rect);
}

void CEF_CALLBACK cef_render_handler_t_on_paint(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, cef_paint_element_type_t type,
      size_t dirtyRectsCount, cef_rect_t const* dirtyRects, const void* buffer,
      int width, int height) {
      //DEBUG_CALLBACK("render_handler->on_paint");
      go_RenderHandlerOnPaint(browser->get_identifier(browser), type, dirtyRectsCount, dirtyRects, buffer, width, height);
}

void CEF_CALLBACK cef_render_handler_t_on_cursor_change(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, cef_cursor_handle_t cursor,
      cef_cursor_type_t type,
      const struct _cef_cursor_info_t* custom_cursor_info) {
      //DEBUG_CALLBACK("render_handler->on_cursor_change");
      go_RenderHandlerOnCursorChange(browser->get_identifier(browser), cursor, type, custom_cursor_info);
}

void CEF_CALLBACK cef_render_handler_t_on_scroll_offset_changed(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser) {
      DEBUG_CALLBACK("render_handler->on_scroll_offset_changed");
      go_RenderHandlerOnScrollOffsetChanged(browser->get_identifier(browser));
}


void initialize_display_handler() {
    cef_display_handler_t* displayHandler = (cef_display_handler_t*)calloc(1, sizeof(cef_display_handler_t));
    displayHandler->base.size = sizeof(cef_display_handler_t);
    initialize_cef_base((cef_base_t*) displayHandler);
    // callbacks
    displayHandler->on_console_message = cef_display_handler_t_on_console_message;
    go_client->display_handler = displayHandler;
}

void initialize_life_span_handler() {
    cef_life_span_handler_t* lifeHandler = (cef_life_span_handler_t*)calloc(1, sizeof(cef_life_span_handler_t));
    //DEBUG_CALLBACK("client->initialize_life_span_handler\n");
    lifeHandler->base.size = sizeof(cef_life_span_handler_t);
    initialize_cef_base((cef_base_t*) lifeHandler);
    // callbacks
    lifeHandler->on_after_created = cef_life_span_handler_t_on_after_created;
    go_client->life_span_handler = lifeHandler;
}

void initialize_render_handler() {
    //DEBUG_CALLBACK("initialize_render_handler");
    cef_render_handler_t* renderHandler = (cef_render_handler_t*)calloc(1, sizeof(cef_render_handler_t));
    renderHandler->base.size = sizeof(cef_render_handler_t);
    initialize_cef_base((cef_base_t*) renderHandler);
    // callbacks
    renderHandler->get_root_screen_rect = cef_render_handler_t_get_root_screen_rect;
    renderHandler->get_view_rect = cef_render_handler_t_get_view_rect;
    renderHandler->get_screen_point = cef_render_handler_t_get_screen_point;
    renderHandler->get_screen_info = cef_render_handler_t_get_screen_info;
    renderHandler->on_popup_show = cef_render_handler_t_on_popup_show;
    renderHandler->on_popup_size = cef_render_handler_t_on_popup_size;
    renderHandler->on_paint = cef_render_handler_t_on_paint;
    renderHandler->on_cursor_change = cef_render_handler_t_on_cursor_change;
    renderHandler->on_scroll_offset_changed = cef_render_handler_t_on_scroll_offset_changed;
    DEBUG_POINTER("render_handler", renderHandler);
    //go_AddRef((void *) renderHandler);
    go_client->render_handler = renderHandler;
}

struct _cef_display_handler_t* CEF_CALLBACK get_display_handler(
        struct _cef_client_t* self) {
    //DEBUG_CALLBACK("get_display_handler");
    go_AddRef((void *) go_client->display_handler);
    return go_client->display_handler;
}

struct _cef_life_span_handler_t* CEF_CALLBACK get_life_span_handler(
        struct _cef_client_t* self) {
    //DEBUG_CALLBACK("get_life_span_handler");
    go_AddRef((void *) go_client->life_span_handler);
    return go_client->life_span_handler;
}

struct _cef_render_handler_t* CEF_CALLBACK get_render_handler(
        struct _cef_client_t* self) {
    //DEBUG_CALLBACK("get_render_handler");
    go_AddRef((void *) go_client->render_handler);
    return go_client->render_handler;
}

void initialize_client_handler(struct _cef_client_t* client) {
    DEBUG_POINTER("initialize_client_handler", client);
    go_client = (cef_go_client*)calloc(1, sizeof(cef_go_client));
    initialize_display_handler();
    initialize_life_span_handler();
    initialize_render_handler();

    client->base.size = sizeof(cef_client_t);
    initialize_cef_base((cef_base_t*)client);
    // callbacks
    //DEBUG_CALLBACK("set_display_handler");
    client->get_display_handler = get_display_handler;
    client->get_life_span_handler = get_life_span_handler;
    client->get_render_handler = get_render_handler;
}
