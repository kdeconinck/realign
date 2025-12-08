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

package set

// Set is a generic set of values of type T.
//
// Internally it is represented as a map[T]struct{}. The zero value of Set is nil; use [New] or [WithCapacity] to
// initialize a non-nil set before adding values.
type Set[T comparable] map[T]struct {
	// NOTE: Intentionally left blank.
}

// New creates an empty [Set] for values of type T.
func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}

// WithCapacity creates an empty [Set] with space reserved for cap elements.
//
// This can reduce allocations if the number of elements to be added is known in advance.
func WithCapacity[T comparable](cap int) Set[T] {
	return make(map[T]struct{}, cap)
}

// Add inserts v into the set.
//
// If v is already present, Add has no effect.
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// Has reports whether v is in the set.
func (s Set[T]) Has(v T) bool {
	_, ok := s[v]

	return ok
}

// Values returns a slice containing all elements of the set.
//
// The order of values is not specified and may vary from one call to the next.
func (s Set[T]) Values() []T {
	out := make([]T, 0, len(s))

	for v := range s {
		out = append(out, v)
	}

	return out
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s)
}
