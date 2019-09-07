package text

// Color represents one of the standard 16 ANSI colors.
type Color int

const (
	ForegroundOff Color = 39
	BackgroundOff Color = 49
)

const (
	Black         Color = 30
	Red           Color = 31
	Green         Color = 32
	Yellow        Color = 33
	Blue          Color = 34
	Magenta       Color = 35
	Cyan          Color = 36
	Gray          Color = 37
	DarkGray      Color = 90
	BrightRed     Color = 91
	BrightGreen   Color = 92
	BrightYellow  Color = 93
	BrightBlue    Color = 94
	BrightMagenta Color = 95
	BrightCyan    Color = 96
	White         Color = 97
)

const (
	ColorOff = ForegroundOff
	FgOff    = ForegroundOff
	BgOff    = BackgroundOff
)

func (c Color) Format(text string) string { return (c.On() + text + c.Off()) }

func (c Color) On() string  { return FormatSequence(c.Code()) }
func (c Color) Off() string { return FormatSequence(ColorOff) }

func (c Color) ForegroundCode() string { return int(c) }
func (c Color) BackgroundCode() string { return (Code(c) + 10) }

func (c Color) FgCode() string { return c.ForegroundCode() }
func (c Color) BgCode() string { return c.BackgroundCode() }

func (c Color) String() string {
	switch c {
	case Black:
		return "black"
	case Red:
		return "red"
	case Green:
		return "green"
	case Yellow:
		return "yellow"
	case Blue:
		return "blue"
	case Magenta:
		return "magenta"
	case Cyan:
		return "cyan"
	case Gray:
		return "gray"
	case DarkGray:
		return "dark gray"
	case BrightRed:
		return "bright red"
	case BrightGreen:
		return "bright green"
	case BrightYellow:
		return "bright yellow"
	case BrightBlue:
		return "bright blue"
	case BrightMagenta:
		return "bright magenta"
	case BrightCyan:
		return "bright cyan"
	case White:
		return "white"
	}
}
