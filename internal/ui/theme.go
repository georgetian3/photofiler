package ui

import (
	"image/color"
	"log/slog"
	"os"
	"runtime"

	"gioui.org/widget/material"
)

func newWindowsTheme() *material.Theme {
	slog.Info("Initializing Windows theme")
	// 1. Initialize the standard Material theme with Go fonts
	th := material.NewTheme()

	// 2. Define Windows system colors (Light Mode)
	winBg := color.NRGBA{R: 243, G: 243, B: 243, A: 255}       // #F3F3F3
	winFg := color.NRGBA{R: 0, G: 0, B: 0, A: 255}             // #000000
	winAccent := color.NRGBA{R: 0, G: 120, B: 212, A: 255}     // #0078D4
	winAccentFg := color.NRGBA{R: 255, G: 255, B: 255, A: 255} // #FFFFFF

	// 3. Apply colors to the theme
	th.Palette.Bg = winBg
	th.Palette.Fg = winFg
	th.Palette.ContrastBg = winAccent
	th.Palette.ContrastFg = winAccentFg


	return th
}

func newMacOSTheme() *material.Theme {
	slog.Info("Initializing macOS theme")
	th := material.NewTheme()

	// macOS System Colors (Dark Mode)
	macBG := color.NRGBA{R: 28, G: 28, B: 30, A: 255}
	macAccentBlue := color.NRGBA{R: 10, G: 132, B: 255, A: 255}
	macText := color.NRGBA{R: 255, G: 255, B: 255, A: 240}

	// Apply to Theme
	th.Palette.Bg = macBG
	th.Palette.Fg = macText
	th.Palette.ContrastBg = macAccentBlue
	th.Palette.ContrastFg = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

	// macOS typically uses slightly smaller, tighter text for utility apps

	return th
}

func newTheme() *material.Theme {
	switch os.Getenv("OS") {
	case "windows":
		return newWindowsTheme()
	case "macos":
		return newMacOSTheme()
	}

	switch runtime.GOOS {
	case "darwin":
		return newMacOSTheme()
	case "windows":
		return newWindowsTheme()
	default:
	}
	return newWindowsTheme()
}
