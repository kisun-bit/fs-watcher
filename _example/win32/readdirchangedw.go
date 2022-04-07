package main

import (
	"fmt"
	watcher "github.com/kisun-bit/go-watcher"
)

func main() {
	w, err := watcher.NewWatcher("C:\\", true)
	if err != nil {
		fmt.Println("NewWatcher Error: ", err.Error())
		return
	}
	cs, err := w.Start()
	if err != nil {
		fmt.Println("Start Error: ", err.Error())
		return
	}

	for c := range cs {
		switch c.Action {
		case watcher.Win32EventCreate:
			fmt.Println("++++++++++++++++++++ ", c.String())
		case watcher.Win32EventDelete:
			fmt.Println("-------------------- ", c.String())
		case watcher.Win32EventUpdate:
			fmt.Println("#################### ", c.String())
		case watcher.Win32EventRenameTo:
			fmt.Println(">>>>>>>>>>>>>>>>>>>> ", c.String())
		case watcher.Win32EventRenameFrom:
			fmt.Println("<<<<<<<<<<<<<<<<<<<< ", c.String())
		}
	}
}
