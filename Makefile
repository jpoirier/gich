include $(GOROOT)/src/Make.inc

ifeq ($(GOOS),windows)
EXT=.exe
endif

TARG=giche

GOFILES=\
	vars.go\
	giche.go\

CLEANFILES+=giche_test$(EXE)

all:

include $(GOROOT)/src/Make.cmd

test:
	$(QUOTED_GOBIN)/$(GC) -o giche_test.$O giche_test.go vars.go
	$(QUOTED_GOBIN)/$(LD) -o giche_test$(EXE) giche_test.$O

