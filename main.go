package main

import (
	"photofiler/internal/sources"
)




func main() {

	// source := "google"
	// var structure interface{}
	// switch source {
	// case "google":
	// 	structure = &sources.GoogleTakeoutMetadata{}
	// case "apple":
	// 	structure = &sources.AppleMetadata{}
	// default:
	// 	fmt.Println("unsupported source:", source)
	// 	return
	// }
	// fmt.Println("Selected structure: %v", structure)
	// counter := 0
	// maxIter := 20

	rootDir := "D:/photos takeout"
	sources.ValidateSourceData(rootDir)
	// err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {


}
