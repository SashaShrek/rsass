TARGET = rsass
PREFIX = /usr/local/bin

.PHONY: all install clean

all:
	go mod init rsass
	go build
	sudo cp rsass $(PREFIX)/rsass
	rm -f rsass go.mod

install:
	sudo cp rsass $(PREFIX)/rsass

clean:
	sudo rm -f $(PREFIX)/rsass $(PREFIX)/keys.pubk rsass keys.pubk go.mod
