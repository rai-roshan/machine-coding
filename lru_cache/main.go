package main

import (
	"fmt"
	"sync"
)

type Node struct {
	key   string
	data  uint32
	Left  *Node
	Right *Node
}

type Dequeue interface {
	PushBack(key string, value uint32) *Node
	PopFromFront()
	MoveToBack(node *Node)
	GetSize() uint32
	GetFront() *Node
	GetBack() *Node
}

type DequeueImp struct {
	Front *Node
	Back  *Node
	Size  uint32
}

func NewDequeue() Dequeue {
	return &DequeueImp{
		Front: nil,
		Back:  nil,
		Size:  0,
	}
}

func (dq *DequeueImp) PushBack(key string, value uint32) *Node {
	newNode := &Node{
		key:  key,
		data: value,
	}

	if dq.Front == nil && dq.Back == nil {
		dq.Front = newNode
		dq.Back = newNode
	} else {
		newNode.Left = dq.Back
		dq.Back.Right = newNode
		dq.Back = newNode
	}
	dq.Size += 1
	return newNode
}

func (dq *DequeueImp) PopFromFront() {
	if dq.Front != nil {
		dq.Front = dq.Front.Right
		dq.Size -= 1
	}
}

func (dq *DequeueImp) MoveToBack(node *Node) {
	nodeLeft := node.Left
	nodeRight := node.Right

	if dq.Size == 1 {
		return
	}
	if nodeLeft != nil {
		nodeLeft.Right = nodeRight
	}
	if nodeRight != nil {
		nodeRight.Left = nodeLeft
	}

	if node == dq.Front {
		dq.Front = nodeRight
	}
	dq.Back.Right = node
	node.Left = dq.Back
	node.Right = nil
	dq.Back = node
}

func (dq *DequeueImp) GetSize() uint32 {
	return dq.Size
}

func (dq *DequeueImp) GetFront() *Node {
	return dq.Front
}

func (dq *DequeueImp) GetBack() *Node {
	return dq.Back
}

type LRU interface {
	Insert(key string, value uint32)
	Get(key string) int
}

type LRUImp struct {
	size         uint32
	keyToNodeMap map[string]*Node
	queue        Dequeue
	lock         sync.RWMutex
}

func NewLRU(size uint32) LRU {
	return &LRUImp{
		size:         size,
		keyToNodeMap: make(map[string]*Node),
		queue:        NewDequeue(),
	}
}

func (lru *LRUImp) Insert(key string, value uint32) {
	lru.lock.Lock()
	defer func() {
		lru.lock.Unlock()
	}()

	// 1. check if the key already present
	node, ok := lru.keyToNodeMap[key]
	if ok {
		node.data = value
		lru.queue.MoveToBack(node)
	} else {
		newNode := lru.queue.PushBack(key, value)
		lru.keyToNodeMap[key] = newNode
	}

	fmt.Println("font : ", *lru.queue.GetFront())

	// 2. check size of queue
	if lru.queue.GetSize() > lru.size {
		frontNode := lru.queue.GetFront()
		fmt.Println("front to delete : ", frontNode)
		delete(lru.keyToNodeMap, frontNode.key)
		lru.queue.PopFromFront()
	}
}

func (lru *LRUImp) Get(key string) int {

	lru.lock.RLock()
	node, ok := lru.keyToNodeMap[key]
	if !ok {
		fmt.Printf("key : %s || value : %d\n", key, -1)
		lru.lock.RUnlock()
		return -1
	}
	lru.lock.RUnlock()

	lru.lock.Lock()
	lru.queue.MoveToBack(node)
	lru.lock.Unlock()

	fmt.Printf("key : %s || value : %d\n", key, node.data)
	return int(node.data)
}

func main() {
	lru := NewLRU(uint32(3))
	lru.Insert("rai0", uint32(0))
	lru.Insert("rai1", uint32(1))
	lru.Insert("rai2", uint32(2))
	// lru.Insert("rai3", uint32(3))
	lru.Get("rai0")
	lru.Get("rai1")
	lru.Get("rai2")
	lru.Get("rai0")

	lru.Insert("rai3", uint32(3))
	fmt.Println("====")
	lru.Get("rai0")
	lru.Get("rai1")
	lru.Get("rai2")
	lru.Get("rai3")

}
