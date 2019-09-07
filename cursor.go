package ansi

import (
	"strconv"
)

func HideCursor() string            { return esc + "[?25l" }
func ShowCursor() string            { return esc + "[?25h" }
func CursorUp(n int) string         { return esc + "[" + strconv.Itoa(n) + "A" }
func CursorDown(n int) string       { return esc + "[" + strconv.Itoa(n) + "B" }
func CursorForward(n int) string    { return esc + "[" + strconv.Itoa(n) + "C" }
func CursorBackward(n int) string   { return esc + "[" + strconv.Itoa(n) + "D" }
func CursorStart(n int) string      { return esc + "[" + strconv.Itoa(n) + "G" }
func SaveCursorPosition() string    { return esc + "[s" }
func RestoreCursorPosition() string { return esc + "[u" }
func GetCursorPosition() string     { return esc + "[6n" }
