package text

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
