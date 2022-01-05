# light
.POSIX:

include config.mk

all: clean build

build:
	go build

clean:
	rm -f light

install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f light $(DESTDIR)$(PREFIX)/bin
	chown root:video $(DESTDIR)$(PREFIX)/bin/light
	chmod 4750 $(DESTDIR)$(PREFIX)/bin/light

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/light

.PHONY: all build clean install uninstall
