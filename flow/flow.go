package flow

import (
	"flag"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/reassembly"
	"os"
	"os/signal"
	"time"
)

const closeTimeout time.Duration = time.Hour * 24
const timeout time.Duration = time.Minute * 5

var (
	pcapFile = flag.String("pcap", "", "PCAP file to use")
	debug    = flag.Bool("debug", false, "Enable debug mode")
)

func main() {
	var handle *pcap.Handle
	var err error
	var dec gopacket.Decoder
	var count = 0

	if handle, err = pcap.OpenOffline(*pcapFile); err != nil {
		fmt.Printf("Failed to open pcap file %s: %s\n", *pcapFile, err)
	}

	source := gopacket.NewPacketSource(handle, dec)

	factory := &TCPStreamFactory{}
	streamPool := reassembly.NewStreamPool(factory)
	streamAssembler := reassembly.NewAssembler(streamPool)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	for packet := range source.Packets() {
		count++
		tcp := packet.Layer(layers.LayerTypeTCP)
		if tcp != nil {
			tcpLayer := tcp.(*layers.TCP)
			c := &AssembleContext{
				packet.Metadata().CaptureInfo,
			}
			streamAssembler.AssembleWithContext(packet.NetworkLayer().NetworkFlow(), tcpLayer, c)
		}
		if count%1000 == 0 {
			ref := packet.Metadata().Timestamp
			flushed, closed := streamAssembler.FlushWithOptions(reassembly.FlushOptions{T: ref.Add(-timeout), TC: ref.Add(-closeTimeout)})
			Debug("Flushed %d packets, closed %d packets", flushed, closed)
		}
		done := false
		select {
		case <-signalChan:
			fmt.Fprintf(os.Stderr, "\nReceived an interrupt, stopping capture...\n")
			done = true
		default:

		}
		if done {
			break
		}
	}

	closed := streamAssembler.FlushAll()
	Debug("Final flush: %d closed", closed)
}
