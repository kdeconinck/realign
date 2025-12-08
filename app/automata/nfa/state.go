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

import "github.com/kdeconinck/realign/collections/mvmap"

// State represents a single state in an [Nfa] over symbols of type S and accepting values of type V.
//
// A state may have:
//   - Zero or more outgoing transitions on concrete symbols.
//   - Zero or more predicate-based transitions.
//   - Zero or more epsilon transitions.
//   - An optional accepting index and value.
type State[S comparable, V any] struct {
	id                   int
	edge                 edge[S, V]
	transitions          *mvmap.MvMap[S, *State[S, V]]
	predicateTransitions []predicateTransition[S, V]
	eTransitions         []*State[S, V]
	acceptIdx            int
	value                V // The accepting value (if any).
}

// NewState returns a new, non-accepting [State].
func (machine *Nfa[S, V]) NewState() *State[S, V] {
	id := machine.nextStateID
	machine.nextStateID++

	return &State[S, V]{
		id:                   id,
		predicateTransitions: nil,
		transitions:          nil,
		eTransitions:         nil,
		acceptIdx:            -1,
	}
}

// ID returns the unique, builder-assigned identifier (starting at 0).
func (state *State[S, V]) ID() int { return state.id }

// AcceptIdx returns the acceptance index of the state.
func (state *State[S, V]) AcceptIdx() int {
	return state.acceptIdx
}

// IsAccepting reports whether the state is accepting.
func (state *State[S, V]) IsAccepting() bool { return state.AcceptIdx() > -1 }

// AcceptValue returns the accepting value if any of the default value of V if the state is NOT accepting.
func (state *State[S, V]) AcceptValue() V {
	var defaultValue V

	if state.acceptIdx <= -1 {
		return defaultValue
	}

	return state.value
}

// Epsilon returns the reachable [State]s following epsilon transitions from the state.
func (state *State[S, V]) Epsilon() []*State[S, V] {
	return state.eTransitions
}

// OutgoingSymbols returns all the symbols that have at least one outgoing transition from this state.
// Note: The order is undefined.
func (state *State[S, V]) OutgoingSymbols() []S {
	if state.transitions == nil {
		if !state.edge.has {
			return nil
		}

		return []S{
			state.edge.symbol,
		}
	}

	return state.transitions.Keys()
}

// OutgoingFor returns all the [State]s reachable from the state by consuming symbol or nil if there are no transitions.
func (state *State[S, V]) OutgoingFor(symbol S) []*State[S, V] {
	if state.transitions == nil {
		if state.edge.has && state.edge.symbol == symbol {
			return []*State[S, V]{
				state.edge.endState,
			}
		}

		return nil
	}

	return state.transitions.Get(symbol)
}

func (state *State[S, V]) put(symbol S, endState *State[S, V]) {
	if !state.edge.has && state.transitions == nil {
		state.edge = newEdge(symbol, endState)

		return
	}

	if state.edge.has {
		if state.transitions == nil {
			state.transitions = mvmap.New[S, *State[S, V]]()
		}

		state.transitions.Put(state.edge.symbol, state.edge.endState)
		state.resetEdge()
	}

	state.transitions.Put(symbol, endState)
}

func (state *State[S, V]) resetEdge() {
	state.edge = edge[S, V]{
		// NOTE: Intentionally left blank.
	}
}
