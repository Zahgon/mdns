// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MIT

package mdns

import (
	"net"

	"github.com/miekg/dns"
)

const (
	// defaultTTL is the default TTL value in returned DNS records in seconds.
	defaultTTL = 120
)

// Zone is the interface used to integrate with the server and
// to serve records dynamically
type Zone interface {
	// Records returns DNS records in response to a DNS question.
	Records(q dns.Question) []dns.RR
}

// MDNSService is used to export a named service by implementing a Zone
type MDNSService struct {
	Instance string   // Instance name (e.g. "hostService name")
	Service  string   // Service name (e.g. "_http._tcp.")
	Domain   string   // If blank, assumes "local"
	HostName string   // Host machine DNS name (e.g. "mymachine.net.")
	Port     int      // Service Port
	IPs      []net.IP // IP addresses for the service's host
	TXT      []string // Service TXT records

	serviceAddr  string // Fully qualified service address
	instanceAddr string // Fully qualified instance address
	enumAddr     string // _services._dns-sd._udp.<domain>
}

// validateFQDN returns an error if the passed string is not a fully qualified
// hdomain name (more specifically, a hostname).
func validateFQDN(s string) error { _ = "STUB: not implemented"; return nil }

// TODO(reddaly): Perform full validation.

// NewMDNSService returns a new instance of MDNSService.
//
// If domain, hostName, or ips is set to the zero value, then a default value
// will be inferred from the operating system.
//
// TODO(reddaly): This interface may need to change to account for "unique
// record" conflict rules of the mDNS protocol.  Upon startup, the server should
// check to ensure that the instance name does not conflict with other instance
// names, and, if required, select a new name.  There may also be conflicting
// hostName A/AAAA records.
func NewMDNSService(instance, service, domain, hostName string, port int, ips []net.IP, txt []string) (*MDNSService, error) {
	_ = "STUB: not implemented"
	// Sanity check inputs
	return nil, nil
}

// Set default domain

// Get host information if no host is specified.

// Try appending the host domain suffix and lookup again
// (required for Linux-based hosts)

// trimDot is used to trim the dots from the start or end of a string
func trimDot(s string) string { _ = "STUB: not implemented"; return "" }

// Records returns DNS records in response to a DNS question.
func (m *MDNSService) Records(q dns.Question) []dns.RR { _ = "STUB: not implemented"; return nil }

func (m *MDNSService) serviceEnum(q dns.Question) []dns.RR { _ = "STUB: not implemented"; return nil }

// serviceRecords is called when the query matches the service name
func (m *MDNSService) serviceRecords(q dns.Question) []dns.RR {
	_ = "STUB: not implemented"
	return nil
}

// Build a PTR response for the service

// Get the instance records

// Return the service record with the instance records

// serviceRecords is called when the query matches the instance name
func (m *MDNSService) instanceRecords(q dns.Question) []dns.RR {
	_ = "STUB: not implemented"
	return nil
}

// Get the SRV, which includes A and AAAA

// Add the TXT record

// TODO(reddaly): IPv4 addresses could be encoded in IPv6 format and
// putinto AAAA records, but the current logic puts ipv4-encodable
// addresses into the A records exclusively.  Perhaps this should be
// configurable?

// Create the SRV Record

// Add the A record

// Add the AAAA record
