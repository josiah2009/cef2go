package cef

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
int XErrorHandlerImpl(Display *display, XErrorEvent *event) {
  go_Log("X error received");
  //LOG(WARNING)
  //      << "X error received: "
  //      << "type " << event->type << ", "
  //      << "serial " << event->serial << ", "
  //      << "error_code " << static_cast<int>(event->error_code) << ", "
  //      << "request_code " << static_cast<int>(event->request_code) << ", "
  //      << "minor_code " << static_cast<int>(event->minor_code);
  return 0;
}
int XIOErrorHandlerImpl(Display *display) {
  return 0;
}
void XlibRegisterHandlers() {
    XSetErrorHandler(XErrorHandlerImpl);
    XSetIOErrorHandler(XIOErrorHandlerImpl);
}
*/
import "C"

func XlibRegisterHandlers() {
	C.XlibRegisterHandlers()
}
