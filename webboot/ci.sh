#!/bin/bash
set -e

go build -mod=mod .
gopath=`go env GOPATH`
ls -l $gopath

./webboot --wpa-version=2.9
if [ ! -f "/tmp/initramfs.linux_amd64.cpio" ]; then
    echo "Initrd was not created."
    exit 1
fi


wget -O TinyCorePure64.iso https://github.com/u-root/webboot/blob/main/pkg/bootiso/testdata/TinyCorePure64.iso?raw=true
rm -f ../bootiso/testdata/TinyCorePure64.iso
ln TinyCorePure64.iso  ../bootiso/testdata/
mkdir -p ../cmds/webboot/testdata/dirlevel1/dirlevel2/
rm -f ../cmds/webboot/testdata/dirlevel1/dirlevel2/TinyCorePure64.iso
ln TinyCorePure64.iso ../cmds/webboot/testdata/dirlevel1/dirlevel2/
# until it is fixed ... (cd ../cmds/webboot && go test -v)
# (cd pkg/menu && go test -v)
# (cd pkg/bootiso && sudo -E env "PATH=$PATH" go test -v) # need sudo to mount the test iso
