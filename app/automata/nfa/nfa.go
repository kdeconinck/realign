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

package nfa

// Nfa represents a non-deterministic finite automaton for symbols of type S with acceptance metadata of type V.
type Nfa[S comparable, V any] struct {
	startState      *State[S, V]
	nextStateID     int
	nextAcceptIndex int
}

// New creates a new [Nfa] with an initial start [State].
func New[S comparable, V any]() *Nfa[S, V] {
	machine := &Nfa[S, V]{
		nextStateID:     0,
		nextAcceptIndex: 0,
	}

	machine.startState = machine.newState()

	return machine
}

// Start returns the start [State] of the nfa.
func (machine *Nfa[S, V]) Start() *State[S, V] { return machine.startState }

// Add adds and returns a new transition starting from startState for symbol.
// Adding a transition causes a new [State] to be generated.
func (machine *Nfa[S, V]) Add(startState *State[S, V], symbol S) *State[S, V] {
	state := machine.newState()
	startState.put(symbol, state)

	return state
}

// AddEpsilonTransition adds and returns an epsilon transition starting from startState.
// Adding an epsilon transition causes a new [State] to be generated.
func (machine *Nfa[S, V]) AddEpsilonTransition(startState *State[S, V]) *State[S, V] {
	state := machine.NewState()
	machine.ConnectEpsilon(startState, state)

	return state
}

// AddAcceptingEpsilonTransition adds and returns an epsilon transition starting from startState.
// Adding an epsilon transition causes a new accepting [State] with metadata value to be generated.
func (machine *Nfa[S, V]) AddAcceptingEpsilonTransition(startState *State[S, V], value V) *State[S, V] {
	state := machine.newAcceptingState(value)
	startState.eTransitions = append(startState.eTransitions, state)

	return state
}

// AddPredicateTransition adds and returns a new predicate transition from startState to endState.
// The predicate function fn is used to determine if the transition is valid for a given symbol.
func (machine *Nfa[S, V]) AddPredicateTransition(startState, endState *State[S, V], fn func(S) bool) {
	transition := predicateTransition[S, V]{
		EndState: endState,
		Fn:       fn,
	}

	startState.predicateTransitions = append(startState.predicateTransitions, transition)
}

// ConnectEpsilon adds an epsilon transition from from to to.
func (machine *Nfa[S, V]) ConnectEpsilon(startState *State[S, V], endState *State[S, V]) {
	startState.eTransitions = append(startState.eTransitions, endState)
}

func (machine *Nfa[S, V]) newState() *State[S, V] {
	id := machine.nextStateID
	machine.nextStateID++

	return &State[S, V]{
		id:                   id,
		predicateTransitions: nil,
		eTransitions:         nil,
		acceptIdx:            -1,
	}
}

func (machine *Nfa[S, V]) newAcceptingState(value V) *State[S, V] {
	id := machine.nextStateID
	machine.nextStateID++

	state := &State[S, V]{
		id:           id,
		eTransitions: nil,
		acceptIdx:    -1,
	}

	machine.markAccepting(state, value)

	return state
}

func (machine *Nfa[S, V]) markAccepting(state *State[S, V], value V) {
	idx := machine.nextAcceptIndex
	machine.nextAcceptIndex++

	state.acceptIdx = idx
	state.value = value
}
