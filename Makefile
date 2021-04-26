TARGET = rsass
PREFIX = /usr/local/bin

.PHONY: all install uninstall clean

all:
	go mod init rsass
	go build
	sudo cp rsass $(PREFIX)/rsass
	rm -f rsass go.mod

install:
	sudo cp rsass $(PREFIX)/rsass

uninstall:
	rm -f rsass keys.pubk

clean:
	sudo rm -f $(PREFIX)/rsass $(PREFIX)/keys.pubk rsass keys.pubk go.mod
