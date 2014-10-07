// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cefcapi

#include "include/capi/cef_base_capi.h"
#include <stdio.h>

// Print only the first execution of the callback,
// ignore the subsequent.
#define DEBUG_CALLBACK(x) { go_Log(x); }

extern void initialize_cef_base(cef_base_t* base);
