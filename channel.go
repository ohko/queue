package queue

import "sync"

// QueueChannel 队列
type QueueChannel struct {
	queue  chan *qOneChannel // 数据
	putPos uint32            // 写位置
	getPos uint32            // 读位置
	max    uint32            // 容量
	lock   sync.Mutex
}

type qOneChannel struct {
	data interface{}
	used bool
}

// NewQueueChannel ...
func NewQueueChannel(size uint32) *QueueChannel {
	o := new(QueueChannel)
	o.queue = make(chan *qOneChannel, size)
	return o
}

// Get 读取一个
func (o *QueueChannel) Get() (interface{}, bool) {
	select {
	case oo := <-o.queue:
		return oo.data, true
	default:
		return nil, false
	}
}

// Put 写入一个
func (o *QueueChannel) Put(data interface{}) bool {
	select {
	case o.queue <- &qOneChannel{data: data}:
		return true
	default:
		return false
	}
}
