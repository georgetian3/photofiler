package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gabriel-vasile/mimetype"
	"os"
	"path/filepath"
)

func mapMetadataToPhotos(inputDir string) {
	fmt.Println("entering", inputDir)
	defer fmt.Println("exiting", inputDir)
	entries, err := os.ReadDir(inputDir)
	if err != nil {
		fmt.Println("read dir:", err)
		return

	}

	for _, e := range entries {
		path := filepath.Join(inputDir, e.Name())
		if e.IsDir() {
			mapMetadataToPhotos(path)

			continue
		}

		stem := filepath.Base(e.Name())
		suffix := filepath.Ext(e.Name())
		mime, _ := mimetype.DetectFile(path)

		fmt.Println("Filename:", e.Name(), "Stem:", stem, "Suffix:", suffix, "type", mime)

		// TODO: extract metadata and map to photo
		// consider using goroutines and channels to parallelize
		// be sure to synchronize access to shared resources

	}
}



