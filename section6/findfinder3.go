package section6

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Find and write the names of subdirectories to a file
//
//  Create a program that can get multiple directory paths from
//  the command-line, and prints only their subdirectories into a
//  file named: dirs.txt
//
//
//  1. Get the directory paths from command-line
//
//  2. Append the names of subdirectories inside each directory
//     to a byte slice
//
//  3. Write that byte slice to dirs.txt file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Please provide directory paths
//
//   go run main.go dir/ dir2/
//
//   cat dirs.txt
//
//     dir/
//             subdir1/
//             subdir2/
//
//     dir2/
//             subdir1/
//             subdir2/
//             subdir3/
//
//
// HINTS
//
//   ONLY READ THIS IF YOU GET STUCK
//
//   + Get all the files in a directory using ioutil.ReadDir
//     (A directory is also a file)
//
//   + You can use IsDir method of a FileInfo value to detect
//     whether a file is a directory or not.
//
//     Check out its documentation:
//
//     go doc os.FileInfo.IsDir
//
//     # or using godocc
//     godocc os.FileInfo.IsDir
//
//   + You can use '\t' escape sequence for indenting the subdirs.
//
//   + You can find a sample directory structure under:
//     solution/ directory
//
// ---------------------------------------------------------
func FindWriteSubdirectoriesToFile() {
	paths := os.Args[1:]
	if len(paths) == 0 {
		fmt.Println("Please provide directory paths")
	}

	dirs := make([]byte, 0, len(paths))

	for _, dir := range paths {
		infos, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println(err)
			return
		}

		dirs = append(dirs, dir...)
		dirs = append(dirs, '\n')

		for _, info := range infos {
			dirs = append(dirs, '\t')
			dirs = append(dirs, info.Name()...)
			dirs = append(dirs, '/', '\n')
		}

		dirs = append(dirs, '\n')
	}

	err := ioutil.WriteFile("dirs.txt", dirs, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func FindWriteSubdirectoriesToFileDA() {
	// get multiple directory paths from the command-line
	// prints only their subdirectories into a file named: dirs.txt
	paths := os.Args[1:]
	if len(paths) == 0 {
		fmt.Println("Please provide directory paths")
		return
	}

	var dirs []byte

	for _, dir := range paths {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println(err)
			return
		}

		dirs = append(dirs, dir...)
		dirs = append(dirs, '\n')

		for _, file := range files {
			if file.IsDir() {
				dirs = append(dirs, '\t')
				dirs = append(dirs, file.Name()...)
				dirs = append(dirs, '/', '\n')
			}
		}

		dirs = append(dirs, '\n')
	}

	err := ioutil.WriteFile("dirs.txt", dirs, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
