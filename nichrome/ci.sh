#!/bin/bash

set -e

echo BUILD usb command
(cd usb && go build .)

go install github.com/u-root/u-root/cmds/exp/tcz

echo RUN usb command
./usb/usb --apt=true --fetch=true --dev=/dev/null
