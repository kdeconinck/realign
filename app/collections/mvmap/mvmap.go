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

package mvmap

// MvMap is a generic multi-value map from keys of type K to slices of values of type V.
//
// For each key, MvMap stores zero or more values in a slice. Missing keys return a nil slice.
// The zero value is not ready for use; call [New] or [WithCapacity] to initialize the map.
type MvMap[K comparable, V any] struct {
	data map[K][]V
}

// New creates an empty [MvMap] with keys of type K and values of type V.
//
// The underlying map is initialized without any capacity.
func New[K comparable, V any]() *MvMap[K, V] {
	return &MvMap[K, V]{
		data: make(map[K][]V),
	}
}

// WithCapacity creates an empty [MvMap] with space reserved for cap keys.
//
// The capacity applies to the underlying map from keys to slices, not to the per-key value slices. To reserve capacity
// for values of a specific key, use [MvMap.SetKeyCap].
func WithCapacity[K comparable, V any](cap int) *MvMap[K, V] {
	return &MvMap[K, V]{
		data: make(map[K][]V, cap),
	}
}

// SetKeyCap initializes the slice for key k with zero length and the given capacity.
// This can be used to preallocate space for values that will be added via [MvMap.Put].
// Any existing values for k are discarded.
func (m *MvMap[K, V]) SetKeyCap(k K, size int) {
	m.data[k] = make([]V, 0, size)
}

// Put appends v to the slice of values associated with key k.
//
// If k has no existing values, Put creates a new slice containing v.
func (m *MvMap[K, V]) Put(k K, v V) {
	m.data[k] = append(m.data[k], v)
}

// Get returns the slice of values associated with key k.
//
// If k has no associated values, Get returns nil. The returned slice is the one stored inside the map; callers should
// not modify it if other goroutines may access the map concurrently.
func (m *MvMap[K, V]) Get(k K) []V {
	return m.data[k]
}

// Len returns the number of keys stored in the map.
//
// Each key may have zero or more associated values.
func (m *MvMap[K, V]) Len() int {
	return len(m.data)
}

// Keys returns a slice containing all keys currently stored in the map.
//
// The order of keys is not specified and may vary from one call to the next.
// The returned slice is a copy; callers are allowed to modify it at any point.
func (m *MvMap[K, V]) Keys() []K {
	out := make([]K, 0, m.Len())

	for v := range m.data {
		out = append(out, v)
	}

	return out
}
