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
	"testing"

	"github.com/kdeconinck/realign/assert"
)

// UT: Compare the length of a slice against 0.
func Test_Emptyf(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	const msgFmt string = "Not equal - got %d, want 0."

	for tcName, tc := range map[string]struct {
		gotInput []int
		want     string
	}{
		"When 'got' has NO elements, the assertion should NOT fail.": {
			gotInput: newSlice[int](),
		},
		"When 'got' has elements, the assertion should fail.": {
			gotInput: newSlice(1, 2, 3, 4, 5),
			want:     "Not equal - got 5, want 0.",
		},
	} {
		tcName, tc := tcName, tc // Rebind: Needed for parallel support.

		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// Arrange.
			tbSpy := SpyTb(t)

			// Act.
			assert.Emptyf(tbSpy, tc.gotInput, msgFmt, len(tc.gotInput))

			// Assert.
			if tbSpy.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, tc.want)
			}
		})
	}
}

// UT: Compare the length of a slice against a number (greater than).
func Test_GreaterThanf(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	const msgFmt string = "Not equal - got %d, want > %d."

	for tcName, tc := range map[string]struct {
		gotInput  []int
		wantInput int
		want      string
	}{
		"When 'got' has enough elements, the assertion should NOT fail.": {
			gotInput:  newSlice(1, 2, 3),
			wantInput: 2,
		},
		"When 'got' has the exact amount of elements, the assertion should fail.": {
			gotInput:  newSlice(1, 2),
			wantInput: 2,
			want:      "Not equal - got 2, want > 2.",
		},
		"When 'got' has not enough elements, the assertion should fail.": {
			gotInput:  newSlice(1),
			wantInput: 2,
			want:      "Not equal - got 1, want > 2.",
		},
	} {
		tcName, tc := tcName, tc // Rebind: Needed for parallel support.

		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// Arrange.
			tbSpy := SpyTb(t)

			// Act.
			assert.GreaterThanf(tbSpy, tc.gotInput, tc.wantInput, msgFmt, len(tc.gotInput), tc.wantInput)

			// Assert.
			if tbSpy.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, tc.want)
			}
		})
	}
}

// UT: Compare the length of a slice against a number (less than).
func Test_LessThanf(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	const msgFmt string = "Not equal - got %d, want < %d."

	for tcName, tc := range map[string]struct {
		gotInput  []int
		wantInput int
		want      string
	}{
		"When 'got' has not too much elements, the assertion should NOT fail.": {
			gotInput:  newSlice(1, 2, 3),
			wantInput: 5,
		},
		"When 'got' has the exact amount of elements, the assertion should fail.": {
			gotInput:  newSlice(1, 2),
			wantInput: 2,
			want:      "Not equal - got 2, want < 2.",
		},
		"When 'got' has too much elements, the assertion should fail.": {
			gotInput:  newSlice(1, 2, 3),
			wantInput: 2,
			want:      "Not equal - got 3, want < 2.",
		},
	} {
		tcName, tc := tcName, tc // Rebind: Needed for parallel support.

		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// Arrange.
			tbSpy := SpyTb(t)

			// Act.
			assert.LessThanf(tbSpy, tc.gotInput, tc.wantInput, msgFmt, len(tc.gotInput), tc.wantInput)

			// Assert.
			if tbSpy.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, tc.want)
			}
		})
	}
}
