package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/u-root/webboot/pkg/menu"
	"github.com/u-root/webboot/pkg/wifi"
	"github.com/vishvananda/netlink"
)

// Collect stdout and stderr from the network setup.
// Declare globally because wifi.Connect() triggers
// go routines that might still be running after return.
var wifiStdout, wifiStderr bytes.Buffer

func connected() bool {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	if _, err := client.Get("http://google.com"); err != nil {
		return false
	}
	return true
}

func wirelessIfaceEntries() ([]menu.Entry, error) {
	interfaces, err := netlink.LinkList()
	if err != nil {
		return nil, err
	}

	var ifEntries []menu.Entry
	for _, iface := range interfaces {
		if interfaceIsWireless(iface.Attrs().Name) {
			ifEntries = append(ifEntries, &Interface{label: iface.Attrs().Name})
		}
	}
	return ifEntries, nil
}

func interfaceIsWireless(ifname string) bool {
	devPath := fmt.Sprintf("/sys/class/net/%s/wireless", ifname)
	if _, err := os.Stat(devPath); err != nil {
		return false
	}
	return true
}

func setupNetwork(uiEvents <-chan ui.Event) (bool, error) {
	iface, err := selectNetworkInterface(uiEvents)
	if err != nil {
		return false, err
	} else if menu.IsBackOption(iface) {
		return false, nil
	}

	return selectWirelessNetwork(uiEvents, iface.Label())
}

func selectNetworkInterface(uiEvents <-chan ui.Event) (menu.Entry, error) {
	ifEntries, err := wirelessIfaceEntries()
	if err != nil {
		return nil, err
	}

	iface, err := menu.PromptMenuEntry("Network Interfaces", "Choose an option", ifEntries, uiEvents)
	if err != nil {
		return nil, err
	}

	return iface, nil
}

func selectWirelessNetwork(uiEvents <-chan ui.Event, iface string) (bool, error) {
	worker, err := wifi.NewIWLWorker(&wifiStdout, &wifiStderr, iface)
	if err != nil {
		return false, err
	}

	for {
		progress := menu.NewProgress("Scanning for wifi networks", true)
		networkScan, err := worker.Scan(&wifiStdout, &wifiStderr)
		progress.Close()
		if err != nil {
			return false, err
		}

		netEntries := []menu.Entry{}
		for _, network := range networkScan {
			netEntries = append(netEntries, &Network{info: network})
		}

		entry, err := menu.PromptMenuEntry("Wireless Networks", "Choose an option", netEntries, uiEvents)
		if err != nil {
			return false, err
		} else if menu.IsBackOption(entry) {
			return false, nil
		}

		network, ok := entry.(*Network)
		if !ok {
			return false, fmt.Errorf("Bad menu entry.")
		}

		completed, err := connectWirelessNetwork(uiEvents, worker, network.info)
		if err == io.EOF { // user typed <Ctrl+d> to exit
			return false, err
		} else if err != nil { // connection error
			menu.DisplayResult([]string{err.Error()}, uiEvents)
			continue
		} else if !completed { // user typed <Esc> to go back
			continue
		}

		return true, nil
	}
}

func connectWirelessNetwork(uiEvents <-chan ui.Event, worker wifi.WiFi, network wifi.Option) (bool, error) {
	var setupParams = []string{network.Essid}
	authSuite := network.AuthSuite

	if authSuite == wifi.NotSupportedProto {
		return false, fmt.Errorf("Security protocol is not supported.")
	} else if authSuite == wifi.WpaPsk || authSuite == wifi.WpaEap {
		credentials, err := enterCredentials(uiEvents, authSuite)
		if err != nil {
			return false, err
		} else if credentials == nil {
			return false, nil
		}
		setupParams = append(setupParams, credentials...)
	}

	progress := menu.NewProgress("Connecting to network", true)
	err := worker.Connect(&wifiStdout, &wifiStderr, setupParams...)
	progress.Close()
	if err != nil {
		return false, err
	}

	return true, nil
}

func enterCredentials(uiEvents <-chan ui.Event, authSuite wifi.SecProto) ([]string, error) {
	var credentials []string
	pass, err := menu.PromptTextInput("Enter password:", menu.AlwaysValid, uiEvents)
	if err != nil {
		return nil, err
	} else if pass == "<Esc>" {
		return nil, nil
	}

	credentials = append(credentials, pass)
	if authSuite == wifi.WpaPsk {
		return credentials, nil
	}

	// If not WpaPsk, the network uses WpaEap and also needs an identity
	identity, err := menu.PromptTextInput("Enter identity:", menu.AlwaysValid, uiEvents)
	if err != nil {
		return nil, err
	} else if identity == "<Esc>" {
		return nil, nil
	}

	credentials = append(credentials, identity)
	return credentials, nil
}
