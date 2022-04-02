package go_watcher

import "runtime"

type Event uint32

const (
	Win32EventCreate Event = iota + 1
	Win32EventDelete
	Win32EventUpdate
	Win32EventRenameFrom
	Win32EventRenameTo
)

func (e Event) String() string {
	switch runtime.GOOS {
	case "windows":
		switch e {
		case Win32EventCreate:
			return "Create"
		case Win32EventDelete:
			return "Delete"
		case Win32EventUpdate:
			return "Update"
		case Win32EventRenameFrom:
			return "RenameFrom"
		case Win32EventRenameTo:
			return "RenameTo"
		}
	}
	return ""
}
