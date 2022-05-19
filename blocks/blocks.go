package blocks

import (
	"context"
	"time"
)

type Display interface {
	Display() []byte
}

type BuildContext struct {
	baseCooldown    time.Duration
	currentCooldown time.Duration
	channel         string
	context         context.Context
	dataText        []byte
	errorText       []byte
}

// implements context.Context.
func (bc *BuildContext) Deadline() (deadline time.Time, ok bool) {
	return bc.context.Deadline()
}

// implements context.Context.
func (bc *BuildContext) Done() <-chan struct{} {
	return bc.context.Done()
}

// implements context.Context
func (bc *BuildContext) Err() error {
	return bc.Err()
}

// // implements context.Context.
func (bc *BuildContext) Value(key interface{}) interface{} {
	return bc.context.Value(key)
}

// Gets the base cooldown for current build.
func (b BuildContext) BaseCooldown() time.Duration {
	return b.baseCooldown
}

// Gets the cooldown for current iteration.
func (b BuildContext) CurrentCooldown() time.Duration {
	return b.currentCooldown
}

// Gets the data text.
func (b BuildContext) DataText() []byte {
	return b.dataText
}

// Gets the channel the message will be sent to.
func (b BuildContext) Channel() string {
	return b.channel
}
