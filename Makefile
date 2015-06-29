.PHONY: detect_os Linux clean

UNAME_S = $(shell uname -s)
CEF=github.com/mmatey/cef2go/cef
INC=-I. \
	-I/usr/include/gtk-2.0 \
	-I/usr/include/glib-2.0 \
	-I/usr/include/cairo \
	-I/usr/include/pango-1.0 \
	-I/usr/include/gdk-pixbuf-2.0 \
	-I/usr/include/atk-1.0 \
	-I/usr/lib/x86_64-linux-gnu/glib-2.0/include \
	-I/usr/lib/x86_64-linux-gnu/gtk-2.0/include \
	-I/usr/lib/i386-linux-gnu/gtk-2.0/include \
	-I/usr/lib/i386-linux-gnu/glib-2.0/include \
	-I/usr/lib64/glib-2.0/include \
	-I/usr/lib64/gtk-2.0/include
export CC=gcc $(INC)
export CGO_LDFLAGS=-L $(PWD)/Release -lcef


detect_os:
	make $(UNAME_S)

build:
	go build -ldflags "-r ." -o ../../bin/cef2go main_linux.go

Linux:
	make clean
	go get -u $(CEF)
	go install $(CEF)
	go build -ldflags "-r ." -o ../../bin/cef2go main_linux.go

clean:
	go clean -i $(CEF) || true
