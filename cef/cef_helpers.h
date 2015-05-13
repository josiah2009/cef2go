#include "include/internal/cef_string.h"
#include "include/capi/cef_v8_capi.h"

extern cef_string_utf8_t * cefSourceToString(cef_string_t * source);
extern cef_string_userfree_t v8ValueToString(cef_v8value_t * str);
extern int32 v8ValueToInt32(cef_v8value_t * i);
extern int v8ValueToBool(cef_v8value_t * b);
extern double v8ValueToDouble(cef_v8value_t * d);
extern void setCefRectDimensions(cef_rect_t * rect, int x, int y, int width, int height);
