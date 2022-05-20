package blocks

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/francoispqt/gojay"
)

type Display interface {
	Display() []byte
}

type BuildContext struct {
	BaseCooldown      time.Duration
	Channel           string
	Context           context.Context
	CooldownThreshold time.Duration
	CurrentCooldown   time.Duration
	DataText          []byte
	ErrorText         []byte
	Forced            bool
	Iteration         int
	Message           string
	Next              Timing
	Previous          Timing
}

func (bc BuildContext) MarshalJSON() ([]byte, error) {
	return gojay.MarshalJSONObject(&bc)
}

type Timing struct {
	CooldownDebounce time.Time
	Send             time.Time
	Threshold        time.Time
}

func (t Timing) MarshalJSON() ([]byte, error) {
	return gojay.MarshalJSONObject(&t)
}

func (t Timing) MarshalJSONObject(enc *gojay.Encoder) {
	enc.TimeKey("cooldown_debounce", &t.CooldownDebounce, time.RFC3339)
	enc.TimeKey("send", &t.Send, time.RFC3339)
	enc.TimeKey("threshold", &t.Threshold, time.RFC3339)
}

func (t Timing) IsNil() bool {
	return false
}

func (bc *BuildContext) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Int64Key("base_cooldown", int64(bc.BaseCooldown))
	enc.StringKey("channel", bc.Channel)
	enc.Int64Key("cooldown_threshold", int64(bc.CooldownThreshold))
	enc.Int64Key("currentCooldown", int64(bc.CurrentCooldown))
	enc.AddStringKey("data", base64.StdEncoding.EncodeToString(bc.DataText))
	enc.AddStringKey("error", base64.StdEncoding.EncodeToString(bc.ErrorText))
	enc.BoolKey("forced", bc.Forced)
	enc.IntKey("iteration", bc.Iteration)
	enc.StringKey("message", bc.Message)
	enc.AddObjectKey("next", bc.Next)
	enc.AddObjectKey("previous", bc.Previous)
}

func (bc *BuildContext) IsNil() bool {
	return bc == nil
}

// implements context.Context.
func (bc *BuildContext) Deadline() (deadline time.Time, ok bool) {
	return bc.Context.Deadline()
}

// implements context.Context.
func (bc *BuildContext) Done() <-chan struct{} {
	return bc.Context.Done()
}

// implements context.Context
func (bc *BuildContext) Err() error {
	return bc.Err()
}

// implements context.Context.
func (bc *BuildContext) Value(key interface{}) interface{} {
	return bc.Context.Value(key)
}
