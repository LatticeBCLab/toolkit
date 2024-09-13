package sync

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	m    sync.Mutex
	done atomic.Uint32 // 0 代表没有执行，1代表已经执行了
}

// Do will return the error from the function f if initialization fails.
func (o *Once) Do(f func() error) error {
	if o.done.Load() == 0 { // fast path
		return o.doSlow(f)
	}
	return nil
}

func (o *Once) doSlow(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done.Load() == 0 { // 双检查
		err = f()
		if err == nil { // 初始化成功才改变 done
			o.done.Store(1)
		}
	}
	return err
}

// Done 返回此Once是否执行过
func (o *Once) Done() bool {
	return o.done.Load() == 1
}
