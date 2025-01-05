package netutils

import "net"

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
