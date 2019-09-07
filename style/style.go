package style

import (
	"strconv"
)

const (
	prefix = "\x1b["
	suffix = "m"
)

const (
	reset = 0
	// ANSI Styles
	bold            = 1
	dim             = 2 // Decreased intensity
	italic          = 3 // Not widely support, sometimes inverse.
	underline       = 4
	slowBlink       = 5 // Less than 150 times per minute
	fastBlink       = 6 // Over 150 times per minute
	inverse         = 7 // Swaps fg and bg colors
	conceal         = 8 // Not widely supported
	strikethrough   = 9 // Crossed-out
	defaultFont     = 10
	doubleUnderline = 21
	framed          = 51
	encircle        = 52
	overline        = 53
)

func Off(code int) int {
	switch {
	case code == 1:
		return 22
	case (code >= 2 && code <= 10):
		return code + 20
	default:
		return 0
	}
}

func Sequence(code int) string          { return prefix + strconv.Itoa(code) + suffix }
func Reset() string                     { return Sequence(reset) }
func Text(code int, text string) string { return Sequence(code) + text + Sequence(Off(code)) }

// Style Text
///////////////////////////////////////////////////////////////////////////////
func Strong(text string) string        { return Text(bold, text) }
func Bold(text string) string          { return Text(bold, text) }
func Italic(text string) string        { return Text(italic, text) }
func Emphasis(text string) string      { return Text(italic, text) }
func Dim(text string) string           { return Text(dim, text) }
func Thin(text string) string          { return Text(dim, text) }
func Underline(text string) string     { return Text(underline, text) }
func Strikethrough(text string) string { return Text(strikethrough, text) }
func Crossout(text string) string      { return Text(strikethrough, text) }
func Blink(text string) string         { return Text(slowBlink, text) }
func SlowBlink(text string) string     { return Text(slowBlink, text) }
func FastBlink(text string) string     { return Text(fastBlink, text) }
func RapidBlink(text string) string    { return Text(fastBlink, text) }
func Inverse(text string) string       { return Text(inverse, text) }
func Reverse(text string) string       { return Text(inverse, text) }
func Conceal(text string) string       { return Text(conceal, text) }
func Hide(text string) string          { return Text(conceal, text) }
func Framed(text string) string        { return Text(framed, text) }
func Encircle(text string) string      { return Text(encircle, text) }
func Overline(text string) string      { return Text(overline, text) }
