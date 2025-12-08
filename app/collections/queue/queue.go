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

package queue

// Queue is a generic first-in, first-out (FIFO) queue of values of type T.
// The zero value is an empty queue ready to use.
type Queue[T any] struct {
	data []T
}

// New creates an empty [Queue] for values of type T.
//
// The returned queue has no pre-allocated capacity.
func New[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

// WithCapacity creates an empty [Queue] with space reserved for cap elements.
//
// This can reduce allocations if the number of enqueued elements is known in advance.
func WithCapacity[T any](cap int) *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0, cap),
	}
}

// Enqueue adds v to the end of the queue.
func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

// Dequeue removes and returns the value at the front of the queue.
// If the queue is empty, Dequeue returns the zero value of T and false.
// If not it returns the front value and true.
func (q *Queue[T]) Dequeue() (v T, ok bool) {
	if len(q.data) == 0 {
		return v, false
	}

	v = q.data[0]
	q.data = q.data[1:]

	return v, true
}

// Len returns the number of elements currently stored in the queue.
func (q *Queue[T]) Len() int {
	return len(q.data)
}
