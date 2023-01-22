package main

import (
	"fmt"

	ansi "github.com/multiverse-os/ansi"
)

func main() {
	// TODO: Fix the order
	// Primary 8 ANSI Colors
	fmt.Printf(" +=====================+\n")
	fmt.Printf(" | Primary ANSI Colors |\n")
	fmt.Printf(" +=====================+\n")
	fmt.Printf(" | %s   | %s   |  \n", ansi.Black("Black"), ansi.BlackBg("BlackBg"))
	fmt.Printf(" | %s  | %s  | \n", ansi.Maroon("Maroon"), ansi.MaroonBg("MaroonBg"))
	fmt.Printf(" | %s   | %s   |  \n", ansi.Green("Green"), ansi.GreenBg("GreenBg"))
	fmt.Printf(" | %s   | %s   | \n", ansi.Olive("Olive"), ansi.OliveBg("OliveBg"))
	fmt.Printf(" | %s    | %s    |  \n", ansi.Blue("Blue"), ansi.BlueBg("BlueBg"))
	fmt.Printf(" | %s | %s |  \n", ansi.Magenta("Magenta"), ansi.MagentaBg("MagentaBg"))
	fmt.Printf(" | %s    | %s    |  \n", ansi.Cyan("Cyan"), ansi.CyanBg("CyanBg"))
	fmt.Printf(" | %s  | %s  | \n", ansi.Silver("Silver"), ansi.SilverBg("SilverBg"))
	fmt.Printf(" +---------+-----------+\n\n")

	// Secondary 8 ANSI Colors
	fmt.Printf(" +=====================+\n")
	fmt.Printf(" |Secondary ANSI Colors|\n")
	fmt.Printf(" +=====================+\n")
	fmt.Printf(" | %s    | %s    | \n", ansi.Gray("Gray"), ansi.GrayBg("GrayBg"))
	fmt.Printf(" | %s     | %s     |  \n", ansi.Red("Red"), ansi.RedBg("RedBg"))
	fmt.Printf(" | %s    | %s    | \n", ansi.Lime("Lime"), ansi.LimeBg("LimeBg"))
	fmt.Printf(" | %s  | %s  |  \n", ansi.Yellow("Yellow"), ansi.YellowBg("YellowBg"))
	fmt.Printf(" | %s | %s | \n", ansi.SkyBlue("SkyBlue"), ansi.SkyBlueBg("SkyBlueBg"))
	fmt.Printf(" | %s | %s | \n", ansi.Fuchsia("Fuchsia"), ansi.FuchsiaBg("FuchsiaBg"))
	fmt.Printf(" | %s    | %s    | \n", ansi.Aqua("Aqua"), ansi.AquaBg("AquaBg"))
	fmt.Printf(" | %s   | %s   |  \n", ansi.White("White"), ansi.WhiteBg("WhiteBg"))
	fmt.Printf(" +---------+-----------+\n\n")

	fmt.Printf(" ANSI Style Options\n")
	fmt.Printf(" +=====================+\n")
	fmt.Printf("  %s  \n", ansi.Bold("Bold"))
	fmt.Printf("  %s  \n", ansi.Dim("Dim"))
	fmt.Printf("  %s  \n", ansi.Italic("Italic"))
	fmt.Printf("  %s  \n", ansi.Underline("Underline"))
	fmt.Printf("  %s  \n", ansi.SlowBlink("Slow Blink"))
	fmt.Printf("  %s  \n", ansi.FastBlink("Fast Blink"))
	fmt.Printf("  %s  \n", ansi.Inverse("Inverse"))
	fmt.Printf("  %s  \n", ansi.Conceal("Conceal"))
	fmt.Printf("  %s  \n", ansi.Strikethrough("Strikethough"))
	fmt.Printf("  %s  \n", ansi.Framed("Framed"))
	fmt.Printf("  %s  \n", ansi.Encircle("Encircle"))
	fmt.Printf("  %s  \n", ansi.Overline("Overline"))
}
