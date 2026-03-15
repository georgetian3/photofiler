package sources

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/schollz/progressbar/v3"
)

func countFiles(rootDir string) (int, error) {
	counter := 0
	err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("Error walking dir: %w", err)
		}
		if !d.IsDir() {
			counter++
		}
		return nil
	})
	return counter, err
}

func getFilenames(rootDir string) ([]string, error) {
	var fileNames []string
	err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("Error walking dir: %w", err)
		}
		if !d.IsDir() {
			fileNames = append(fileNames, path)
		}
		return nil
	})
	return fileNames, err
}

func getMimetypes(files []string, bar *progressbar.ProgressBar) (map[string][]string, error) {
	mimetypes := make(map[string][]string)
	for _, file := range files {
		mime, err := mimetype.DetectFile(file)
		if err != nil {
			return nil, err
		}
		mimetypes[mime.String()] = append(mimetypes[mime.String()], file)
		if bar != nil {
			bar.Add(1)
		}
	}
	bar.Finish()
	return mimetypes, nil
}

var ALLOWED_MIMETYPES = map[string]struct{}{
	"image/jpeg":       {},
	"image/png":        {},
	"image/heic":       {},
	"image/heif":       {},
	"image/gif":        {},
	"video/mp4":        {},
	"video/quicktime":  {},
	"application/json": {},
}

func isSupportedMimetype(mime string) bool {
	return mime == "application/json" || strings.Contains(mime, "image/") || strings.Contains(mime, "video/")
}

func ValidateSourceData(rootDir string) {
	slog.Info(fmt.Sprintf("Counting files in directory: %s", rootDir))
	filenames, err := getFilenames(rootDir)
	if err != nil {
		slog.Error("Error counting files:", slog.Any("error", err))
		return
	}
	slog.Info(fmt.Sprintf("File count: %d", len(filenames)))
	bar := progressbar.Default(int64(len(filenames)))
	mimetypes, err := getMimetypes(filenames, bar)
	if err != nil {
		slog.Error("Error counting mimetypes:", slog.Any("error", err))
		return
	}
	keys := make([]string, 0, len(mimetypes))
	for k := range mimetypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	unsupportedMimetypes := make([]string, 0)
	for _, mime := range keys {
		files := mimetypes[mime]
		if isSupportedMimetype(mime) {
			slog.Info(fmt.Sprintf("  Supported mimetype: %-20scount: %d", mime, len(files)))
			} else {
			unsupportedMimetypes = append(unsupportedMimetypes, files...)
			slog.Warn(fmt.Sprintf("Unsupported mimetype: %-20scount: %d", mime, len(files)))
		}
	}
	unsupportedMimetypesFilename := "unsupported_mimetypes.txt"
	if len(unsupportedMimetypes) > 0 {
		slog.Warn("##############################################################")
		slog.Warn("Source contains files with unsupported mimetypes")
		slog.Warn("These files have been written to " + unsupportedMimetypesFilename)
		slog.Warn("Consider filtering or converting these files before processing")
		slog.Warn("##############################################################")

	}
	os.WriteFile(unsupportedMimetypesFilename, []byte(strings.Join(unsupportedMimetypes, "\n")), 0644)

}
