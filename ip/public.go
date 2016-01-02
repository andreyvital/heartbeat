package ip

import "github.com/miekg/dns"

const (
	OpenDNSIPv4 = "208.67.222.222:53"
)

func Public() string {
	c := dns.Client{}

	m := dns.Msg{}
	m.SetQuestion("myip.opendns.com.", dns.TypeA)

	r, _, err := c.Exchange(&m, OpenDNSIPv4)

	if err != nil {
		return ""
	}

	if len(r.Answer) == 0 {
		return ""
	}

	return r.Answer[0].(*dns.A).A.String()
}
