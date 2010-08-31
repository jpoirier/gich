include $(GOROOT)/src/Make.$(GOARCH)

TARG=gogallery

GOFILES=\
        http.go \
        sql.go \
        html.go \
        main.go

include $(GOROOT)/src/Make.cmd



------------
include $(GOROOT)/src/Make.inc

all : wgo

TARG=wgo

GOFILES=\
	const.go \
	Torrent.go \
	Tracker.go \
	Files.go \
	Wire.go \
	Bitfield.go \
	Peer.go \
	PeerMgr.go \
	PieceMgr.go \
	PeerQueue.go \
	PieceData.go \
	Stats.go \
	Listener.go \
	logger.go \
	test.go \

include $(GOROOT)/src/Make.cmd
