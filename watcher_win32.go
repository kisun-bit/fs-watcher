package go_watcher

import (
	"os"
	"syscall"
	"time"
	"unsafe"
)

const (
	WatchingBufferSize = 1024
	WatchingQueueSize  = 4
)

type Win32Watcher struct {
	root           string
	rootPtr        *uint16
	rootDescriptor syscall.Handle
	recursive      bool
	buffer         []byte
	notify         chan FileWatchingObject
}

func NewWatcher(Dir string, IsRecursive bool) (w *Win32Watcher, err error) {
	if _, err = os.Stat(Dir); err != nil {
		return nil, err
	}

	w = new(Win32Watcher)
	w.root = Dir
	w.recursive = IsRecursive
	w.buffer = make([]byte, WatchingBufferSize)
	w.notify = make(chan FileWatchingObject, WatchingQueueSize)
	return w, err
}

func (w *Win32Watcher) Start() (_ chan FileWatchingObject, err error) {
	if err = w.changeDirAccessRight(); err != nil {
		return nil, err
	}

	if err = w.asyncLoop(); err != nil {
		return nil, err
	}
	return w.notify, nil
}

func (w *Win32Watcher) changeDirAccessRight() (err error) {
	if w.rootPtr, err = syscall.UTF16PtrFromString(w.root); err != nil {
		return err
	}

	w.rootDescriptor, err = syscall.CreateFile(w.rootPtr,
		syscall.FILE_LIST_DIRECTORY,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE|syscall.FILE_SHARE_DELETE,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_FLAG_BACKUP_SEMANTICS,
		0,
	)
	if err != nil {
		return err
	}
	return nil
}

func (w *Win32Watcher) asyncLoop() (err error) {
	go func() {
		defer syscall.CloseHandle(w.rootDescriptor)
		defer close(w.notify)

		for {
			if err = w.readDirectoryChanges(); err != nil {
				break
			}
			w.collectChangeInfo()
		}
	}()
	return
}

func (w *Win32Watcher) readDirectoryChanges() (err error) {
	return syscall.ReadDirectoryChanges(
		w.rootDescriptor,
		&w.buffer[0],
		WatchingBufferSize,
		w.recursive,
		syscall.FILE_NOTIFY_CHANGE_FILE_NAME|
			syscall.FILE_NOTIFY_CHANGE_DIR_NAME|
			syscall.FILE_NOTIFY_CHANGE_ATTRIBUTES|
			syscall.FILE_NOTIFY_CHANGE_SIZE|
			syscall.FILE_NOTIFY_CHANGE_LAST_WRITE,
		nil,
		&syscall.Overlapped{},
		0,
	)
}

func (w *Win32Watcher) collectChangeInfo() {
	var offset uint32
	for {
		raw := (*syscall.FileNotifyInformation)(unsafe.Pointer(&w.buffer[offset]))
		buf := (*[syscall.MAX_PATH]uint16)(unsafe.Pointer(&raw.FileName))
		name := syscall.UTF16ToString(buf[:raw.FileNameLength/2])
		info := FileWatchingObject{
			Action:      convertToWin32Event(raw.Action),
			AbsPath:     w.root + name,
			TriggerTime: time.Now(),
		}
		w.notify <- info
		if raw.NextEntryOffset == 0 {
			break
		}
		offset += raw.NextEntryOffset
		if offset >= WatchingBufferSize {
			break
		}
	}
}
