package file_system

type FileSystem struct {
	filePath string
	offset   uint32
}

func (fs *FileSystem) GetRecentTenLines() (data []string) {

	return
}

func (fs *FileSystem) GetNewLinesFromLastOffset() (data []string) {
	return
}

func NewFileSystem(filePath string) *FileSystem {
	return &FileSystem{
		filePath: filePath,
		offset: 0,
	}
} 