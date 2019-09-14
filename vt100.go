package ansi

import (
	"fmt"
	"strings"
)

// Attribute constants for use with SetAttrs.
const (
	ATTR_RESET = 0

	//ATTR_BRIGHT = 1
	//ATTR_BOLD = 1  // alias for BRIGHT
	//ATTR_DIM = 2
	//ATTR_FAINT = 2  // alias for DIM
	//ATTR_ITALIC = 3  // not widely supported
	//ATTR_UNDERLINE = 4
	//ATTR_BLINK = 5
	//ATTR_BLINK_FAST = 6  // not widely supported
	//ATTR_REVERSE = 7
	//ATTR_CONCEAL = 8
	//ATTR_STRIKETHROUGH = 9  // not widely supported

	ATTR_FONT_STD   = 10
	ATTR_FONT_ALT_1 = 11
	ATTR_FONT_ALT_2 = 12
	ATTR_FONT_ALT_3 = 13
	ATTR_FONT_ALT_4 = 14
	ATTR_FONT_ALT_5 = 15
	ATTR_FONT_ALT_6 = 16
	ATTR_FONT_ALT_7 = 17
	ATTR_FONT_ALT_8 = 18
	ATTR_FONT_ALT_9 = 19

	ATTR_NORMAL = 22 // cancels out ATTR_BOLD/ATTR_BRIGHT and ATTR_FAINT/ATTR_DIM
	//ATTR_NO_ITALIC = 23  // cancels out ATTR_ITALIC
	//ATTR_NO_UNDERLINE = 24  // cancels out ATTR_UNDERLINE
	//ATTR_NO_BLINK = 25  // cancels out ATTR_BLINK and ATTR_BLINK_FAST
	//ATTR_NO_CONCEAL = 28  // cancels out ATTR_CONCEAL
	//ATTR_NO_STRIKETHROUGH = 29  // cancels out ATTR_STRIKETHROUGH

	//ATTR_FG_BLACK = 30
	//ATTR_FG_RED = 31
	//ATTR_FG_GREEN = 32
	//ATTR_FG_YELLOW = 33
	//ATTR_FG_BLUE = 34
	//ATTR_FG_MAGENTA = 35
	//ATTR_FG_CYAN = 36
	//ATTR_FG_WHITE = 37
	//ATTR_FG_256 = 38  // used in setting foreground in 256-colour terminals, should not be used directly (use SetFG256(n) instead)
	ATTR_FG_RESET = 39 // reset foreground to default

	//ATTR_BG_BLACK = 40
	//ATTR_BG_RED = 41
	//ATTR_BG_GREEN = 42
	//ATTR_BG_YELLOW = 43
	//ATTR_BG_BLUE = 44
	//ATTR_BG_MAGENTA = 45
	//ATTR_BG_CYAN = 46
	//ATTR_BG_WHITE = 47
	//ATTR_BG_256 = 38  // used in setting background in 256-colour terminals, should not be used directly (use SetBG256(n) instead)
	ATTR_BG_RESET = 39 // reset background to default
)

func emitEscape(baseCode string, args ...interface{}) {
	strArgs := make([]string, len(args))
	for i, arg := range args {
		strArgs[i] = fmt.Sprint(arg)
	}
	fmt.Print("\x1B[" + strings.Join(strArgs, ";") + baseCode)
}

// Move the cursor up r rows.
func CursorUp(r uint) {
	emitEscape("A", r)
}

// Move the cursor down r rows.
func CursorDown(r uint) {
	emitEscape("B", r)
}

// Move the cursor forward c columns.
func CursorForward(c uint) {
	emitEscape("C", c)
}

// Move the cursor back c columns.
func CursorBackward(c uint) {
	emitEscape("D", c)
}

// Move the cursor down r rows, and to the leftmost column.
func CursorNextLine(r uint) {
	emitEscape("E", r)
}

// Move the cursor up r rows, and to the leftmost column.
func CursorPrevLine(r uint) {
	emitEscape("F", r)
}

// Move the cursor to the column c, while remaining on the current row.
func CursorToColumn(row, col uint) {
	emitEscape("G", row, col)
}

// Move the cursor to the cell at [row, col].
func CursorTo(row, col uint) {
	emitEscape("H", row, col)
}

// Move the cursor to the home position.
func CursorHome() {
	emitEscape("H")
}

// Clear everything from the current character to the end of the screen (moving along columns, then rows).
func ClearScreenPartialForward() {
	emitEscape("J")
}

// Clear everything from the current character to the start of the screen (moving along columns, then rows).
func ClearScreenPartialBackward() {
	emitEscape("J", 1)
}

// Clear the entire screen. May also return the cursor to the home position, but this is not universal, so do not rely on it.
func ClearScreen() {
	emitEscape("J", 2)
}

// Clear everything from the current character to the end of the line.
func ClearLinePartialForward() {
	emitEscape("K")
}

// Clear everything from the current character to the start of the line.
func ClearLinePartialBackward() {
	emitEscape("K", 1)
}

// Clear the current line.
func ClearLine() {
	emitEscape("K", 2)
}

// Set one or more attributes. They will be set in order, so if any attribute conflicts with another, the last one will be the one which is set when the function returns.
// This allows you to, for example, use ATTR_RESET as the first attribute to ensure only the other attributes supplied will be active.
func SetAttrs(attrs ...uint) {
	attrsInterfaced := make([]interface{}, len(attrs))
	for i, attr := range attrs {
		attrsInterfaced[i] = attr
	}
	emitEscape("m", attrsInterfaced...)
}

// Reset all attributes to default.
func ResetAttrs() {
	emitEscape("m")
}

// Set a foreground colour by number (n) from the 256-colour set. This is generally only supported on xterm and related terminals.
func SetFG256(n uint) {
	SetAttrs(ATTR_FG_256, 5, n)
}

// Set a background colour by number (n) from the 256-colour set. This is generally only supported on xterm and related terminals.
func SetBG256(n uint) {
	SetAttrs(ATTR_BG_256, 5, n)
}

// Ask the terminal for the cursor position: will be returned in the form \x1B[{row};{col}R, where {row} and {col} are replaced with the current row/column. Handling this sequence is not a part of package vt100.
func ReportCursorPos() {
	emitEscape("n", 6)
}

// Save the cursor position, to be restored by a later RestoreCursorPos()
func SaveCursorPos() {
	emitEscape("s")
}

// Restore the cursor position to the position saved by an earlier SaveCursorPos()
func RestoreCursorPos() {
	emitEscape("u")
}

// Set cursor visibility.
func SetCursorVisible(visible bool) {
	if visible {
		emitEscape("?25h")
	} else {
		emitEscape("?25l")
	}
}

// Clear the screen, reset all attributes, and return the cursor to home position.
func ResetScreen() {
	ClearScreen()
	ResetAttrs()
	CursorHome()
}

// Print some text at a given [row, col].
func PrintAt(row, col uint, text string) {
	CursorTo(row, col)
	fmt.Print(text)
}
