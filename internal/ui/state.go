package ui

import (
	"log/slog"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/ncruces/zenity"
)

type FolderSelection struct {
	Path       string
	PathEditor widget.Editor
	Button     widget.Clickable
	PathChan   chan string
	ErrChan    chan error
}

type ProcessingOptions struct {
	// Placeholder for future processing options
	IgnoreUnsupportedFiles widget.Bool
	ProcessInPlace         widget.Bool
}

type UiState struct {
	SelectedFolder    FolderSelection
	ProcessingOptions ProcessingOptions
}

func newUiState() *UiState {
	state := &UiState{
		SelectedFolder: FolderSelection{
			PathChan: make(chan string),
			ErrChan:  make(chan error),
		},
	}
	state.SelectedFolder.PathEditor.SingleLine = true
	state.SelectedFolder.PathEditor.ReadOnly = true
	return state
}

func (state *UiState) update(w *app.Window, gtx layout.Context) {

	if state.SelectedFolder.Button.Clicked(gtx) {
		selectFolder(state.SelectedFolder.PathChan, state.SelectedFolder.ErrChan)
	}

	select {
	case path := <-state.SelectedFolder.PathChan:
		slog.Info("Folder selected", "path", path)
		state.SelectedFolder.Path = path
		state.SelectedFolder.PathEditor.SetText(path)
		// go sources.ValidateSourceData(path)
		w.Invalidate()
	case err := <-state.SelectedFolder.ErrChan:
		if err == zenity.ErrCanceled {
			slog.Info("Folder selection canceled by user")
			return
		}
		slog.Error("Error selecting folder", "error", err)
		w.Invalidate()
	default:
	}
}
