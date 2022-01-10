TARGET = rsass
PREFIX = /usr/local/bin

.PHONY: all clean

all:
	go mod init rsass
	go mod tidy
	GOOS=linux GOARCH=amd64 go build
	sudo cp rsass $(PREFIX)/rsass
	rm -f rsass go.mod

clean:
	sudo rm -rvf $(PREFIX)/rsass $(PREFIX)/keys.pubk ${HOME}/rsass
