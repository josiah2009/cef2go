#include "include/internal/cef_string.h"
#include "include/capi/cef_v8_capi.h"

cef_string_utf8_t * cefSourceToString(cef_string_t * source) {
    cef_string_utf8_t * output = cef_string_userfree_utf8_alloc();
    if (source == 0) {
        return output;
    }
    cef_string_to_utf8(source->str, source->length, output);
    return output;
}

cef_string_userfree_t v8ValueToString(cef_v8value_t * str) {
      return str->get_string_value(str);
}

int32 v8ValueToInt32(cef_v8value_t * i) {
      return i->get_int_value(i);
}

int v8ValueToBool(cef_v8value_t * b) {
      return b->get_bool_value(b);
}
