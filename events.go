package swaybarprotocol

// ClickEvent represents an event that swaybar writes to the standard input
// when the user clicks on a block.
type ClickEvent struct {
	Name      string `json:"name"`
	Instance  string `json:"instance"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Button    int    `json:"button"`
	Event     int    `json:"event"`
	RelativeX int    `json:"relative_x"`
	RelativeY int    `json:"relative_y"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}
