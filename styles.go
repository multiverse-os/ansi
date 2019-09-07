package text

type Style int

const (
	Bold          Style = 1
	Faint         Style = 2 // Decreased intensity
	Italic        Style = 3 // Not widely support, sometimes inverse.
	Underline     Style = 4
	SlowBlink     Style = 5 // Less than 150 times per minute
	RapidBlink    Style = 6 // Over 150 times per minute
	Reverse       Style = 7 // Swaps fg and bg colors
	Conceal       Style = 8 // Not widely supported
	Strikethrough Style = 9 // Crossed-out
)

const (
	Strong   = Bold
	Emphasis = Italic // These help with those familiar with HTML
	Dim      = Faint
	Thin     = Faint
	Inverse  = Reverse
	Blink    = SlowBlink
	Hide     = Conceal
	CrossOut = Strikethrough
)

func (s Style) On() attribute  { return Sequence(s.Code()) }
func (s Style) Off() attribute { return Sequence((s.Code() + 20)) }

func (s Style) String() string {
	switch s {
	case Bold:
		return "bold"
	case Faint:
		return "faint"
	case Italic:
		return "italic"
	case Underline:
		return "underline"
	case SlowBlink:
		return "blink"
	case FastBlink:
		return "fast blink"
	case Reverse:
		return "reverse"
	case Conceal:
		return "conceal"
	case Strikethrough:
		return "strikethrough"
	}
}
