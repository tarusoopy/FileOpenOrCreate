package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filename := "/var/log/vario/sample.log"
	pathdir := filepath.Dir(filename)
	if _, err := os.Stat(pathdir); os.IsNotExist(err) {
		err = os.Mkdir(pathdir, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	_, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
