package swaybarprotocol

import (
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	h := &Header{
		Version:     1,
		ClickEvents: true,
		ContSignal:  18,
		StopSignal:  19,
	}

	w := &strings.Builder{}

	Init(w, h)

	json := `{"version":1,"click_events":true,"cont_signal":18,"stop_signal":19}`
	expected := json + "\n" + "["
	actual := w.String()

	if actual != expected {
		t.Errorf("expected: %s, got: %s", expected, actual)
	}
}

func TestOutput(t *testing.T) {
	arr := []*Body{
		{
			FullText:            "fullText",
			ShortText:           "shortText",
			Color:               "#ccccccff",
			Background:          "#111111ff",
			Border:              "#222222ff",
			BorderTop:           1,
			BorderBottom:        1,
			BorderLeft:          1,
			BorderRight:         1,
			MinWidth:            100,
			Align:               AlignCenter,
			Name:                "name",
			Instance:            "instance",
			Urgent:              true,
			Separator:           true,
			SeparatorBlockWidth: 5,
			Markup:              MarkupNone,
		},
	}

	w := &strings.Builder{}

	Output(w, arr)

	json := `[{"instance":"instance","border":"#222222ff","markup":"none","background":"#111111ff","full_text":"fullText","short_text":"shortText","name":"name","align":"center","color":"#ccccccff","border_top":1,"border_left":1,"border_bottom":1,"min_width":100,"separator_block_width":5,"border_right":1,"urgent":true,"separator":true}]`
	expected := json + "\n" + ","
	actual := w.String()

	if actual != expected {
		t.Errorf("expected %s, got: %s", expected, actual)
	}
}
