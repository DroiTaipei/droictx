package droictx

import (
	"testing"
	"time"

	"github.com/DroiTaipei/droipkg"
)

func TestDone(t *testing.T) {
	ctx := &DoneContext{}
	ctx.SetTimeout(5000*time.Millisecond, nil)
	go func(ctx Context) {
		time.Sleep(500 * time.Millisecond)
		ctx.Finish()
	}(ctx)

	select {
	case <-ctx.Done():
		// check finish() have stopped timer
		if ctx.StopTimer() {
			t.Error("Finish() haven't stop timer")
		}
	case <-ctx.Timeout():
		t.Error("done doesn't work")
	}
}

func TestTimeout(t *testing.T) {
	err := droipkg.ConstDroiError("1000000 timeout")
	ctx := &DoneContext{}
	ctx.SetTimeout(50*time.Millisecond, err)

	select {
	case <-time.After(1 * time.Second):
		ctx.StopTimer()
		t.Error("context overslept")
	case <-ctx.Timeout():
		errTimeout := ctx.DroiErr()
		if errTimeout.Error() != err.Error() {
			t.Error("timeout error fail")
		}
	}

	if !ctx.IsTimeout() {
		t.Error("isTimeout doesn't work")
	}
}
