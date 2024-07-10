package flow

import "github.com/google/gopacket"

// Assemble

type AssembleContext struct {
	CaptureInfo gopacket.CaptureInfo
}

func (a *AssembleContext) GetCaptureInfo() gopacket.CaptureInfo {
	return a.CaptureInfo
}
