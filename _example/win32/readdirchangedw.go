package main

import (
	"fmt"
	go_watcher "github.com/kisun-bit/go-watcher"
)

func main() {
	w, err := go_watcher.NewWatcher("C:\\", true)
	if err != nil {
		fmt.Println("NewWatcher Error: %s", err.Error())
		return
	}
	cs, err := w.Start()
	if err != nil {
		fmt.Println("Start Error: %s", err.Error())
		return
	}

	for c := range cs {
		switch c.Action {
		case go_watcher.Win32EventCreate:
			fmt.Println("++++++++++++++++++++ ", c.String())
		case go_watcher.Win32EventDelete:
			fmt.Println("-------------------- ", c.String())
		case go_watcher.Win32EventUpdate:
			fmt.Println("#################### ", c.String())
		case go_watcher.Win32EventRenameTo:
			fmt.Println(">>>>>>>>>>>>>>>>>>>> ", c.String())
		case go_watcher.Win32EventRenameFrom:
			fmt.Println("<<<<<<<<<<<<<<<<<<<<", c.String())
		}
	}
}
