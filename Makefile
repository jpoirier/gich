include $(GOROOT)/src/Make.inc

all: giche

TARG=giche

GOFILES=\
	giche.go \

include $(GOROOT)/src/Make.cmd
