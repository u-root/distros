package main

import (
	"fmt"

	"github.com/u-root/webboot/pkg/menu"
	"github.com/u-root/webboot/pkg/wifi"
)

type Distro struct {
	url          string
	bootConfig   string
	kernelParams string
}

var supportedDistros = map[string]Distro{
	"Tinycore": Distro{
		url:          "http://tinycorelinux.net/11.x/x86_64/release/TinyCorePure64-11.1.iso",
		bootConfig:   "syslinux",
		kernelParams: "iso=",
	},
	"Ubuntu": Distro{
		url:          "https://releases.ubuntu.com/20.04.1/ubuntu-20.04.1-desktop-amd64.iso",
		bootConfig:   "grub",
		kernelParams: "iso-scan/filename=",
	},
}

// ISO contains information of the iso user want to boot
type ISO struct {
	label string
	path  string
}

var _ = menu.Entry(&ISO{})

// Label is the string this iso displays in the menu page.
func (i *ISO) Label() string {
	return i.label
}

// Config represents one kind of configure of booting an iso
type Config struct {
	label string
}

var _ = menu.Entry(&Config{})

// Label is the string this iso displays in the menu page.
func (c *Config) Label() string {
	return c.label
}

// DownloadOption let user download an iso then boot it
type DownloadOption struct {
}

var _ = menu.Entry(&DownloadOption{})

// Label is the string this iso displays in the menu page.
func (d *DownloadOption) Label() string {
	return "Download an ISO"
}

// DirOption represents a directory under cache directory
// it displays it's sub-directory or iso files
type DirOption struct {
	label string
	path  string
}

var _ = menu.Entry(&DirOption{})

// Label is the string this option displays in the menu page.
func (d *DirOption) Label() string {
	return d.label
}

type Interface struct {
	label string
}

func (i *Interface) Label() string {
	return i.label
}

type Network struct {
	info wifi.Option
}

func (n *Network) Label() string {
	switch n.info.AuthSuite {
	case wifi.NoEnc:
		return fmt.Sprintf("%s: No Passphrase\n", n.info.Essid)
	case wifi.WpaPsk:
		return fmt.Sprintf("%s: WPA-PSK (only passphrase)\n", n.info.Essid)
	case wifi.WpaEap:
		return fmt.Sprintf("%s: WPA-EAP (passphrase and identity)\n", n.info.Essid)
	case wifi.NotSupportedProto:
		return fmt.Sprintf("%s: Not a supported protocol\n", n.info.Essid)
	}
	return "Invalid wifi network."
}

// BackOption let user back to the upper menu
type BackOption struct {
}

var _ = menu.Entry(&BackOption{})

// Label is the string this iso displays in the menu page.
func (b *BackOption) Label() string {
	return "Go Back"
}
