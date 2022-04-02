package go_watcher

import (
	"fmt"
	"time"
)

type FileWatchingObject struct {
	Action      Event
	AbsPath     string
	TriggerTime time.Time
}

func (fwo *FileWatchingObject) String() string {
	return fmt.Sprintf("<FileWatchingObject | Action(%s) Path(%s) Time(%s)>",
		fwo.Action.String(), fwo.AbsPath, fwo.TriggerTime.Format("2006-01-02 15:04:05"))
}
