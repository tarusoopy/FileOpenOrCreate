package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filename := "/var/log/sample/sample.log"
	pathdir := filepath.Dir(filename)
	if _, err := os.Stat(pathdir); os.IsNotExist(err) {
		err = os.Mkdir(pathdir, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err = os.Create(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
