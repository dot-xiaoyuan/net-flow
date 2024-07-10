package flow

import "github.com/google/gopacket/tcpassembly/tcpreader"

type Reader struct {
	readerStream tcpreader.ReaderStream
	ident        string
	parent       *TCPStream
	isClient     bool
}
