// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MIT

package mdns

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/miekg/dns"
)

// ServiceEntry is returned after we query for a service
type ServiceEntry struct {
	Name         string
	Host         string
	AddrV4       net.IP
	AddrV6       net.IP // @Deprecated
	AddrV6IPAddr *net.IPAddr
	Port         int
	Info         string
	InfoFields   []string

	Addr net.IP // @Deprecated

	hasTXT bool
	sent   bool
}

// complete is used to check if we have all the info we need
func (s *ServiceEntry) complete() bool { _ = "STUB: not implemented"; return false }

// QueryParam is used to customize how a Lookup is performed
type QueryParam struct {
	Service             string               // Service to lookup
	Domain              string               // Lookup domain, default "local"
	Timeout             time.Duration        // Lookup timeout, default 1 second
	Interface           *net.Interface       // Multicast interface to use
	Entries             chan<- *ServiceEntry // Entries Channel
	WantUnicastResponse bool                 // Unicast response desired, as per 5.4 in RFC
	DisableIPv4         bool                 // Whether to disable usage of IPv4 for MDNS operations. Does not affect discovered addresses.
	DisableIPv6         bool                 // Whether to disable usage of IPv6 for MDNS operations. Does not affect discovered addresses.
	Logger              *log.Logger          // Optionally provide a *log.Logger to better manage log output.
}

// DefaultParams is used to return a default set of QueryParam's
func DefaultParams(service string) *QueryParam { _ = "STUB: not implemented"; return nil }

// TODO(reddaly): Change this default.

// Query looks up a given service, in a domain, waiting at most
// for a timeout before finishing the query. The results are streamed
// to a channel. Sends will not block, so clients should make sure to
// either read or buffer.
func Query(params *QueryParam) error { _ = "STUB: not implemented"; return nil }

// QueryContext looks up a given service, in a domain, waiting at most
// for a timeout before finishing the query. The results are streamed
// to a channel. Sends will not block, so clients should make sure to
// either read or buffer. QueryContext will attempt to stop the query
// on cancellation.
func QueryContext(ctx context.Context, params *QueryParam) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a new client

// Set the multicast interface

// Ensure defaults are set

// Run the query

// Lookup is the same as Query, however it uses all the default parameters
func Lookup(service string, entries chan<- *ServiceEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// Client provides a query interface that can be used to
// search for service providers using mDNS
type client struct {
	use_ipv4 bool
	use_ipv6 bool

	ipv4UnicastConn *net.UDPConn
	ipv6UnicastConn *net.UDPConn

	ipv4MulticastConn *net.UDPConn
	ipv6MulticastConn *net.UDPConn

	closed   int32
	closedCh chan struct{} // TODO(reddaly): This doesn't appear to be used.

	log *log.Logger
}

// NewClient creates a new mdns Client that can be used to query
// for records
func newClient(v4 bool, v6 bool, logger *log.Logger) (*client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(reddaly): At least attempt to bind to the port required in the spec.
// Create a IPv4 listener

// Establish unicast connections

// Establish multicast connections

// Check that unicast and multicast connections have been made for IPv4 and IPv6
// and disable the respective protocol if not.

// Close is used to cleanup the client
func (c *client) Close() error { _ = "STUB: not implemented"; return nil }

// something else already closed it

// setInterface is used to set the query interface, uses system
// default if not provided
func (c *client) setInterface(iface *net.Interface) error { _ = "STUB: not implemented"; return nil }

// msgAddr carries the message and source address from recv to message processing.
type msgAddr struct {
	msg *dns.Msg
	src *net.UDPAddr
}

// query is used to perform a lookup and stream results
func (c *client) query(params *QueryParam) error {
	_ = "STUB: not implemented"
	// Create the service name
	return nil
}

// Start listening for response packets

// Send the query

// RFC 6762, section 18.12.  Repurposing of Top Bit of qclass in Question
// Section
//
// In the Question Section of a Multicast DNS query, the top bit of the qclass
// field is used to indicate that unicast responses are preferred for this
// particular question.  (See Section 5.4.)

// Map the in-progress responses

// Listen until we reach the timeout

// TODO(reddaly): Check that response corresponds to serviceAddr?

// Create new entry for this

// Check for a target mismatch

// Get the port

// Pull out the txt

// Pull out the IP

// @Deprecated

// Pull out the IP

// @Deprecated
// @Deprecated

// link-local IPv6 addresses must be qualified with a zone (interface). Zone is
// specific to this machine/network-namespace and so won't be carried in the
// mDNS message itself. We borrow the zone from the source address of the UDP
// packet, as the link-local address should be valid on that interface.

// Check if this entry is complete

// Fire off a node specific query

// sendQuery is used to multicast a query out
func (c *client) sendQuery(q *dns.Msg) error { _ = "STUB: not implemented"; return nil }

// recv is used to receive until we get a shutdown
func (c *client) recv(l *net.UDPConn, msgCh chan *msgAddr) { _ = "STUB: not implemented"; return }

// ensureName is used to ensure the named node is in progress
func ensureName(inprogress map[string]*ServiceEntry, name string) *ServiceEntry {
	_ = "STUB: not implemented"
	return nil
}

// alias is used to setup an alias between two entries
func alias(inprogress map[string]*ServiceEntry, src, dst string) { _ = "STUB: not implemented"; return }
