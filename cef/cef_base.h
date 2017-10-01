// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cefcapi

#include "include/capi/cef_base_capi.h"
#include <stdio.h>

// Print only the first execution of the callback,
// ignore the subsequent.
#define DEBUG_CALLBACK(x) { go_Log(x); }
#define DEBUG_POINTER(x, p) { go_LogPointer(x, p); }

extern void initialize_cef_base_ref_counted(cef_base_ref_counted_t* base);
