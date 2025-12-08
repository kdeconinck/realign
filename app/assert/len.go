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
	"testing"
)

// Emptyf asserts that the slice got is empty.
//
// Emptyf is generic over any slice type S. The assertion succeeds if got contains no elements. If got contains any
// elements, Emptyf calls [testing.TB.Fatalf] with a message formatted according to format and args, in the same style
// as [fmt.Sprintf].
func Emptyf[S ~[]V, V any](tb testing.TB, got S, format string, args ...any) {
	tb.Helper()

	if len(got) > 0 {
		tb.Fatalf(format, args...)
	}
}

// GreaterThanf asserts that the length of the slice got is greater than want.
//
// GreaterThanf is generic over any slice type S. The assertion succeeds if got contains more elements than want.
// If got doesn't contain more elements than want, GreaterThanf calls [testing.TB.Fatalf] with a message formatted
// according to format and args, in the same style as [fmt.Sprintf].
func GreaterThanf[S ~[]V, V any](tb testing.TB, got S, want int, format string, args ...any) {
	tb.Helper()

	if len(got) <= want {
		tb.Fatalf(format, args...)
	}
}

// LessThanf asserts that the length of the slice got is less than want.
//
// LessThanf is generic over any slice type S. The assertion succeeds if got contains fewer elements than want.
// If got doesn't contain fewer elements than want, LessThanf calls [testing.TB.Fatalf] with a message formatted
// according to format and args, in the same style as [fmt.Sprintf].
func LessThanf[S ~[]V, V any](tb testing.TB, got S, want int, format string, args ...any) {
	tb.Helper()

	if len(got) >= want {
		tb.Fatalf(format, args...)
	}
}
