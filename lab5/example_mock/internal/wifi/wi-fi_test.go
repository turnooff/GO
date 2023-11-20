package wifi_test

import (
	"errors"
	myWifi "example_mock/internal/wifi"
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/mdlayher/wifi"
	//"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

type rowTestSysInfo struct {
	addrs       []string
	errExpected error
}

var testTable = []rowTestSysInfo{
	{
		addrs: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
	},
	{
		errExpected: errors.New("ExpectedError"),
	},
}

func TestGetAddresses(t *testing.T) {
	mockWifi := NewWiFi(t)
	wifiService := myWifi.Service{WiFi: mockWifi}
	for i, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.addrs), row.errExpected)
		actualAddrs, err := wifiService.GetAddresses()
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			continue
		}
		if !reflect.DeepEqual(actualAddrs, parseMACs(row.addrs)) {
			t.Errorf("expected names to be %v, got %v", actualAddrs, parseMACs(row.addrs))
		}
	}
}

func TestGetNames(t *testing.T) {
	mockWifi := NewWiFi(t)
	wifiService := myWifi.Service{WiFi: mockWifi}
	mockWifi.On("Interfaces").Unset()
	mockWifi.On("Interfaces").Return([]*wifi.Interface{{Name: "WiFiName1"}, {Name: "WiFiName2"}}, nil)
	names, err := wifiService.GetNames()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expectedNames := []string{"WiFiName1", "WiFiName2"}
	if !reflect.DeepEqual(names, expectedNames) {
		t.Errorf("expected names to be %v, got %v", expectedNames, names)
	}
	// _, err = myWifi.WiFi.Interfaces(mockWifi)
	// if err != nil {
	// 	t.Errorf("unexpected error: %v", err)
	// }
}

func TestNew(t *testing.T) {
	mockWifi := NewWiFi(t)
	wifiService := myWifi.Service{WiFi: mockWifi}
	result := myWifi.New(mockWifi)
	if result != wifiService {
		t.Errorf("expected to be %v, got %v", wifiService, result)
	}
}

func mockIfaces(addrs []string) []*wifi.Interface {
	var interfaces []*wifi.Interface
	for i, addrStr := range addrs {
		hwAddr := parseMAC(addrStr)
		if hwAddr == nil {
			continue
		}
		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         fmt.Sprintf("eth%d", i+1),
			HardwareAddr: hwAddr,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		}
		interfaces = append(interfaces, iface)
	}
	return interfaces
}

func parseMACs(macStr []string) []net.HardwareAddr {
	var addrs []net.HardwareAddr
	for _, addr := range macStr {
		addrs = append(addrs, parseMAC(addr))
	}
	return addrs
}

func parseMAC(macStr string) net.HardwareAddr {
	hwAddr, err := net.ParseMAC(macStr)
	if err != nil {
		return nil
	}
	return hwAddr
}
