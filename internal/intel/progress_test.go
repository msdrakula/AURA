package intel

import "testing"

func TestEmitTickThrottles(t *testing.T) {
	var n int
	e := &Engine{Progress: func(ProgressEvent) { n++ }}
	for i := 1; i <= 100; i++ {
		e.emitTick("dirs", i, 100, 25, ProgressEvent{})
	}
	if n < 5 || n > 8 {
		t.Fatalf("expected ~5 ticks (1,25,50,75,100), got %d", n)
	}
}

func TestEmitNilProgressSafe(t *testing.T) {
	var e Engine
	e.emit(ProgressEvent{Action: "start", Stage: "tech"})
}
