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

package nfa_test

import (
	"testing"

	"github.com/kdeconinck/realign/assert"
	"github.com/kdeconinck/realign/automata/nfa"
)

// UT: Get the start 'State' of an 'Nfa'.
func TestNfa_Start(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[int, int]()

	// Act.
	got := machine.Start()

	// Assert.
	assert.NotNilf(t, got, "\n\n"+
		"UT Name:  When creating a new 'Nfa', the start 'State' is NOT nil.\n"+
		"\033[32mExpected: NOT <nil>.\033[0m\n"+
		"\033[31mActual:   <nil>.\033[0m\n\n")
}

// UT: Add a new transition to an 'Nfa'.
func TestNfa_Add(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[int, int]()
	startState := machine.Start()

	// Act.
	state := machine.Add(startState, 42)

	// Assert.
	assert.NotNilf(t, state, "\n\n"+
		"UT Name:  When adding a transition, a new 'State' is returned.\n"+
		"\033[32mExpected: NOT <nil>.\033[0m\n"+
		"\033[31mActual:   <nil>.\033[0m\n\n")
}

// UT: Add multiple transitions to an 'Nfa'.
func TestNfa_AddMultipleTransitions(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[int, int]()
	startState := machine.Start()

	// Act.
	state1 := machine.Add(startState, 1)
	state2 := machine.Add(startState, 2)

	// Assert.
	assert.NotNilf(t, state1, "\n\n"+
		"UT Name:  When adding a transition ('State #1'), a new 'State' is returned.\n"+
		"\033[32mExpected: NOT <nil>.\033[0m\n"+
		"\033[31mActual:   <nil>.\033[0m\n\n")

	assert.NotNilf(t, state2, "\n\n"+
		"UT Name:  When adding a transition ('State #2'), a new 'State' is returned.\n"+
		"\033[32mExpected: NOT <nil>.\033[0m\n"+
		"\033[31mActual:   <nil>.\033[0m\n\n")
}

// UT: Add an epsilon transition to an 'Nfa'.
func TestNfa_AddEpsilonTransition(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[int, int]()
	startState := machine.Start()

	// Act.
	state := machine.AddEpsilonTransition(startState)

	// Assert.
	assert.NotNilf(t, state, "\n\n"+
		"UT Name:  When adding an epsilon transition, a new 'State' is returned.\n"+
		"\033[32mExpected: NOT <nil>.\033[0m\n"+
		"\033[31mActual:   <nil>.\033[0m\n\n")
}

// UT: Add an accepting epsilon transition to an 'Nfa'.
func TestNfa_AddAcceptingEpsilonTransition(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[int, int]()
	startState := machine.Start()

	// Act.
	state := machine.AddAcceptingEpsilonTransition(startState, 42)

	// Assert.
	assert.NotNilf(t, state, "\n\n"+
		"UT Name:  When adding an accepting epsilon transition, a new 'State' is returned.\n"+
		"\033[32mExpected: NOT <nil>.\033[0m\n"+
		"\033[31mActual:   <nil>.\033[0m\n\n")

	t.Run("When adding an accepting epsilon transition, the new 'State' is accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got := state.IsAccepting()

		assert.Truef(t, got, "\n\n"+
			"UT Name:  When adding an accepting epsilon transition, the new 'State' is accepting.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("When adding an accepting epsilon transition, the acceptance value is correct.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := state.AcceptValue(), 42

		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When adding an accepting epsilon transition, the acceptance value is correct.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}

// UT: Add a predicate transition to an `Nfa`.
func TestNfa_AddPredicateTransition(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[int, int]()
	startState := machine.Start()
	endState := machine.NewState()

	// Act.
	fn := func() { machine.AddPredicateTransition(startState, endState, func(i int) bool { return i > 0 }) }

	// Assert.
	assert.NoPanicf(t, fn, "\n\n"+
		"UT Name:  When adding a predicate transition, the function should NOT panic.\n"+
		"\033[32mExpected: NOT panic.\033[0m\n"+
		"\033[31mActual:   panic.\033[0m\n\n")
}

// UT: Connect two 'State's with an epsilon transition.
func TestNfa_ConnectEpsilon(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[int, int]()
	startState := machine.Start()
	endState := machine.NewState()

	// Act.
	fn := func() { machine.ConnectEpsilon(startState, endState) }

	// Assert.
	assert.NoPanicf(t, fn, "\n\n"+
		"UT Name:  When connecting two 'State's with an epsilon transition, the function should NOT panic.\n"+
		"\033[32mExpected: NOT panic.\033[0m\n"+
		"\033[31mActual:   panic.\033[0m\n\n")
}

var benchmarkOutput *nfa.State[int, int] // Output of the benchmark(s). Used to avoid compiler optimizations.

// Benchmark(s): Add elements in a linear chain.
func BenchmarkNfa_Add_Linear_1(b *testing.B)       { benchmarkNfa_Add_Linear(1, b) }
func BenchmarkNfa_Add_Linear_10(b *testing.B)      { benchmarkNfa_Add_Linear(10, b) }
func BenchmarkNfa_Add_Linear_100(b *testing.B)     { benchmarkNfa_Add_Linear(100, b) }
func BenchmarkNfa_Add_Linear_1000(b *testing.B)    { benchmarkNfa_Add_Linear(1_000, b) }
func BenchmarkNfa_Add_Linear_1000000(b *testing.B) { benchmarkNfa_Add_Linear(1_000_000, b) }

// Benchmark(s): Add elements from a single state (fan-out).
func BenchmarkNfa_Add_FanOut_1(b *testing.B)       { benchmarkNfa_Add_FanOut(1, b) }
func BenchmarkNfa_Add_FanOut_10(b *testing.B)      { benchmarkNfa_Add_FanOut(10, b) }
func BenchmarkNfa_Add_FanOut_100(b *testing.B)     { benchmarkNfa_Add_FanOut(100, b) }
func BenchmarkNfa_Add_FanOut_1000(b *testing.B)    { benchmarkNfa_Add_FanOut(1_000, b) }
func BenchmarkNfa_Add_FanOut_1000000(b *testing.B) { benchmarkNfa_Add_FanOut(1_000_000, b) }

func benchmarkNfa_Add_Linear(count int, b *testing.B) {
	var startState *nfa.State[int, int]

	for b.Loop() {
		machine := nfa.New[int, int]()
		startState = machine.Start()

		for idx := range count {
			startState = machine.Add(startState, idx)
		}
	}

	benchmarkOutput = startState
}

func benchmarkNfa_Add_FanOut(count int, b *testing.B) {
	var state *nfa.State[int, int]

	for b.Loop() {
		machine := nfa.New[int, int]()
		startState := machine.Start()

		for idx := range count {
			state = machine.Add(startState, idx)
		}
	}

	benchmarkOutput = state
}
