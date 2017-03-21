package queue

import "testing"
import "time"

import "fmt"

import "runtime"

func TestQueue(t *testing.T) {
	runtime.GOMAXPROCS(4)

	run("Queue", 100000, t)
	run1("EsQueue", 100000, t)
	run2("LockQueue", 100000, t)
	run3("ChanQueue", 100000, t)
	fmt.Println()
	run("Queue", 1000000, t)
	run1("EsQueue", 1000000, t)
	run2("LockQueue", 1000000, t)
	run3("ChanQueue", 1000000, t)
	fmt.Println()
	// run("Queue", 10000000, t)
	// run1("EsQueue", 10000000, t)
	// run2("LockQueue", 10000000, t)
	// run3("ChanQueue", 10000000, t)
}
func getSize(size uint32) uint32 {
	return size * 10
}

func run(name string, size uint32, t *testing.T) {
	var t1, t2 time.Duration
	var sum uint64
	q := NewQueue(getSize(size))

	// 生产
	// _t1 := time.Now()
	for i := uint32(0); i < size; i++ {
		go func(i uint32) {
			for {
				if q.Put(i) {
					break
				}
				runtime.Gosched()
			}
			// t1 += time.Now().Sub(_t1)
		}(i)
	}

	// 消费
	_t2 := time.Now()
	for i := uint32(0); i < size; i++ {
		for {
			if n, b := q.Get(); b {
				sum += uint64(n.(uint32))
				break
			}
			runtime.Gosched()
		}
	}
	t2 = time.Now().Sub(_t2)

	// 判断
	fmt.Print("name:", name, "\t")
	fmt.Print("size:", size, "\t")
	fmt.Print("put:", t1.String(), "\t")
	fmt.Print("get:", t2.String(), "\t")
	fmt.Print("sum:", sum, "\n")
}

func run1(name string, size uint32, t *testing.T) {
	var t1, t2 time.Duration
	var sum uint64
	q := NewQueue1(getSize(size))

	// 生产
	// _t1 := time.Now()
	for i := uint32(0); i < size; i++ {
		go func(i uint32) {
			for {
				if b, _ := q.Put(i); b {
					break
				}
				runtime.Gosched()
			}
			// t1 += time.Now().Sub(_t1)
		}(i)
	}

	// 消费
	_t2 := time.Now()
	for i := uint32(0); i < size; i++ {
		for {
			if n, b, _ := q.Get(); b {
				sum += uint64(n.(uint32))
				break
			}
			runtime.Gosched()
		}
	}
	t2 = time.Now().Sub(_t2)

	// 判断
	fmt.Print("name:", name, "\t")
	fmt.Print("size:", size, "\t")
	fmt.Print("put:", t1.String(), "\t")
	fmt.Print("get:", t2.String(), "\t")
	fmt.Print("sum:", sum, "\n")
}

func run2(name string, size uint32, t *testing.T) {
	var t1, t2 time.Duration
	var sum uint64
	q := NewQueueLock(getSize(size))

	// 生产
	// _t1 := time.Now()
	for i := uint32(0); i < size; i++ {
		go func(i uint32) {
			for {
				if q.Put(i) {
					break
				}
				runtime.Gosched()
			}
			// t1 += time.Now().Sub(_t1)
		}(i)
	}

	// 消费
	_t2 := time.Now()
	for i := uint32(0); i < size; i++ {
		for {
			if n, b := q.Get(); b {
				sum += uint64(n.(uint32))
				break
			}
			runtime.Gosched()
		}
	}
	t2 = time.Now().Sub(_t2)

	// 判断
	fmt.Print("name:", name, "\t")
	fmt.Print("size:", size, "\t")
	fmt.Print("put:", t1.String(), "\t")
	fmt.Print("get:", t2.String(), "\t")
	fmt.Print("sum:", sum, "\n")
}

func run3(name string, size uint32, t *testing.T) {
	var t1, t2 time.Duration
	var sum uint64
	q := NewQueueChannel(getSize(size))

	// 生产
	// _t1 := time.Now()
	for i := uint32(0); i < size; i++ {
		go func(i uint32) {
			for {
				if q.Put(i) {
					break
				}
				runtime.Gosched()
			}
			// t1 += time.Now().Sub(_t1)
		}(i)
	}

	// 消费
	_t2 := time.Now()
	for i := uint32(0); i < size; i++ {
		for {
			if n, b := q.Get(); b {
				sum += uint64(n.(uint32))
				break
			}
			runtime.Gosched()
		}
	}
	t2 = time.Now().Sub(_t2)

	// 判断
	fmt.Print("name:", name, "\t")
	fmt.Print("size:", size, "\t")
	fmt.Print("put:", t1.String(), "\t")
	fmt.Print("get:", t2.String(), "\t")
	fmt.Print("sum:", sum, "\n")
}
