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

package assert_test

import (
	"fmt"
	"testing"
)

// TbSpy wraps a [testing.TB] and records the last failure message passed to [TbSpy.Fatalf] instead of failing the test
// immediately.
//
// It is intended for use in tests of the assert package itself, to verify that assertion helpers call Fatalf with the
// expected message.
type TbSpy struct {
	testing.TB
	failureMsg string
}

// SpyTb returns a new [TbSpy] that wraps the provided [testing.TB].
//
// The returned spy forwards all methods to tb except for Fatalf, which is overridden to capture the formatted failure
// message instead of marking the test as failed.
func SpyTb(tb testing.TB) *TbSpy {
	return &TbSpy{
		TB: tb,
	}
}

// Fatalf records the formatted failure message in failureMsg.
//
// Fatalf does not call the embedded [testing.TB.Fatalf], so using TbSpy does not fail the underlying test. This allows
// tests to inspect the message that would have been reported by assertion helpers.
func (tb *TbSpy) Fatalf(format string, args ...any) {
	tb.failureMsg = fmt.Sprintf(format, args...)
}

func newSlice[T any](args ...T) []T {
	container := make([]T, len(args))
	copy(container, args)

	return container
}
