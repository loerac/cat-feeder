CATFEEDER = /usr/local/share/.cat-feeder
CURDIR = $(shell pwd)

configure:
	mkdir -p $(CATFEEDER)
	go get github.com/gorilla/mux

install:
	cp $(CURDIR)/common/common-names.json $(CATFEEDER)/
