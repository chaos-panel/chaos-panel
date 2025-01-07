package netutils

import (
	"net"
	"strings"
)

// get ipmac of lan
func GetLanMacAll() []string {
	returns := make([]string, 0, len(GetLanAll()))
	for _, v := range GetLanAll() {
		returns = append(returns, v)
	}
	return returns
}

// get ipv4 all of lan
func GetLanIpv4All() []string {
	returns := make([]string, 0, len(GetLanAll()))
	for k := range GetLanAll() {
		returns = append(returns, k)
	}
	return returns
}

// get ipv4 all of lan: map[ipv4]mac
func GetLanAll() map[string]string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	lanAll := make(map[string]string)
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() || ip.IsLinkLocalUnicast() ||
				ip.Equal(net.ParseIP("0.0.0.0")) || ip.Equal(net.ParseIP("127.0.0.1")) {
				continue
			}
			if strings.Contains(i.Name, "#") || strings.Contains(i.Name, "virtual") || strings.Contains(i.Name, "docker") || strings.Contains(i.Name, "veth") || strings.Contains(i.Name, "cali") ||
				strings.Contains(i.Name, "vbox") || strings.Contains(i.Name, "vmnet") || strings.Contains(i.Name, "vnic") ||
				strings.Contains(i.Name, "flannel") || strings.Contains(i.Name, "br-") || strings.Contains(i.Name, "macvlan") ||
				strings.Contains(i.Name, "weave") || strings.Contains(i.Name, "virbr") || strings.Contains(i.Name, "ovs") ||
				strings.Contains(i.Name, "br0") || strings.Contains(i.Name, "br1") || strings.Contains(i.Name, "tun") ||
				strings.Contains(i.Name, "wg") || strings.Contains(i.Name, "tap") {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue // 非ipv4
			}
			mac := i.HardwareAddr
			if mac == nil {
				continue
			}
			lanAll[ip.String()] = mac.String()
		}
	}
	return lanAll
}

func IsIpv4(ip string) bool {
	return net.ParseIP(ip).To4() != nil
}
