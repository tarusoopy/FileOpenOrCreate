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
	basename := filepath.Base(filename)
	if _, err := os.Stat(basename); os.IsNotExist(err) {
		f, err := os.Create(basename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
