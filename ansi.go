package ansi

//  8-bit
//
//  As 256-color lookup tables became common on graphic cards, escape sequences were added to select from a pre-defined set of 256 colors:[citation needed]
//
//    ESC[ 38;5;⟨n⟩ m Select foreground color
//    ESC[ 48;5;⟨n⟩ m Select background color
//      0-  7 :  standard colors (as in ESC [ 30–37 m)
//      8- 15 :  high intensity colors (as in ESC [ 90–97 m)
//      16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//    232-255 :  grayscale from black to white in 24 steps

const (
	esc    = "\u001B"
	prefix = "\x1b["
	suffix = "m"
)

const (
	NormalSequence = prefix + "0" + suffix
	ResetSequence  = prefix + "39" + suffix
)

const (
	Normal = 0
	Reset  = 39
)

func Sequence(code int) string { return prefix + strconv.Itoa(c.Code()) + suffix }

//
// Standardized Theme Color Functions
///////////////////////////////////////////////////////////////////////////////
func SystemDefault(text string) { return ResetSequence + text }
func Reset(text string)         { return SystemDefault(text) }
func Normal(text string)        { return SystemDefault(text) }

// Style Text
///////////////////////////////////////////////////////////////////////////////
func Strong(text string) string        { return (Strong.On() + text + Strong.Off()) }
func Bold(text string) string          { return (Bold.On() + text + Bold.Off()) }
func Italic(text string) string        { return (Italic.On() + text + Italic.Off()) }
func Emphasis(text string) string      { return (Emphasis.On() + text + Emphasis.Off()) }
func Faint(text string) string         { return (Faint.On() + text + Faint.Off()) }
func Light(text string) string         { return (Light.On() + text + Light.Off()) }
func Thin(text string) string          { return (Thin.On() + text + Thin.Off()) }
func Underline(text string) string     { return (Underline.On() + text + Underline.Off()) }
func Strikethrough(text string) string { return (Strikethrough.On() + text + Strikethrough.Off()) }
func Crossout(text string) string      { return (CrossOut.On() + text + CrossOut.Off()) }
func Blink(text string) string         { return (Blink.On() + text + Blink.Off()) }
func FastBlink(text string) string     { return (FastBlink.On() + text + FastBlink.Off()) }
func Inverse(text string) string       { return (Inverse.On() + text + Inverse.Off()) }
func Reverse(text string) string       { return (Reverse.On() + text + Reverse.Off()) }
func Conceal(text string) string       { return (Conceal.On() + text + Conceal.Off()) }
func Hide(text string) string          { return (Hide.On() + text + Hide.Off()) }

// Standard Color Functions
///////////////////////////////////////////////////////////////////////////////
func Black(text string) string        { return (Black.On() + text + Black.Off()) }
func Red(text string) string          { return (Red.On() + text + Red.Off()) }
func Green(text string) string        { return (Green.On() + text + Green.Off()) }
func Yellow(text string) string       { return (Yellow.On() + text + Yellow.Off()) }
func Blue(text string) string         { return (Blue.On() + text + Blue.Off()) }
func Magenta(text string) string      { return (Magenta.On() + text + Magenta.Off()) }
func Purple(text string) string       { return (Purple.On() + text + Magenta.Off()) }
func Cyan(text string) string         { return (Cyan.On() + text + Cyan.Off()) }
func Gray(text string) string         { return (Gray.On() + text + Gray.Off()) }
func DarkGray(text string) string     { return (DarkGray.On() + text + DarkGray.Off()) }
func LightRed(text string) string     { return (LightRed.On() + text + LightRed.Off()) }
func LightGreen(text string) string   { return (LightGreen.On() + text + LightGreen.Off()) }
func LightYellow(text string) string  { return (LightYellow.On() + text + LightYellow.Off()) }
func LightBlue(text string) string    { return (LightBlue.On() + text + LightBlue.Off()) }
func LightMagenta(text string) string { return (LightMagenta.On() + text + LightMagneta.Off()) }
func LightPurple(text string) string  { return (LightPurple.On() + text + LightPurple.Off()) }
func LightCyan(text string) string    { return (LightCyan.On() + text + LightCyan.Off()) }
func White(text string) string        { return (White.On() + text + White.Off()) }
