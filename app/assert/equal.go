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

package assert

import (
	_ "fmt"
	"slices"
	"testing"
)

// Equalf asserts that got and want are equal.
//
// Equalf is generic over any [comparable] type V and compares values using the == operator.
// If got and want are not equal, Equalf calls [testing.TB.Fatalf] with a message formatted according to format and args,
// in the same style as [fmt.Sprintf].
//
// For slice equality, use [EqualSf].
func Equalf[V comparable](tb testing.TB, got, want V, format string, args ...any) {
	tb.Helper()

	if got != want {
		tb.Fatalf(format, args...)
	}
}

// EqualSf asserts that the slices got and want are equal.
//
// EqualSf is generic over any slice type S whose element type V is [comparable]. It compares slices using
// [slices.Equal]. If got and want are not equal, EqualSf calls [testing.TB.Fatalf] with a message formatted according
// to format and args, in the same style as [fmt.Sprintf].
func EqualSf[S ~[]V, V comparable](tb testing.TB, got, want S, format string, args ...any) {
	tb.Helper()

	if !slices.Equal(got, want) {
		tb.Fatalf(format, args...)
	}
}
