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

package dfa_test

import (
	"testing"

	"github.com/kdeconinck/realign/assert"
	"github.com/kdeconinck/realign/automata/dfa"
	"github.com/kdeconinck/realign/automata/nfa"
)

// UT: Convert a complex 'Nfa' to a 'Dfa'.
func TestFromNfa_Complex(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	nMachine := nfa.New[string, string]()
	s0 := nMachine.Start()

	// s0 --'a'--> s1
	// s0 --'a'--> s2
	s1 := nMachine.Add(s0, "a")
	s2 := nMachine.Add(s0, "a")

	// s2 --eps--> s4 (Accepting "WIN", Priority 1 / Index 0)
	// We create this first to ensure lower AcceptIdx
	s4 := nMachine.AddAcceptingEpsilonTransition(s2, "WIN")

	// s1 --eps--> s3 (Accepting "LOSE", Priority 2 / Index 1)
	s3 := nMachine.AddAcceptingEpsilonTransition(s1, "LOSE")

	// s4 --'b'--> s5
	s5 := nMachine.Add(s4, "b")

	// s5 --'c'--> s5 (Self Loop via epsilon bridge)
	s5Loop := nMachine.Add(s5, "c")
	nMachine.ConnectEpsilon(s5Loop, s5)

	// s3 --'b'--> s0 (Loop back via epsilon bridge)
	s3Loop := nMachine.Add(s3, "b")
	nMachine.ConnectEpsilon(s3Loop, s0)

	// Act.
	dMachine := dfa.FromNfa(nMachine)
	dStart := dMachine.Start()

	// Assert 1: Start State (Key: {s0})
	t.Run("Start state is not accepting and transitions on 'a'", func(t *testing.T) {
		assert.Falsef(t, dStart.IsAccepting(), "\n\n"+
			"UT Name:  Start state should not be accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   true.\033[0m\n\n")
		assert.NotNilf(t, dStart.OutgoingFor("a"), "\n\n"+
			"UT Name:  Start state should transition on 'a'.\n"+
			"\033[32mExpected: NOT <nil>.\033[0m\n"+
			"\033[31mActual:   <nil>.\033[0m\n\n")
		assert.Nilf(t, dStart.OutgoingFor("b"), "\n\n"+
			"UT Name:  Start state should not transition on 'b'.\n"+
			"\033[32mExpected: <nil>.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", dStart.OutgoingFor("b"))
	})

	// Assert 2: State A (Key: {s1, s2, s3, s4})
	// s3 is accepting ("LOSE", idx 1), s4 is accepting ("WIN", idx 0).
	// Priority rule: lowest index wins. Expected: "WIN".
	stateA := dStart.OutgoingFor("a")
	t.Run("Merged state follows priority rules", func(t *testing.T) {
		assert.Truef(t, stateA.IsAccepting(), "\n\n"+
			"UT Name:  StateA should be accepting.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   false.\033[0m\n\n")
		assert.Equalf(t, stateA.AcceptValue(), "WIN", "\n\n"+
			"UT Name:  StateA should value 'WIN' due to priority.\n"+
			"\033[32mExpected: WIN.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", stateA.AcceptValue())
		assert.NotNilf(t, stateA.OutgoingFor("b"), "\n\n"+
			"UT Name:  StateA should transition on 'b'.\n"+
			"\033[32mExpected: NOT <nil>.\033[0m\n"+
			"\033[31mActual:   <nil>.\033[0m\n\n")
	})

	// Assert 3: State AB (Key: {s3Loop, s0, s5})
	// s3Loop -> s0, s4 -> s5.
	stateAB := stateA.OutgoingFor("b")
	t.Run("Cycle correctly links back", func(t *testing.T) {
		// Transition 'a' from State AB should lead back to State A (since s0 --a--> {s1,s2} -> A)
		// s5 has no 'a'.
		stateABA := stateAB.OutgoingFor("a")
		assert.Equalf(t, stateABA, stateA, "\n\n"+
			"UT Name:  Transition 'a' from StateAB should loop back to StateA.\n"+
			"\033[32mExpected: Same State pointer.\033[0m\n"+
			"\033[31mActual:   Different.\033[0m\n\n")
	})

	// Assert 4: State ABC (Key: {s5Loop, s5})
	// s5 --c--> s5Loop -> s5.
	stateABC := stateAB.OutgoingFor("c")
	t.Run("Self loop persists", func(t *testing.T) {
		assert.NotNilf(t, stateABC, "\n\n"+
			"UT Name:  StateAB should transition on 'c'.\n"+
			"\033[32mExpected: NOT <nil>.\033[0m\n"+
			"\033[31mActual:   <nil>.\033[0m\n\n")
		// Transition 'c' from State ABC should loop to itself
		stateABCC := stateABC.OutgoingFor("c")
		assert.Equalf(t, stateABCC, stateABC, "\n\n"+
			"UT Name:  StateABC should self-loop on 'c'.\n"+
			"\033[32mExpected: Same State pointer.\033[0m\n"+
			"\033[31mActual:   Different.\033[0m\n\n")
	})
}

var benchmarkOutput *dfa.Dfa[int, int] // Output of the benchmark(s).

// Benchmark: Linear NFa to Dfa.
func BenchmarkFromNfa_Linear_10(b *testing.B)   { benchmarkFromNfa_Linear(10, b) }
func BenchmarkFromNfa_Linear_100(b *testing.B)  { benchmarkFromNfa_Linear(100, b) }
func BenchmarkFromNfa_Linear_1000(b *testing.B) { benchmarkFromNfa_Linear(1_000, b) }

// Benchmark: FanOut NFa to Dfa.
func BenchmarkFromNfa_FanOut_10(b *testing.B)   { benchmarkFromNfa_FanOut(10, b) }
func BenchmarkFromNfa_FanOut_100(b *testing.B)  { benchmarkFromNfa_FanOut(100, b) }
func BenchmarkFromNfa_FanOut_1000(b *testing.B) { benchmarkFromNfa_FanOut(1_000, b) }

func benchmarkFromNfa_Linear(count int, b *testing.B) {
	nMachine := nfa.New[int, int]()
	current := nMachine.Start()
	for i := 0; i < count; i++ {
		current = nMachine.Add(current, i)
	}

	b.ResetTimer()
	for b.Loop() {
		benchmarkOutput = dfa.FromNfa(nMachine)
	}
}

func benchmarkFromNfa_FanOut(count int, b *testing.B) {
	nMachine := nfa.New[int, int]()
	nStart := nMachine.Start()

	for i := 0; i < count; i++ {
		// All transitions on same symbol '1'
		nMachine.Add(nStart, 1)
	}

	b.ResetTimer()
	for b.Loop() {
		benchmarkOutput = dfa.FromNfa(nMachine)
	}
}
