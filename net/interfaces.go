package net

import "net"

func Interfaces() (res []Interface) {
	ifaces, err := net.Interfaces()

	if err != nil {
		return res
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()

		if err != nil {
			continue
		}

		for _, addr := range addrs {
			switch addr.(type) {
			case *net.IPNet:
				ipn := addr.(*net.IPNet)

				if ipn.IP.DefaultMask() == nil {
					continue
				}

				res = append(res, Interface{
					Index:    iface.Index,
					Name:     iface.Name,
					Addr:     ipn.IP.String(),
					Loopback: ipn.IP.IsLoopback(),
				})
			}
		}
	}

	return res
}
