// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

// copied from
// https://github.com/cockroachdb/cockroach/blob/master/pkg/rpc/breaker.go
// see the license & copyright info above

package netutil

import (
	"time"

	"github.com/lni/goutils/netutil/cenk/backoff"
	"github.com/lni/goutils/netutil/facebookgo/clock"
	"github.com/lni/goutils/netutil/rubyist/circuitbreaker"
)

const maxBackoff = time.Second

type breakerClock struct{}

func (c *breakerClock) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

func (c *breakerClock) AfterFunc(d time.Duration, f func()) *clock.Timer {
	panic("unimplemented")
}

func (c *breakerClock) Now() time.Time {
	return time.Now()
}

func (c *breakerClock) Sleep(d time.Duration) {
	panic("unimplemented")
}

func (c *breakerClock) Tick(d time.Duration) <-chan time.Time {
	panic("unimplemented")
}

func (c *breakerClock) Ticker(d time.Duration) *clock.Ticker {
	panic("unimplemented")
}

func (c *breakerClock) Timer(d time.Duration) *clock.Timer {
	panic("unimplemented")
}

var _ clock.Clock = &breakerClock{}

// newBackOff creates a new exponential backoff properly configured for RPC
// connection backoff.
func newBackOff(clock backoff.Clock) backoff.BackOff {
	// This exponential backoff limits the circuit breaker to 1 second
	// intervals between successive attempts to resolve a node address
	// and connect via GRPC.
	//
	// NB (nota Ben): MaxInterval should be less than the Raft election timeout
	// (1.5s) to avoid disruptions. A newly restarted node will be in follower
	// mode with no knowledge of the Raft leader. If it doesn't hear from a
	// leader before the election timeout expires, it will start to campaign,
	// which can be disruptive. Therefore the leader needs to get in touch (via
	// Raft heartbeats) with such nodes within one election timeout of their
	// restart, which won't happen if their backoff is too high.
	b := &backoff.ExponentialBackOff{
		InitialInterval:     500 * time.Millisecond,
		RandomizationFactor: 0.5,
		Multiplier:          1.5,
		MaxInterval:         maxBackoff,
		MaxElapsedTime:      0,
		Clock:               clock,
	}
	b.Reset()
	return b
}

// NewBreaker returns a new breaker.
func NewBreaker() *circuit.Breaker {
	clock := &breakerClock{}
	return circuit.NewBreakerWithOptions(&circuit.Options{
		BackOff:    newBackOff(clock),
		Clock:      clock,
		ShouldTrip: circuit.ThresholdTripFunc(1),
	})
}
