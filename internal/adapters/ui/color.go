package ui

import (
	"github.com/fatih/color"
)

var (
	Blue    = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	Green   = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	Red     = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	Magenta = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()
	Cyan    = color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
)

func SetColor(nameFile string, styleColor color.Attribute) string {
	switch styleColor {
	case color.FgBlue:
		return Blue(nameFile)
	case color.FgGreen:
		return Green(nameFile)
	case color.FgRed:
		return Red(nameFile)
	case color.FgMagenta:
		return Magenta(nameFile)
	case color.FgCyan:
		return Cyan(nameFile)
	default:
		return nameFile
	}
}
