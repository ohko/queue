package queue

import "sync"

// QueueLock 队列
type QueueLock struct {
	queue  []qOneLock // 数据
	putPos uint32     // 写位置
	getPos uint32     // 读位置
	max    uint32     // 容量
	lock   sync.Mutex
}

type qOneLock struct {
	data interface{}
	used bool
}

// NewQueueLock ...
func NewQueueLock(size uint32) *QueueLock {
	o := new(QueueLock)
	o.max = size
	o.queue = make([]qOneLock, size)
	return o
}

// Get 读取一个
func (o *QueueLock) Get() (interface{}, bool) {
	o.lock.Lock()
	defer o.lock.Unlock()

	if o.queue[o.getPos].used {
		o.queue[o.getPos].used = false
		r := o.queue[o.getPos].data
		o.getPos++
		if o.getPos >= o.max {
			o.getPos = 0
		}
		return r, true
	}
	return nil, false
}

// Put 写入一个
func (o *QueueLock) Put(data interface{}) bool {
	o.lock.Lock()
	defer o.lock.Unlock()

	if !o.queue[o.putPos].used {
		o.queue[o.putPos].data = data
		o.queue[o.putPos].used = true
		o.putPos++
		if o.putPos >= o.max {
			o.putPos = 0
		}
		return true
	}
	return false
}
