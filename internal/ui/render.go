package ui

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func renderFolderSelection(gtx layout.Context, th *material.Theme, state *UiState) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Body1(th, "Selected Folder").Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(4)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return widget.Border{
						Color:        th.Palette.Fg,
						CornerRadius: unit.Dp(4),
						Width:        unit.Dp(1),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return material.Editor(th, &state.SelectedFolder.PathEditor, "No folder selected").Layout(gtx)
						})
					})
				}),
				layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &state.SelectedFolder.Button, "Browse...").Layout(gtx)
				}),
			)
		}),
	)
}

func renderOptions(gtx layout.Context, th *material.Theme, state *UiState) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H6(th, "Processing Options").Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// When checked, the checkbox background becomes th.Palette.ContrastBg (the accent color),
			// which provides the "highlight" effect.
			cb := material.CheckBox(th, &state.ProcessingOptions.IgnoreUnsupportedFiles, "Ignore Unsupported Files")
			return cb.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			cb := material.CheckBox(th, &state.ProcessingOptions.ProcessInPlace, "Process In Place")
			return cb.Layout(gtx)
		}),
	)
}

func render(gtx layout.Context, th *material.Theme, state *UiState) {
	layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceStart}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return drawTime(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return renderFolderSelection(gtx, th, state)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return renderOptions(gtx, th, state)
			})
		}),
	)
}
