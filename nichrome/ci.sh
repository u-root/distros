#!/bin/bash

set -e

echo BUILD usb command
(cd usb && go build .)

echo RUN usb command
./usb/usb --apt=true --fetch=true --dev=/dev/null
