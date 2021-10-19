package goka

import (
	"testing"

	"github.com/eko/goka/internal/test"
)

func TestSignal_SetState(t *testing.T) {

	sig := NewSignal(0, 1, 2)
	test.AssertTrue(t, sig.IsState(0))
	test.AssertFalse(t, sig.IsState(1))

	sig.SetState(1)
	test.AssertTrue(t, sig.IsState(1))
	test.AssertFalse(t, sig.IsState(0))

	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("Expected panic, which didn't occur")
		}
	}()

	// set some invalid state, this will panic
	sig.SetState(3)
}

func TestSignal_Wait(t *testing.T) {
	sig := NewSignal(0, 1, 2)

	<-sig.WaitForState(0)
	// should continue right now since

	var (
		done     = make(chan struct{})
		hasState bool
	)
	go func() {
		defer close(done)
		<-sig.WaitForState(1)
		hasState = true
	}()

	test.AssertFalse(t, hasState)
	sig.SetState(1)
	// wait for the goroutine to catchup with the state
	<-done
	test.AssertTrue(t, hasState)
}

func TestSignalWaitMin(t *testing.T) {
	sig := NewSignal(0, 1, 2)

	var (
		done     = make(chan struct{})
		hasState bool
	)
	go func() {
		defer close(done)
		<-sig.WaitForStateMin(1)
		hasState = true
	}()

	test.AssertFalse(t, hasState)
	sig.SetState(2)
	// wait for the goroutine to catchup with the state
	<-done
	test.AssertTrue(t, hasState)
}
