package queue

import "sync/atomic"
import "runtime"

// Queue 队列
type Queue struct {
	queue  []qOne // 数据
	putPos uint32 // 写位置
	getPos uint32 // 读位置
	max    uint32 // 容量
}

type qOne struct {
	data interface{}
	used bool
}

// NewQueue ...
func NewQueue(size uint32) *Queue {
	o := new(Queue)
	o.max = size
	o.queue = make([]qOne, size)
	return o
}

// Get 读取一个
func (o *Queue) Get() (interface{}, bool) {
	var getPos, nextPos uint32

	// 查找下一个写坐标
	for {
		getPos = o.getPos
		nextPos = getPos + 1

		// 循环
		if nextPos >= o.max {
			nextPos = 0
		}
		if atomic.CompareAndSwapUint32(&o.getPos, getPos, nextPos) {
			if o.queue[nextPos].used {
				o.queue[nextPos].used = false
				return o.queue[nextPos].data, true
			}
			return nil, false
		}
		runtime.Gosched()
	}
}

// Put 写入一个
func (o *Queue) Put(data interface{}) bool {
	var putPos, nextPos uint32

	// 查找下一个写坐标
	for {
		putPos = o.putPos
		nextPos = putPos + 1

		// 循环
		if nextPos >= o.max {
			nextPos = 0
		}
		if atomic.CompareAndSwapUint32(&o.putPos, putPos, nextPos) {
			if !o.queue[nextPos].used {
				o.queue[nextPos].data = data
				o.queue[nextPos].used = true
				return true
			}
			return false
		}
		runtime.Gosched()
	}
}
