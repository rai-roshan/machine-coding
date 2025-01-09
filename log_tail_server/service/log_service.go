package log_service

import (
	"log_tail_server/domain/file_system"
	"log_tail_server/domain/log_queue"
)

type LogService interface {
	CheckClientConnection() (bool, error)
	AddClientConnection() error
	PublishLogsOnWriteEvent() error
	PublishInititalLogs() error // 1. add to queue, 2. publish
}

type logService struct {
	connectionRecord map[uint32]bool
	logQueue         *log_queue.LogQueue
	fileSystem       *file_system.FileSystem
}

func NewLogService(logQueue *log_queue.LogQueue, fileSystem *file_system.FileSystem) LogService {
	return &logService{
		logQueue:         logQueue,
		fileSystem:       fileSystem,
		connectionRecord: make(map[uint32]bool),
	}
}

func (logService) CheckClientConnection() (bool, error) {
	return false, nil
}
func (logService) AddClientConnection() error {
	return nil
}
func (logService) PublishLogsOnWriteEvent() error {
	return nil
}
func (logService) PublishInititalLogs() error {
	return nil
}
