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

// Package assert provides a fluent, comprehensive set of assertion functions for Go's standard testing framework.
//
// It simplifies writing test cases by providing rich, readable assertion methods that accept a [testing.TB] instance
// as the first argument. When an assertion fails, it calls [testing.TB.Fatalf] internally.
//
// Basic usage example:
//
//	func TestMyFunction(t *testing.T) {
//	    result := MyFunction()
//
//	    assert.Truef(t, result > 0, "Expected result (%d) to be positive", result)
//	}
//
// Key assertion categories:
//
//   - Equality: [Equalf], [EqualSf], [Nilf] and [NotNilf].
//   - Boolean: [Truef] and [Falsef].
//   - Comparison: [Emptyf], [GreaterThanf] and [LessThanf].
//   - Error handling: [Errorf], [Panicf], [NoPanicf].
//
// Each function includes the ability to format a custom failure message, similar to [fmt.SPrintf].
package assert

import _ "fmt"
