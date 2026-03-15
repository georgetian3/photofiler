package ui

import (
	"log/slog"
	"photofiler/internal/sources"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/ncruces/zenity"
)

type FolderSelection struct {
	Path        string
	ButtonState widget.Clickable
	PathChan    chan string
	ErrChan     chan error
}

type ProcessingOptions struct {
	// Placeholder for future processing options
	IgnoreUnsupportedFiles bool
	ProcessInPlace         bool
}

type UiState struct {
	SelectedFolder    FolderSelection
	ProcessingOptions ProcessingOptions
}

func newUiState() *UiState {
	return &UiState{
		SelectedFolder: FolderSelection{
			PathChan: make(chan string),
			ErrChan:  make(chan error),
		},
	}
}

func (state *UiState) update(w *app.Window, gtx layout.Context) {

	if state.SelectedFolder.ButtonState.Clicked(gtx) {
		selectFolder(state.SelectedFolder.PathChan, state.SelectedFolder.ErrChan)
	}

	select {
	case path := <-state.SelectedFolder.PathChan:
		slog.Info("Folder selected", "path", path)
		state.SelectedFolder.Path = path
		go sources.ValidateSourceData(path)
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
