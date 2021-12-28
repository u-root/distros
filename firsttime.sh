#!/bin/bash
set -e

sudo apt-get install build-essential kexec-tools libelf-dev libnl-3-dev libnl-genl-3-dev libssl-dev qemu-system-x86 wireless-tools wpasupplicant libflashrom-dev

echo $GOPATH
pwd
git clone git://github.com/u-root/u-root ../u-root
GO111MODULE=on
echo INSTALL
(cd ../u-root && go install -mod=mod)
(cd ../u-root && go mod tidy -compat=1.17)
echo which u-root
which u-root
