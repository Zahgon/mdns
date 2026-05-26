// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MIT

package mdns

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

const (
	ipv4mdns              = "224.0.0.251"
	ipv6mdns              = "ff02::fb"
	mdnsPort              = 5353
	forceUnicastResponses = false
)

var (
	ipv4Addr = &net.UDPAddr{
		IP:   net.ParseIP(ipv4mdns),
		Port: mdnsPort,
	}
	ipv6Addr = &net.UDPAddr{
		IP:   net.ParseIP(ipv6mdns),
		Port: mdnsPort,
	}
)

// Config is used to configure the mDNS server
type Config struct {
	// Zone must be provided to support responding to queries
	Zone Zone

	// Iface if provided binds the multicast listener to the given
	// interface. If not provided, the system default multicase interface
	// is used.
	Iface *net.Interface

	// LogEmptyResponses indicates the server should print an informative message
	// when there is an mDNS query for which the server has no response.
	LogEmptyResponses bool

	// Logger can optionally be set to use an alternative logger instead of the default.
	Logger *log.Logger
}

// mDNS server is used to listen for mDNS queries and respond if we
// have a matching local record
type Server struct {
	config *Config

	ipv4List *net.UDPConn
	ipv6List *net.UDPConn

	shutdown   int32
	shutdownCh chan struct{}
}

// NewServer is used to create a new mDNS server from a config
func NewServer(config *Config) (*Server, error) {
	_ = "STUB: not implemented"
	// Create the listeners
	return nil, nil
}

// Check if we have any listener

// Shutdown is used to shutdown the listener
func (s *Server) Shutdown() error { _ = "STUB: not implemented"; return nil }

// something else already closed us

// recv is a long running routine to receive packets from an interface
func (s *Server) recv(c *net.UDPConn) { _ = "STUB: not implemented"; return }

// parsePacket is used to parse an incoming packet
func (s *Server) parsePacket(packet []byte, from net.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

// handleQuery is used to handle an incoming query
func (s *Server) handleQuery(query *dns.Msg, from net.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

// "In both multicast query and multicast response messages, the OPCODE MUST
// be zero on transmission (only standard queries are currently supported
// over multicast).  Multicast DNS messages received with an OPCODE other
// than zero MUST be silently ignored."  Note: OpcodeQuery == 0

// "In both multicast query and multicast response messages, the Response
// Code MUST be zero on transmission.  Multicast DNS messages received with
// non-zero Response Codes MUST be silently ignored."

// TODO(reddaly): Handle "TC (Truncated) Bit":
//    In query messages, if the TC bit is set, it means that additional
//    Known-Answer records may be following shortly.  A responder SHOULD
//    record this fact, and wait for those additional Known-Answer records,
//    before deciding whether to respond.  If the TC bit is clear, it means
//    that the querying host has no additional Known Answers.

// Handle each question

// See section 18 of RFC 6762 for rules about DNS headers.

// 18.1: ID (Query Identifier)
// 0 for multicast response, query.Id for unicast response

// 18.2: QR (Query/Response) Bit - must be set to 1 in response.

// 18.3: OPCODE - must be zero in response (OpcodeQuery == 0)

// 18.4: AA (Authoritative Answer) Bit - must be set to 1

// The following fields must all be set to 0:
// 18.5: TC (TRUNCATED) Bit
// 18.6: RD (Recursion Desired) Bit
// 18.7: RA (Recursion Available) Bit
// 18.8: Z (Zero) Bit
// 18.9: AD (Authentic Data) Bit
// 18.10: CD (Checking Disabled) Bit
// 18.11: RCODE (Response Code)

// 18.12 pertains to questions (handled by handleQuestion)
// 18.13 pertains to resource records (handled by handleQuestion)

// 18.14: Name Compression - responses should be compressed (though see
// caveats in the RFC), so set the Compress bit (part of the dns library
// API, not part of the DNS packet) to true.

// handleQuestion is used to handle an incoming question
//
// The response to a question may be transmitted over multicast, unicast, or
// both.  The return values are DNS records for each transmission type.
func (s *Server) handleQuestion(q dns.Question) (multicastRecs, unicastRecs []dns.RR) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle unicast and multicast responses.
// TODO(reddaly): The decision about sending over unicast vs. multicast is not
// yet fully compliant with RFC 6762.  For example, the unicast bit should be
// ignored if the records in question are close to TTL expiration.  For now,
// we just use the unicast bit to make the decision, as per the spec:
//     RFC 6762, section 18.12.  Repurposing of Top Bit of qclass in Question
//     Section
//
//     In the Question Section of a Multicast DNS query, the top bit of the
//     qclass field is used to indicate that unicast responses are preferred
//     for this particular question.  (See Section 5.4.)

// sendResponse is used to send a response packet
func (s *Server) sendResponse(resp *dns.Msg, from net.Addr, unicast bool) error {
	_ = "STUB: not implemented"
	// TODO(reddaly): Respect the unicast argument, and allow sending responses
	// over multicast.
	return nil
}

// Determine the socket to send from
