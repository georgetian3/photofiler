package ui

import (
	"log/slog"
	"time"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/ncruces/zenity"
)

func drawTime(gtx layout.Context, th *material.Theme) layout.Dimensions {
	// 1. Get the current time
	formattedTime := time.Now().Format("15:04:05")

	// 2. Schedule the NEXT redraw in 1 second
	// Using gtx.Execute is the modern way to run "commands" in Gio
	gtx.Execute(op.InvalidateCmd{At: gtx.Now.Add(time.Millisecond * 100)})

	// 3. Display the time
	return material.H1(th, formattedTime).Layout(gtx)
}

func selectFolder(folderChan chan<- string, errChan chan<- error) {
	go func() {
		slog.Info("Opening file selection dialog")
		path, err := zenity.SelectFile(zenity.Directory())
		if err != nil {
			errChan <- err
		} else {
			folderChan <- path
		}
		slog.Info("File selection dialog closed")
	}()
}
