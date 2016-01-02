package ip

import "net"

// Internal returns the internal/local host IP address
// http://stackoverflow.com/a/31551220
func Internal() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipn, ok := addr.(*net.IPNet); ok && !ipn.IP.IsLoopback() {
			if ipn.IP.To4() != nil {
				return ipn.IP.String()
			}
		}
	}

	return ""
}
