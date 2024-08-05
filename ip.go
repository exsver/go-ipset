package ipset

import "net"

func isStringValidCidrOrIP(s string) bool {
	if net.ParseIP(s) != nil {
		return true
	}

	if _, _, err := net.ParseCIDR(s); err == nil {
		return true
	}

	return false
}
