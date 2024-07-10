package flow

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/reassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	"sync"
)

type TCPStreamFactory struct {
	wg sync.WaitGroup
}

func (factory *TCPStreamFactory) New(netFlow, tcpFlow gopacket.Flow, tcp *layers.TCP, ac reassembly.AssemblerContext) reassembly.Stream {
	Debug("* New: %s %s\n", netFlow, tcpFlow)

	fsmOptions := reassembly.TCPSimpleFSMOptions{
		SupportMissingEstablishment: false, // 允许缺失 SYN、SYN+ACK、ACK
	}

	stream := &TCPStream{
		net:        netFlow,
		transport:  tcpFlow,
		reversed:   tcp.SrcPort == 80,
		tcpState:   reassembly.NewTCPSimpleFSM(fsmOptions),
		ident:      fmt.Sprintf("%s:%s", netFlow, tcpFlow),
		optChecker: reassembly.NewTCPOptionCheck(),
	}
	stream.client = Reader{
		readerStream: tcpreader.ReaderStream{},
		ident:        fmt.Sprintf("%s %s", netFlow, tcpFlow),
		parent:       stream,
		isClient:     true,
	}
	stream.server = Reader{
		readerStream: tcpreader.ReaderStream{},
		ident:        fmt.Sprintf("%s %s", netFlow.Reverse(), tcpFlow.Reverse()),
		parent:       stream,
		isClient:     false,
	}
}
