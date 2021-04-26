TARGET = rsass
PREFIX = /usr/local/bin

.PHONY: all clean

all:
	go mod init rsass
	go build
	sudo cp rsass $(PREFIX)/rsass
	rm -f rsass go.mod

clean:
	sudo rm -f $(PREFIX)/rsass $(PREFIX)/keys.pubk rsass keys.pubk go.mod
