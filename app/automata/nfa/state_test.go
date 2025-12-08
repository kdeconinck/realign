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
	"sort"
	"testing"

	"github.com/kdeconinck/realign/assert"
	"github.com/kdeconinck/realign/automata/nfa"
)

// UT: Create a new 'State'.
func TestNfa_NewState(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[int, int]()

	// Act.
	state := machine.NewState()

	got := state.IsAccepting()

	// Assert.
	assert.Falsef(t, got, "\n\n"+
		"UT Name:  When creating a new 'State', it is NOT accepting.\n"+
		"\033[32mExpected: false.\033[0m\n"+
		"\033[31mActual:   %t.\033[0m\n\n", got)
}

// UT: Get the ID of a 'State'.
func TestState_ID(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[int, int]()

	startState := machine.Start()

	state1 := machine.NewState()
	state2 := machine.NewState()

	// Act.
	got1, got2, got3, want1, want2, want3 := startState.ID(), state1.ID(), state2.ID(), 0, 1, 2

	// Assert.
	assert.Equalf(t, got1, want1, "\n\n"+
		"UT Name:  The ID of the start 'State' is correct.\n"+
		"\033[32mExpected: %d.\033[0m\n"+
		"\033[31mActual:   %d.\033[0m\n\n", want1, got1)

	assert.Equalf(t, got2, want2, "\n\n"+
		"UT Name:  The ID of the 1st added 'State' is correct.\n"+
		"\033[32mExpected: %d.\033[0m\n"+
		"\033[31mActual:   %d.\033[0m\n\n", want2, got2)

	assert.Equalf(t, got3, want3, "\n\n"+
		"UT Name:  The ID of the 2nd added 'State' is correct.\n"+
		"\033[32mExpected: %d.\033[0m\n"+
		"\033[31mActual:   %d.\033[0m\n\n", want3, got3)
}

// UT: Check which 'State's are reachable by following a single epsilon transition.
func TestState_Epsilon(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[int, int]()

	state1 := machine.NewState()
	state2 := machine.NewState()
	state3 := machine.NewState()
	state4 := machine.NewState()

	machine.ConnectEpsilon(state1, state2)
	machine.ConnectEpsilon(state1, state3)
	machine.ConnectEpsilon(state2, state4)

	// Act.
	got, want := state1.Epsilon(), []*nfa.State[int, int]{state2, state3}

	// Assert.
	assert.EqualSf(t, got, want, "\n\n"+
		"UT Name:  Requesting the 'State's reachable following a single epsilon transition, returns the correct 'State's.\n"+
		"\033[32mExpected: %v.\033[0m\n"+
		"\033[31mActual:   %v.\033[0m\n\n", want, got)
}

// UT: Check which symbols can be consumed from a 'State'.
func TestState_OutgoingSymbols(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("Requesting the consumable symbols from a 'State' (without transitions), returns <nil>.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		machine := nfa.New[int, int]()

		state1 := machine.NewState()

		// Act.
		got := state1.OutgoingSymbols()

		// Assert.
		assert.Nilf(t, got, "\n\n"+
			"UT Name:  Requesting the consumable symbols from a 'State' (without transitions), returns <nil>.\n"+
			"\033[32mExpected: <nil>.\033[0m\n"+
			"\033[31mActual:   NOT <nil>.\033[0m\n\n")
	})

	t.Run("Requesting the consumable symbols from a 'State' (with a single transition), returns the correct symbols.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		machine := nfa.New[int, int]()

		state1 := machine.NewState()

		machine.Add(state1, 1)

		// Act.
		got, want := state1.OutgoingSymbols(), []int{1}

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Requesting the consumable symbols from a 'State' (with a single transition), returns the correct symbols.\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Requesting the consumable symbols from a 'State' (with multiple transitions), returns the correct symbols.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.
		// Arrange.
		machine := nfa.New[int, int]()

		state1 := machine.NewState()
		state2 := machine.NewState()

		machine.Add(state1, 1) // NOTE: Intentionally adding the same symbol multiple times.
		machine.Add(state1, 1)
		machine.Add(state1, 10)
		machine.Add(state1, 100)
		machine.Add(state2, 1000)

		// Act.
		got, want := state1.OutgoingSymbols(), []int{1, 10, 100}

		sort.Ints(got) // NOTE: The order is undefined, so we need to sort the slices before passing them to 'assert.EqualSf'.
		sort.Ints(want)

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Requesting the consumable symbols from a 'State' (with multiple transitions), returns the correct symbols.\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})
}

// UT: Check which 'State's are reachable by consuming a single symbol.
func TestState_OutgoingFor(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("Requesting the 'State's reachable by consuming a symbol (that is not defined), returns <nil>.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		machine := nfa.New[int, int]()

		state1 := machine.NewState()

		// Act.
		got := state1.OutgoingFor(10)

		// Assert.
		assert.Nilf(t, got, "\n\n"+
			"UT Name:  Requesting the 'State's reachable by consuming a symbol (that is not defined), returns <nil>.\n"+
			"\033[32mExpected: <nil>.\033[0m\n"+
			"\033[31mActual:   NOT <nil>.\033[0m\n\n")
	})

	t.Run("Requesting the 'State's reachable by consuming a symbol (on a state with a single transition), returns the correct 'State's.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		machine := nfa.New[int, int]()

		state1 := machine.NewState()
		state2 := machine.Add(state1, 1)

		// Act.
		got, want := state1.OutgoingFor(1), []*nfa.State[int, int]{state2}

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Requesting the 'State's reachable by consuming a symbol (on a state with a single transition), returns the correct 'State's.\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Requesting the 'State's reachable by onsuming a symbol (on a state with multiple transitions), returns the correct 'State's.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		machine := nfa.New[int, int]()

		state1 := machine.NewState()
		state2 := machine.Add(state1, 1)
		state3 := machine.Add(state1, 1)

		machine.Add(state1, 100)

		// Act.
		got, want := state1.OutgoingFor(1), []*nfa.State[int, int]{state2, state3}

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Requesting the 'State's reachable by onsuming a symbol (on a state with multiple transitions), returns the correct 'State's.\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})
}

// UT: Check if a 'State' is accepting.
func TestState_IsAccepting(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When the state is NOT accepting, false is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		machine := nfa.New[int, int]()

		state := machine.NewState()

		// Act.
		got := state.IsAccepting()

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  When the state is NOT accepting, false is returned.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("When the state is accepting, true is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		machine := nfa.New[int, int]()

		startState := machine.Start()

		state := machine.AddAcceptingEpsilonTransition(startState, 42)

		// Act.
		got := state.IsAccepting()

		// Assert.
		assert.Truef(t, got, "\n\n"+
			"UT Name:  When the state IS accepting, true is returned.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})
}

// UT: Get the acceptance value of a 'State'.
func TestState_AcceptValue(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When the state is NOT accepting, the default value is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		machine := nfa.New[int, int]()

		state := machine.NewState()

		// Act.
		got, want := state.AcceptValue(), 0

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When the state is NOT accepting, the default value is returned.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("When the state id accepting, the value is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		machine := nfa.New[int, int]()

		startState := machine.Start()

		state := machine.AddAcceptingEpsilonTransition(startState, 42)

		// Act.
		got, want := state.AcceptValue(), 42

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When the state IS accepting, the value is returned.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}
