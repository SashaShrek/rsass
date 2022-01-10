#!/usr/bin/env bash

#COLORS
COLOR_RED=`tput setaf 1`
COLOR_GREEN=`tput setaf 2`
COLOR_YELLOW=`tput setaf 3`
COLOR_BLUE=`tput setaf 4`
COLOR_RESET=`tput sgr0`

showLogo(){
    echo -e "${COLOR_RED}"
    echo -e "RSASS"
    echo -e "Install script"
    echo -e "${COLOR_RESET}"
}

#START
showLogo
echo "Downloading archive..."
wget -vS https://github.com/SashaShrek/rsass/releases/download/v.1.0.0/rsass64.tar.gz
echo -e "${COLOR_GREEN}OK${COLOR_RESET}"
echo "Unpacking..."
gunzip rsass64.tar.gz && tar -xvf rsass64.tar && rm -rf rsass64.tar
mv rsass ${HOME}/ && cd ${HOME}/rsass/
echo -e "${COLOR_GREEN}OK${COLOR_RESET}"
echo "Installing..."
make
echo -e "${COLOR_GREEN}OK${COLOR_RESET}"
echo "Cleaning..."
rm -rv file rsass.go crypto.go crt_keys.go
echo -e "${COLOR_GREEN}OK${COLOR_RESET}"