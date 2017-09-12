package droictx

import (
	"io"
	"sync"
	"time"

	"github.com/DroiTaipei/droipkg"
)

// copy from golang.org/pkg/context
// closedchan is a reusable closed channel.
var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

type contextKV struct {
	key   string
	value interface{}
}

// change Context to interface for compatibility, or just change?

type Context interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	GetString(key string) (value string, ok bool)
	GetInt(key string) (value int, ok bool)
	GetInt64(key string) (value int64, ok bool)
	Map() (ret map[string]interface{})
	Reset()
	SetTimeout(duration time.Duration, err droipkg.DroiError)
	ResetTimeout(duration time.Duration, err droipkg.DroiError)
	IsTimeout() bool
	Timeout() <-chan time.Time
	Finish()
	StopTimer() bool
	Done() <-chan struct{}
	SetHTTPHeaders(s Setter)
	HeaderMap() (ret map[string]string)
	HeaderSet(headerField, headerValue string)
}

type DoneContext struct {
	kv []contextKV
	// if we have to use in gRPC, maybe we could put golang.org/pkg/context here?
	done chan struct{}
	mu   sync.Mutex
	// use Timer for memory fast gc
	timeout *time.Timer
	// return error when timeout
	errTimeout droipkg.DroiError
	// for isTimeout()... timer cannot check whether is timeout status
	deadline time.Time
}

func (c *DoneContext) Set(key string, value interface{}) {
	args := c.kv
	n := len(args)
	for i := 0; i < n; i++ {
		kv := &args[i]
		if kv.key == key {
			kv.value = value
			return
		}
	}

	ca := cap(args)
	if ca > n {
		args = args[:n+1]
		kv := &args[n]
		kv.key = key
		kv.value = value
		c.kv = args
		return
	}

	kv := contextKV{
		key:   key,
		value: value,
	}
	c.kv = append(args, kv)
}

func (c *DoneContext) Get(key string) interface{} {
	args := c.kv
	n := len(args)
	for i := 0; i < n; i++ {
		kv := &args[i]
		if kv.key == key {
			return kv.value
		}
	}
	return nil
}

func (c *DoneContext) GetString(key string) (value string, ok bool) {
	v := c.Get(key)
	if v == nil {
		return
	}
	value, ok = v.(string)
	return
}

func (c *DoneContext) GetInt(key string) (value int, ok bool) {
	v := c.Get(key)
	if v == nil {
		return
	}
	value, ok = v.(int)
	return
}

func (c *DoneContext) GetInt64(key string) (value int64, ok bool) {
	v := c.Get(key)
	if v == nil {
		return
	}
	value, ok = v.(int64)
	return
}

func (c *DoneContext) Map() (ret map[string]interface{}) {
	ret = make(map[string]interface{})
	args := c.kv
	n := len(args)
	for i := 0; i < n; i++ {
		ret[args[i].key] = args[i].value
	}
	return
}

func (c *DoneContext) Reset() {
	args := c.kv
	n := len(args)
	for i := 0; i < n; i++ {
		v := args[i].value
		if vc, ok := v.(io.Closer); ok {
			vc.Close()
		}
	}
	c.kv = c.kv[:0]
}

func (c *DoneContext) SetTimeout(duration time.Duration, err droipkg.DroiError) {
	c.deadline = time.Now().Add(duration)
	c.timeout = time.NewTimer(duration)
	c.errTimeout = err
}

func (c *DoneContext) ResetTimeout(duration time.Duration, err droipkg.DroiError) {
	c.deadline = time.Now().Add(duration)
	c.timeout.Reset(duration)
	c.errTimeout = err
}

func (c *DoneContext) IsTimeout() bool {
	return time.Now().After(c.deadline)
}

func (c *DoneContext) Timeout() <-chan time.Time {
	if c.timeout == nil {
		return nil
	}
	return c.timeout.C
}

func (c *DoneContext) DroiErr() droipkg.DroiError {
	return c.errTimeout
}

func (c *DoneContext) Finish() {
	if c.done == nil {
		c.done = closedchan
	} else {
		close(c.done)
	}

	if c.timeout != nil {
		c.timeout.Stop()
	}
}

func (c *DoneContext) StopTimer() bool {
	return c.timeout.Stop()
}

// golang.org/pkg/context function

func (c *DoneContext) Done() <-chan struct{} {
	c.mu.Lock()
	if c.done == nil {
		c.done = make(chan struct{})
	}
	d := c.done
	c.mu.Unlock()
	return d
}

func (c *DoneContext) Err() error {
	return c.errTimeout
}

func (c *DoneContext) Value(key interface{}) interface{} {
	return nil
}

func (c *DoneContext) Deadline() (deadline time.Time, ok bool) {
	if !c.deadline.IsZero() {
		deadline = c.deadline
		ok = true
	}
	return
}
