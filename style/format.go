package style

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
