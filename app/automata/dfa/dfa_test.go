// =====================================================================================================================
// = LICENSE:       Copyright (c) 2025 Kevin De Coninck
// =
// =                Permission is hereby granted, free of charge, to any person
// =                obtaining a copy of this software and associated documentation
// =                files (the "Software"), to deal in the Software without
// =                restriction, including without limitation the rights to use,
// =                copy, modify, merge, publish, distribute, sublicense, and/or sell
// =                copies of the Software, and to permit persons to whom the
// =                Software is furnished to do so, subject to the following
// =                conditions:
// =
// =                The above copyright notice and this permission notice shall be
// =                included in all copies or substantial portions of the Software.
// =
// =                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// =                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// =                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// =                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// =                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// =                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// =                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// =                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

package dfa_test

import (
	"testing"

	"github.com/kdeconinck/realign/assert"
	"github.com/kdeconinck/realign/automata/dfa"
	"github.com/kdeconinck/realign/automata/nfa"
)

// UT: Convert a linear 'Nfa' to a 'Dfa'.
func TestFromNfa_Linear(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	nMachine := nfa.New[int, string]()
	nStart := nMachine.Start()
	nEnd := nMachine.Add(nStart, 1)
	nMachine.AddAcceptingEpsilonTransition(nEnd, "OK")

	// Act.
	dMachine := dfa.FromNfa(nMachine)

	// Assert: Start State.
	dStart := dMachine.Start()
	assert.NotNilf(t, dStart, "Start state should not be nil")
	assert.Falsef(t, dStart.IsAccepting(), "Start state should NOT be accepting")

	// Assert: Transition(s).
	dNext := dStart.OutgoingFor(1)
	assert.NotNilf(t, dNext, "Transition for symbol '1' should exist")
	assert.Truef(t, dNext.IsAccepting(), "Next state should be accepting")
	assert.Equalf(t, dNext.AcceptValue(), "OK", "Accept value should be 'OK'")
}

// UT: Convert an 'Nfa' with multiple transitions on the same symbol (subset construction).
func TestFromNfa_FanOut(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	nMachine := nfa.New[int, string]()
	nStart := nMachine.Start()

	// Two transitions on the same symbol '1' leading to different outcomes
	// Path 1: 1 -> Accept("A")
	s1 := nMachine.Add(nStart, 1)
	nMachine.AddAcceptingEpsilonTransition(s1, "A")

	// Path 2: 1 -> Accept("B")
	s2 := nMachine.Add(nStart, 1)
	nMachine.AddAcceptingEpsilonTransition(s2, "B")

	// Act.
	dMachine := dfa.FromNfa(nMachine)

	// Assert.
	dStart := dMachine.Start()
	dNext := dStart.OutgoingFor(1)

	assert.NotNilf(t, dNext, "Transition for symbol '1' should exist")
	assert.Truef(t, dNext.IsAccepting(), "Resulting state should be accepting")
}

// UT: Convert an 'Nfa' with epsilon transitions.
func TestFromNfa_Epsilon(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	nMachine := nfa.New[int, int]()
	nStart := nMachine.Start()

	// Start --(epsilon)--> S1 --(1)--> S2 (Accept)
	s1 := nMachine.AddEpsilonTransition(nStart)
	s2 := nMachine.Add(s1, 1)
	nMachine.AddAcceptingEpsilonTransition(s2, 42)

	// Act.
	dMachine := dfa.FromNfa(nMachine)

	// Assert.
	// The DFA start state should effectively represent {nStart, s1}
	// And transition on '1' should go to a state representing {s2, ...} (accepting)
	dStart := dMachine.Start()
	dNext := dStart.OutgoingFor(1)

	assert.NotNilf(t, dNext, "Transition for symbol '1' should exist from start")
	assert.Truef(t, dNext.IsAccepting(), "Destination state should be accepting")
	assert.Equalf(t, dNext.AcceptValue(), 42, "Value should be 42")
}

// UT: Convert an 'Nfa' with a loop.
func TestFromNfa_Loop(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	nMachine := nfa.New[int, int]()
	nStart := nMachine.Start()

	// Start --(1)--> Start (Loop)
	nMachine.Connect(nStart, nStart, 1)
	// Start is accepting
	nMachine.AddAcceptingEpsilonTransition(nStart, 100)

	// Act.
	dMachine := dfa.FromNfa(nMachine)

	// Assert.
	dStart := dMachine.Start()
	assert.Truef(t, dStart.IsAccepting(), "Start state should be accepting")

	dNext := dStart.OutgoingFor(1)
	assert.NotNilf(t, dNext, "Transition for symbol '1' should exist")
	assert.Truef(t, dNext == dStart, "Transition '1' should loop back to start")
}

var benchmarkOutput *dfa.Dfa[int, int] // Output of the benchmark(s).

// Benchmark: Linear NFA to DFA.
func BenchmarkFromNfa_Linear_10(b *testing.B)   { benchmarkFromNfa_Linear(10, b) }
func BenchmarkFromNfa_Linear_100(b *testing.B)  { benchmarkFromNfa_Linear(100, b) }
func BenchmarkFromNfa_Linear_1000(b *testing.B) { benchmarkFromNfa_Linear(1_000, b) }

// Benchmark: FanOut NFA to DFA.
func BenchmarkFromNfa_FanOut_10(b *testing.B)   { benchmarkFromNfa_FanOut(10, b) }
func BenchmarkFromNfa_FanOut_100(b *testing.B)  { benchmarkFromNfa_FanOut(100, b) }
func BenchmarkFromNfa_FanOut_1000(b *testing.B) { benchmarkFromNfa_FanOut(1_000, b) }

func benchmarkFromNfa_Linear(count int, b *testing.B) {
	nMachine := nfa.New[int, int]()
	current := nMachine.Start()
	for i := 0; i < count; i++ {
		current = nMachine.Add(current, i)
	}

	b.ResetTimer()
	for b.Loop() {
		benchmarkOutput = dfa.FromNfa(nMachine)
	}
}

func benchmarkFromNfa_FanOut(count int, b *testing.B) {
	nMachine := nfa.New[int, int]()
	nStart := nMachine.Start()

	for i := 0; i < count; i++ {
		// All transitions on same symbol '1'
		nMachine.Add(nStart, 1)
	}

	b.ResetTimer()
	for b.Loop() {
		benchmarkOutput = dfa.FromNfa(nMachine)
	}
}
