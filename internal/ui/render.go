package ui

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func render(gtx layout.Context, th *material.Theme, state *UiState) {
	material.H2(th, "Selected Folder: "+state.SelectedFolder.Path).Layout(gtx)
	drawTime(gtx, th)

	layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return material.Button(th, &state.SelectedFolder.ButtonState, "Choose Folder").Layout(gtx)
	})
}
