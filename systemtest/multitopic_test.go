package systemtest

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"github.com/lovoo/goka/internal/test"
	"github.com/lovoo/goka/multierr"
)

// Tests a processor with multiple input topics. Random values are emitted to random topics, the values are accumulated
// for a single key and checked for correctness after emitting a couple of messages.
// This is a regression/showcase test for https://github.com/lovoo/goka/issues/332
func TestMultiTopics(t *testing.T) {

	var (
		group        goka.Group = goka.Group(fmt.Sprintf("%s-%d", "goka-systemtest-multitopic", time.Now().Unix()))
		table                   = goka.GroupTable(group)
		inputStreams []goka.Stream
	)

	for i := 0; i < 5; i++ {
		inputStreams = append(inputStreams, goka.Stream(fmt.Sprintf("%s-input-%d", string(group), i)))
	}

	if !*systemtest {
		t.Skipf("Ignoring systemtest. pass '-args -systemtest' to `go test` to include them")
	}

	tmc := goka.NewTopicManagerConfig()
	tmc.Table.Replication = 1
	tmc.Stream.Replication = 1
	cfg := goka.DefaultConfig()
	tm, err := goka.TopicManagerBuilderWithConfig(cfg, tmc)([]string{*broker})
	test.AssertNil(t, err)

	for _, inStream := range inputStreams {
		err = tm.EnsureStreamExists(string(inStream), 1)
		test.AssertNil(t, err)
	}
	// let the cluster create it
	time.Sleep(5 * time.Second)

	proc, err := goka.NewProcessor([]string{*broker},
		goka.DefineGroup(
			group,
			goka.Inputs(inputStreams, new(codec.Int64), func(ctx goka.Context, msg interface{}) {
				var oldVal int64

				if val := ctx.Value(); val != nil {
					oldVal = val.(int64)
				}

				log.Printf("%s: %03d+%03d->%03d", ctx.Topic(), oldVal, msg, msg.(int64)+oldVal)

				// accumulate with old value
				ctx.SetValue(msg.(int64) + oldVal)
			}),
			goka.Persist(new(codec.Int64)),
		),
		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)),
	)
	test.AssertNil(t, err)

	view, err := goka.NewView([]string{*broker}, table, new(codec.Int64))
	test.AssertNil(t, err)

	var emitters []*goka.Emitter

	for _, input := range inputStreams {
		emitter, err := goka.NewEmitter([]string{*broker}, input, new(codec.Int64))
		test.AssertNil(t, err)
		emitters = append(emitters, emitter)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errg, ctx := multierr.NewErrGroup(ctx)

	errg.Go(func() error {
		return proc.Run(ctx)
	})
	errg.Go(func() error {
		return view.Run(ctx)
	})

	log.Printf("waiting for processor/view to be running")
	pollTimed(t, "proc and view are recovered", 10.0, proc.Recovered, view.Recovered)
	log.Printf("...done")

	var sum int64
	for i := int64(0); i < 100; i++ {
		value := rand.Int63n(100)
		// emit to random emitters in sync
		err := emitters[rand.Intn(len(emitters))].EmitSync("key", value)
		test.AssertNil(t, err)
		// ... and batched
		prom, err := emitters[rand.Intn(len(emitters))].Emit("key", value)
		test.AssertNil(t, err)
		prom.Then(func(err error) {
			test.AssertNil(t, err)
		})

		// accumulate what we have sent so far
		sum += (value * 2)
	}

	for _, emitter := range emitters {
		test.AssertNil(t, emitter.Finish())
	}

	// poll the view and the processor until we're sure that we have
	pollTimed(t, "all messages have been transferred", 10.0,
		func() bool {
			value, err := view.Get("key")
			test.AssertNil(t, err)
			return value != nil && value.(int64) == sum
		},
		func() bool {
			value, err := proc.Get("key")
			test.AssertNil(t, err)
			return value != nil && value.(int64) == sum
		},
	)

	// stop everything and wait until it's shut down
	cancel()
	test.AssertNil(t, errg.Wait().NilOrError())
}
