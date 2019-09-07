package style

const (
	escape = "\x1b"
	prefix = "["
	suffix = "m"
)

const (
	Normal        = 0
	Reset         = 0
	Bold          = 1
	Faint         = 2 // Decreased intensity
	Italic        = 3 // Not widely support, sometimes inverse.
	Underline     = 4
	SlowBlink     = 5 // Less than 150 times per minute
	RapidBlink    = 6 // Over 150 times per minute
	Reverse       = 7 // Swaps fg and bg colors
	Conceal       = 8 // Not widely supported
	Strikethrough = 9 // Crossed-out
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

func Sequence(code int) string { return escape + prefix + strconv.Itoa(code) + suffix }

func On() attribute  { return Sequence(s.Code()) }
func Off() attribute { return Sequence((s.Code() + 20)) }
