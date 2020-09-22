package impl

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/AlexsJones/vinculum/pkg/proto"
	"github.com/AlexsJones/vinculum/pkg/sys"
	"github.com/AlexsJones/vinculum/pkg/tracker"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/peer"
)

type NodeServerImpl struct{}

func init() {
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"169.254.0.0/16", // RFC3927 link-local
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	} {
		_, block, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(fmt.Errorf("parse error on %q: %v", cidr, err))
		}
		privateIPBlocks = append(privateIPBlocks, block)
	}
}

var privateIPBlocks []*net.IPNet

func isPrivateIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

func (NodeServerImpl) Register(ctx context.Context, syn *proto.ConnectionSyn) (*proto.ConnectionAck, error) {

	peer, bool := peer.FromContext(ctx)
	if !bool {
		return nil, errors.New("could not add node - could not resolve peer from context")
	}
	ip, _, err := net.SplitHostPort(peer.Addr.String())
	if err != nil {
		return nil, errors.New("could not resolve node IP")
	}

	if isPrivateIP(net.ParseIP(ip)) {
		log.Debug("Identified local client - Warning this is not recommended for more than one client")
		syn.Node.IpAddr = "localhost"
	} else {
		syn.Node.IpAddr = ip
	}
	syn.Node.Network = peer.Addr.Network()

	if tracker.Instance().Get(syn.Node.Guid) == nil {
		tracker.Instance().Add(syn.Node)
		log.Debugf("Added node %s @ %s", syn.Node.Guid, syn.Node.IpAddr)
	} else {
		log.Debugf("Received UpdateSyn from %s", syn.Node.Guid)
	}

	return &proto.ConnectionAck{
		Node:  sys.GetLocalNode(),
		State: proto.ConnectionAck_Known,
	}, nil
}
