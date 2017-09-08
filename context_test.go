package droictx

import (
	"testing"
	"time"
)

func TestDone(t *testing.T) {
	ctx := &DoneContext{}
	ctx.SetTimeout(5000 * time.Millisecond)
	go func(ctx Context) {
		time.Sleep(500 * time.Millisecond)
		ctx.Finish()
	}(ctx)

	select {
	case <-ctx.Done():
		ctx.StopTimer()
	case <-ctx.Timeout():
		t.Error("done doesn't work")
	}
}

func TestTimeout(t *testing.T) {
	ctx := &DoneContext{}
	ctx.SetTimeout(50 * time.Millisecond)

	select {
	case <-time.After(1 * time.Second):
		ctx.StopTimer()
		t.Error("context overslept")
	case <-ctx.Timeout():
	}

	if !ctx.IsTimeout() {
		t.Error("isTimeout doesn't work")
	}
}
