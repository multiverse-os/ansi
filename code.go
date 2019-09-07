package main

type Code int

func (c Code) Int() string { return int(c) }

const (
	Reset Code = 0 // Normal
	// ANSI Styles
	Bold          Code = 1
	Faint         Code = 2 // Decreased intensity
	Italic        Code = 3 // Not widely support, sometimes inverse.
	Underline     Code = 4
	SlowBlink     Code = 5 // Less than 150 times per minute
	RapidBlink    Code = 6 // Over 150 times per minute
	Reverse       Code = 7 // Swaps fg and bg colors
	Conceal       Code = 8 // Not widely supported
	Strikethrough Code = 9 // Crossed-out
	// Primary 8 ANSI Colors
	Black   Code = 30
	Red     Code = 31
	Green   Code = 32
	Yellow  Code = 33
	Blue    Code = 34
	Magenta Code = 35
	Cyan    Code = 36
	Gray    Code = 37
	// Primary 8 ANSI Background
	BlackBackground   Code = 40
	RedBackground     Code = 41
	GreenBackground   Code = 42
	YellowBackground  Code = 43
	BlueBackground    Code = 44
	MagentaBackground Code = 45
	CyanBackground    Code = 46
	GrayBackground    Code = 47
	// Off Codes
	ColorOff      Code = 39
	BackgroundOff Code = 49
	// Secondary 8 ANSI Colors
	DarkGray      Code = 90
	BrightRed     Code = 91
	BrightGreen   Code = 92
	BrightYellow  Code = 93
	BrightBlue    Code = 94
	BrightMagenta Code = 95
	BrightCyan    Code = 96
	White         Code = 97
	// Secondary 8 ANSI Background
	DarkGrayBackground      Code = 100
	BrightRedBackground     Code = 101
	BrightGreenBackground   Code = 102
	BrightYellowBackground  Code = 103
	BrightBlueBackground    Code = 104
	BrightMagentaBackground Code = 105
	BrightCyanBackground    Code = 106
	WhiteBackground         Code = 107
)

// Aliasing
const (
	// Reset Aliasing
	Normal = Reset
	// Color Aliasing
	Purple = Magenta
	// Style Aliasing
	Strong   = Bold
	Emphasis = Italic // These help with those familiar with HTML
	Dim      = Faint
	Thin     = Faint
	Inverse  = Reverse
	Blink    = SlowBlink
	Hide     = Conceal
	CrossOut = Strikethrough
)

func Reset(text string) string         { return Reset.Sequence() + text }
func ForegroundOff(text string) string { return ForegroundOff.Sequence() + text }
func BackgroundOff(text string) string { return BackgroundOff.Sequence() + text }
func StyleOff(text string) string      { return StyleOff.Sequence() + text }

func Foreground(code Code, text string) string {
	return Sequence(code.Int()) + text + Sequence(ColorOff.Int())
}
func Background(code Code, text string) string {
	return Sequence(code.Int()+10) + text + Sequence(ColorOff.Int()+10)
}
func Style(code Code, text string) string {
	return Sequence(code.Int()) + text + Sequence(code.Int()+20)
}

func (c Code) Color(text string) string {
	switch code := c.Int(); {
	case Reset.Int():
		return Reset(text)
	case (code >= 30 && code <= 37), (code >= 90 && code <= 97):
		return Color(c, text)
	case ForegroundOff.Int():
		return ForegroundOff(text)
	case BackgroundOff.Int():
		return BackgroundOff(text)
	case (code >= 40 && code <= 47), (code >= 100 && code <= 107):
		return Background(c, text)
	default:
		return text
	}
}
func (c Code) Format(text string) string {
	switch code := c.Int(); {
	case (code >= 1 && code <= 9):
		Style(c, text)
	case StyleOff.Int():
		StyleOff(text)
	default:
		return c.Color(text)
	}
}
