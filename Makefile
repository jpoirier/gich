include $(GOROOT)/src/Make.inc

ifeq ($(GOOS),windows)
EXT=.exe
endif

TARG=gich

GOFILES=\
	vars.go\
	gich.go\

CLEANFILES+=gich_test$(EXE)

all:

include $(GOROOT)/src/Make.cmd

test:
	$(QUOTED_GOBIN)/$(GC) -o gich_test.$O gich_test.go vars.go
	$(QUOTED_GOBIN)/$(LD) -o gich_test$(EXE) gich_test.$O

