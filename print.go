package pr

import (
	"fmt"
	C "github.com/fatih/color"
)

var (
	enabled bool
)

func Enabled() bool {
	return enabled
}

func Disable() {
	enabled = false
}

func Enable() {
	enabled = true
}

func PrintBlack(format string, a ...interface{}) {
	base("black", format, a...)
}

func PrintRed(format string, a ...interface{}) {
	base("red", format, a...)
}

func PrintGreen(format string, a ...interface{}) {
	base("green", format, a...)
}

func PrintYellow(format string, a ...interface{}) {
	base("yellow", format, a...)
}

func PrintBlue(format string, a ...interface{}) {
	base("blue", format, a...)
}

func PrintMagenta(format string, a ...interface{}) {
	base("magenta", format, a...)
}

func PrintCyan(format string, a ...interface{}) {
	base("cyan", format, a...)
}

func PrintWhite(format string, a ...interface{}) {
	base("white", format, a...)
}

func SprintBlack(format string, a ...interface{}) string {
	return baseString("black", format, a...)
}

func SprintRed(format string, a ...interface{}) string {
	return baseString("red", format, a...)
}

func SprintGreen(format string, a ...interface{}) string {
	return baseString("green", format, a...)
}

func SprintYellow(format string, a ...interface{}) string {
	return baseString("yellow", format, a...)
}

func SprintBlue(format string, a ...interface{}) string {
	return baseString("blue", format, a...)
}

func SprintMagenta(format string, a ...interface{}) string {
	return baseString("magenta", format, a...)
}

func SprintCyan(format string, a ...interface{}) string {
	return baseString("cyan", format, a...)
}

func SprintWhite(format string, a ...interface{}) string {
	return baseString("white", format, a...)
}

func base(kind string, format string, a ...interface{}) {
	if !enabled {
		return
	}

	switch kind {
	case "black":
		C.Black(format, a...)
		break
	case "red":
		C.Red(format, a...)
		break
	case "green":
		C.Green(format, a...)
		break
	case "yellow":
		C.Yellow(format, a...)
		break
	case "blue":
		C.Blue(format, a...)
		break
	case "magenta":
		C.Magenta(format, a...)
		break
	case "cyan":
		C.Cyan(format, a...)
		break
	case "white":
		C.White(format, a...)
		break
	}
}

func Black(format string, a ...interface{}) {
	base("black", format, a...)
}

func Red(format string, a ...interface{}) {
	base("red", format, a...)
}

func Green(format string, a ...interface{}) {
	base("green", format, a...)
}

func Yellow(format string, a ...interface{}) {
	base("yellow", format, a...)
}

func Blue(format string, a ...interface{}) {
	base("blue", format, a...)
}

func Magenta(format string, a ...interface{}) {
	base("magenta", format, a...)
}

func Cyan(format string, a ...interface{}) {
	base("cyan", format, a...)
}

func White(format string, a ...interface{}) {
	base("white", format, a...)
}

func SBlack(format string, a ...interface{}) string {
	return baseString("black", format, a...)
}

func SRed(format string, a ...interface{}) string {
	return baseString("red", format, a...)
}

func SGreen(format string, a ...interface{}) string {
	return baseString("green", format, a...)
}

func SYellow(format string, a ...interface{}) string {
	return baseString("yellow", format, a...)
}

func SBlue(format string, a ...interface{}) string {
	return baseString("blue", format, a...)
}

func SMagenta(format string, a ...interface{}) string {
	return baseString("magenta", format, a...)
}

func SCyan(format string, a ...interface{}) string {
	return baseString("cyan", format, a...)
}

func SWhite(format string, a ...interface{}) string {
	return baseString("white", format, a...)
}

func baseString(kind string, format string, a ...interface{}) string {
	if !enabled {
		return ""
	}

	switch kind {
	case "black":
		return C.BlackString(format, a...)
	case "red":
		return C.RedString(format, a...)
	case "green":
		return C.GreenString(format, a...)
	case "yellow":
		return C.YellowString(format, a...)
	case "blue":
		return C.BlueString(format, a...)
	case "magenta":
		return C.MagentaString(format, a...)
	case "cyan":
		return C.CyanString(format, a...)
	case "white":
		return C.WhiteString(format, a...)
	}

	return ""
}


func Panic(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	panic(s)
}
func RedPanic(format string, a ...interface{}) {
	s := SRed(format, a...)
	panic(s)
}
