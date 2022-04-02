package go_watcher

func convertToWin32Event(no uint32) Event {
	if no > uint32(Win32EventRenameTo) || no < uint32(Win32EventCreate) {
		panic("invalid win32 event")
	}
	return Event(no)
}
