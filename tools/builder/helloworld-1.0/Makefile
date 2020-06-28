DESTDIR ?=
PREFIX ?= /usr/local

helloworld:
	gcc main.c -o helloworld

install: helloworld
	install -m 0755 -d $(DESTDIR)/$(PREFIX)/bin
	install -m 0755 helloworld $(DESTDIR)/$(PREFIX)/bin

uninstall:
	@echo $(DESTDIR)$(PREFIX)bin/helloworld 
	rm -fv  $(DESTDIR)/$(PREFIX)/bin/helloworld

# Without .PHONY,  clean command line will never execute
.PHONY: clean
clean:
	rm -f helloworld 
	rm -f helloworld.exe
