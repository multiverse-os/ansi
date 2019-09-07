package text

// Style Text
///////////////////////////////////////////////////////////////////////////////
func Strong(text string) string        { return Style(Strong, text) }
func Bold(text string) string          { return Style(Bold, text) }
func Italic(text string) string        { return Style(Italic, text) }
func Emphasis(text string) string      { return Style(Emphasis, text) }
func Faint(text string) string         { return Style(Faint, text) }
func Light(text string) string         { return Style(Light, text) }
func Thin(text string) string          { return Style(Thin, text) }
func Underline(text string) string     { return Style(Underline, text) }
func Strikethrough(text string) string { return Style(Strikethrough, text) }
func Crossout(text string) string      { return Style(CrossOut, text) }
func Blink(text string) string         { return Style(Blink, text) }
func FastBlink(text string) string     { return Style(FastBlink, text) }
func Inverse(text string) string       { return Style(Inverse, text) }
func Reverse(text string) string       { return Style(Reverse, text) }
func Conceal(text string) string       { return Style(Conceal, text) }
func Hide(text string) string          { return Style(Hide, text) }
