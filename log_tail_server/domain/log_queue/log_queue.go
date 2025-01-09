package log_queue

import "fmt"

type LogQueue struct {
	logQueue []string
}

func (lq *LogQueue) Push(data string) {
	lq.logQueue = append(lq.logQueue, data)
}

func (lq *LogQueue) Pop() error {
	len := len(lq.logQueue)
	if len == 0 {
		return fmt.Errorf("no data to pop")
	}
	lq.logQueue = lq.logQueue[1: len-1]
	return nil
}

func (lq *LogQueue) GetList() []string {
	return lq.logQueue
}

func NewLogQueue() *LogQueue {
	return &LogQueue{}
}