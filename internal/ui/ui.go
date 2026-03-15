package ui

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func Run() {
	w := new(app.Window)
	w.Option(app.Title("Photofiler"))
	var ops op.Ops
	uiState := newUiState()
	th := material.NewTheme()

	go func() {
		if err := func() error {
			for {
				switch e := w.Event().(type) {
				case app.DestroyEvent:
					return e.Err
				case app.FrameEvent:
					gtx := app.NewContext(&ops, e)
					uiState.update(w, gtx)
					render(gtx, th, uiState)
					e.Frame(gtx.Ops)
				}
			}
		}(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}
