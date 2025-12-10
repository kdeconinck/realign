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

package dfa

// State is a node in a [Dfa].
type State[S comparable, V any] struct {
	id          int
	transitions map[S]*State[S, V]
	acceptIdx   int
	value       V // The accepting value (if any).
}

// IsAccepting returns true if the state is an accepting state.
func (s *State[S, V]) IsAccepting() bool { return s.acceptIdx > -1 }

// AcceptValue returns the accepting value of the state.
// If the state is not accepting, it returns the zero value of V.
func (s *State[S, V]) AcceptValue() V { return s.value }

// OutgoingFor returns the target state for the given symbol, or nil if no transition exists.
func (s *State[S, V]) OutgoingFor(symbol S) *State[S, V] {
	return s.transitions[symbol]
}

// Returns a new [State].
func (d *Dfa[S, V]) newState() *State[S, V] {
	id := d.nextStateID
	d.nextStateID++

	return &State[S, V]{
		id:          id,
		transitions: make(map[S]*State[S, V]),
		acceptIdx:   -1,
	}
}

// Returns a new accepting [State].
func (d *Dfa[S, V]) newAcceptingState(acceptIdx int, value V) *State[S, V] {
	state := d.newState()
	state.acceptIdx = acceptIdx
	state.value = value

	return state
}
