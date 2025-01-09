package file_watcher

type FileWatcher struct {
	filePath string
}

func (fw *FileWatcher) ListenWriteEvent(callback func() error) {
	callback() // pass event and event error
}

func NewFileWatcher(filePath string) *FileWatcher {
	return &FileWatcher{
		filePath: filePath,
	}
}

