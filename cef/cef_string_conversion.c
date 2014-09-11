#include "include/internal/cef_string.h"

cef_string_utf8_t * cefSourceToString(cef_string_t * source) {
    cef_string_utf8_t * output = cef_string_userfree_utf8_alloc();
    if (source == 0) {
        return output;
    }
    cef_string_to_utf8(source->str, source->length, output);
    return output;
}
