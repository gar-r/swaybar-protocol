package swaybarprotocol

// Body represents a single element of the infinite swaybar array, where each element
// represents a block to be rendered by swaybar. (see: man 7 swaybar-protocol)
type Body struct {
	Instance            string `json:"instance,omitempty"`
	Border              string `json:"border,omitempty"`
	Markup              string `json:"markup,omitempty"`
	Background          string `json:"background,omitempty"`
	FullText            string `json:"full_text"`
	ShortText           string `json:"short_text,omitempty"`
	Name                string `json:"name"`
	Align               string `json:"align,omitempty"`
	Color               string `json:"color,omitempty"`
	BorderTop           int    `json:"border_top,omitempty"`
	BorderLeft          int    `json:"border_left,omitempty"`
	BorderBottom        int    `json:"border_bottom,omitempty"`
	MinWidth            int    `json:"min_width,omitempty"`
	SeparatorBlockWidth int    `json:"separator_block_width,omitempty"`
	BorderRight         int    `json:"border_right,omitempty"`
	Urgent              bool   `json:"urgent,omitempty"`
	Separator           bool   `json:"separator,omitempty"`
}

const (
	AlignLeft   = "left"
	AlignRight  = "right"
	AlignCenter = "center"

	MarkupPango = "pango"
	MarkupNone  = "none"
)
