package flow

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/reassembly"
	"sync"
)

type TCPStream struct {
	sync.Mutex
	client         Reader
	server         Reader
	tcpState       *reassembly.TCPSimpleFSM
	optChecker     reassembly.TCPOptionCheck
	net, transport gopacket.Flow
	reversed       bool
	fsmErr         bool
	urls           []string
	ident          string
}

func (s *TCPStream) Accept(tcp *layers.TCP, ci gopacket.CaptureInfo, dir reassembly.TCPFlowDirection, nextSeq reassembly.Sequence, start *bool, ac reassembly.AssemblerContext) bool {
	//TODO implement me
	panic("implement me")
}

func (s *TCPStream) ReassembledSG(sg reassembly.ScatterGather, ac reassembly.AssemblerContext) {
	//TODO implement me
	panic("implement me")
}

func (s *TCPStream) ReassemblyComplete(ac reassembly.AssemblerContext) bool {
	//TODO implement me
	panic("implement me")
}
