package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var level = 0

func main() {

	// err := filepath.Walk("E:\\GO\\src",

	// 	func(path string, info os.FileInfo, err error) error {

	// 		if err != nil {
	// 			return err
	// 		}

	// 		fmt.Println(path, info.Size())
	// 		return err
	// 	})

	// if err != nil {

	// 	log.Println(err)
	// }
	ScanDirectory("E:\\GO\\TicTacToe")
}

func ScanDirectory(path string) {

	display_filename_prefix_middle := "├──"
	display_filename_prefix_last := "└──"
	// display_parent_prefix_middle := "    "
	// display_parent_prefix_last := "│   "
	// fmt.Print(display_parent_prefix_middle)
	for i := 0; i < level; i++ {
		fmt.Print("    ")

	}
	level++
	fmt.Println(display_filename_prefix_middle, path)
	// fmt.Println(" ")

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			ScanDirectory(filePath)
		} else {

			for i := 0; i < level; i++ {
				fmt.Print("    ")

			}
			fmt.Print(display_filename_prefix_last)
			fmt.Println(file.Name())

		}

	}
	// fmt.Print(display_filename_prefix_last)
	// fmt.Print(display_parent_prefix_last)
	level--

}
