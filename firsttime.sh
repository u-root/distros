#!/bin/bash
set -e

sudo apt-get install build-essential kexec-tools libelf-dev libnl-3-dev libnl-genl-3-dev libssl-dev qemu-system-x86 wireless-tools wpasupplicant

echo $GOPATH
pwd
go get -u github.com/u-root/u-root
GO111MODULE=on
echo INSTALL
go install -mod=mod github.com/u-root/u-root
echo which u-root
which u-root
