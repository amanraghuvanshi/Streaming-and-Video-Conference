package webrtc

import "sync"

type Room struct {
	Peers
}
type Peers struct {
	Listlock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

func (p *Peers) DispatchKeyFrames() {

}
