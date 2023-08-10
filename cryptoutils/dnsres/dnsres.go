package dnsres

import (
	"net"
)

// ResolveDomain resolves a domain to its associated IP addresses.
func ResolveDomain(domain string) ([]net.IP, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}
	return ips, nil
}
