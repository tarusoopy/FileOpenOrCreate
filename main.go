package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filename := "/home/nakano/sample/sample.log"
	pathdir := filepath.Dir(filename)
    fmt.Println(pathdir)
	if _, err := os.Stat(pathdir); os.IsNotExist(err) {
		err = os.Mkdir(pathdir, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
    fmt.Println(filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err = os.Create(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
