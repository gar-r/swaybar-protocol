package swaybarprotocol

// Header represents the swaybar header as defined in the swaybar
// json protocol. (see: man 7 swaybar-protocol)
type Header struct {
	Version     int  `json:"version"`
	ClickEvents bool `json:"click_events"`
	ContSignal  int  `json:"cont_signal"`
	StopSignal  int  `json:"stop_signal"`
}
