package swaybarprotocol

// Body represents a single element of the infinite swaybar array, where each element
// represents a block to be rendered by swaybar. (see: man 7 swaybar-protocol)
type Body struct {
	Instance            string `json:"instance"`
	Border              string `json:"border"`
	Markup              string `json:"markup"`
	Background          string `json:"background"`
	FullText            string `json:"full_text"`
	ShortText           string `json:"short_text"`
	Name                string `json:"name"`
	Align               string `json:"align"`
	Color               string `json:"color"`
	BorderTop           int    `json:"border_top"`
	BorderLeft          int    `json:"border_left"`
	BorderBottom        int    `json:"border_bottom"`
	MinWidth            int    `json:"min_width"`
	SeparatorBlockWidth int    `json:"separator_block_width"`
	BorderRight         int    `json:"border_right"`
	Urgent              bool   `json:"urgent"`
	Separator           bool   `json:"separator"`
}

const (
	AlignLeft   = "left"
	AlignRight  = "right"
	AlignCenter = "center"

	MarkupPango = "pango"
	MarkupNone  = "none"
)
