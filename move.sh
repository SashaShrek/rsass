go mod init rsass
#Linux
GOOS=linux GOARCH=amd64 go build
tar -cvf rsass64.tar rsass LICENSE README.md Makefile && gzip -6 rsass64.tar
scp rsass64.tar.gz sssha256@188.227.84.204:/var/www/rsass/linux/
rm rsass rsass64.tar.gz

#Windows
GOOS=windows GOARCH=amd64 go build
zip -6 rsass64.zip rsass.exe LICENSE README.md
scp rsass64.zip sssha256@188.227.84.204:/var/www/rsass/windows/
rm rsass.exe rsass64.zip go.mod