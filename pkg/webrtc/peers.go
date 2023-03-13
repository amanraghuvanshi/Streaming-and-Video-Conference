package webrtc

import (
	"sync"

	"github.com/amanraghuvanshi/videostreaming/pkg/chat"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}
type Peers struct {
	Listlock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

func (p *Peers) DispatchKeyFrames() {

}
