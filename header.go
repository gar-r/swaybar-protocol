package swaybarprotocol

// Header represents the swaybar header as defined in the swaybar
// json protocol. (see: man 7 swaybar-protocol)
type Header struct {
	Version     int  `json:"version"`
	ClickEvents bool `json:"click_events,omitempty"`
	ContSignal  int  `json:"cont_signal,omitempty"`
	StopSignal  int  `json:"stop_signal,omitempty"`
}
