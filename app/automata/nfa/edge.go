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

import _ "github.com/kdeconinck/realign/collections/mvmap"

// An 'edge' is a "single" transition from one [State] to another.
//
// Reasoning:
// When building a linear chain of transitions (which means that each symbol has 1 outgoing transition), it doesn't make
// a lot of sense to allocate a [mvmap.MvMap].
// Instead of using the [mvmap.MvMap], we use this struct for [State]s that only have a single transition.
// This reduces the amount of allocations because we don't have to allocate a [mvmap.MvMap] for each [State].
type edge[S comparable, V any] struct {
	has      bool
	symbol   S
	endState *State[S, V]
}

func newEdge[S comparable, V any](symbol S, endState *State[S, V]) edge[S, V] {
	return edge[S, V]{
		has:      true,
		symbol:   symbol,
		endState: endState,
	}
}
