// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "strings"
    "testing"
)

func doSolve(t *testing.T, clues *Clues) {
    solver := NewSolver()
    board, ok := solver.SolveClues(clues)
    if expected := strings.TrimSpace(clues.Solution()); expected == "" {
        if ok {
            t.Errorf("Found false solution:\n%v\n", board)
        } else {
            t.Log("Correctly deduced no solution\n")
        }
    } else {
        expected = strings.TrimSpace(expected)
        if !ok {
            t.Errorf("Didn't find solution:\n%v\n", expected)
        } else if actual := strings.TrimSpace(board.String()); expected != actual {
            t.Errorf("Expected solution:\n%v\nActualSolution:\n%v\n", expected, actual)
        } else {
            t.Logf("Correctly deduced solution:\n%v\n", expected)
        }
    }
}

func TestSolver_Solve_SampleClues_0(t *testing.T) {
    doSolve(t, SampleClues()[0])
}

func TestSolver_Solve_SampleClues_1(t *testing.T) {
    doSolve(t, SampleClues()[1])
}

func TestSolver_Solve_SampleClues_2(t *testing.T) {
    doSolve(t, SampleClues()[2])
}

func TestSolver_Solve_SampleClues_3(t *testing.T) {
    doSolve(t, SampleClues()[3])
}

func TestSolver_Solve_SampleClues_4(t *testing.T) {
    doSolve(t, SampleClues()[4])
}

func TestSolver_Solve_SampleClues_5(t *testing.T) {
    doSolve(t, SampleClues()[5])
}

func TestSolver_Solve_SampleClues_6(t *testing.T) {
    doSolve(t, SampleClues()[6])
}

func TestSolver_Solve_SampleClues_7(t *testing.T) {
    doSolve(t, SampleClues()[7])
}

func TestSolver_Solve_SampleClues_8(t *testing.T) {
    doSolve(t, SampleClues()[8])
}

func TestSolver_Solve_SampleClues_9(t *testing.T) {
    doSolve(t, SampleClues()[9])
}

func TestSolver_Solve_SampleClues_10(t *testing.T) {
    doSolve(t, SampleClues()[10])
}

func TestSolver_Solve_SampleClues_11(t *testing.T) {
    doSolve(t, SampleClues()[11])
}

func solverBenchmark(b *testing.B, index int) {
    solver := NewSolver()
    clues := SampleClues()[index]
    for i := 0; i < b.N; i++ {
        solver.SolveClues(clues)
    }
}

func BenchmarkSolver_Solve_SampleClues_0(b *testing.B) {
    solverBenchmark(b, 0)
}

func BenchmarkSolver_Solve_SampleClues_1(b *testing.B) {
    solverBenchmark(b, 1)
}

func BenchmarkSolver_Solve_SampleClues_2(b *testing.B) {
    solverBenchmark(b, 2)
}

func BenchmarkSolver_Solve_SampleClues_3(b *testing.B) {
    solverBenchmark(b, 3)
}

func BenchmarkSolver_Solve_SampleClues_4(b *testing.B) {
    solverBenchmark(b, 4)
}

func BenchmarkSolver_Solve_SampleClues_5(b *testing.B) {
    solverBenchmark(b, 5)
}

func BenchmarkSolver_Solve_SampleClues_6(b *testing.B) {
    solverBenchmark(b, 6)
}

func BenchmarkSolver_Solve_SampleClues_7(b *testing.B) {
    solverBenchmark(b, 7)
}

func BenchmarkSolver_Solve_SampleClues_8(b *testing.B) {
    solverBenchmark(b, 8)
}

func BenchmarkSolver_Solve_SampleClues_9(b *testing.B) {
    solverBenchmark(b, 9)
}

func BenchmarkSolver_Solve_SampleClues_10(b *testing.B) {
    solverBenchmark(b, 10)
}

func BenchmarkSolver_Solve_SampleClues_11(b *testing.B) {
    solverBenchmark(b, 11)
}

func BenchmarkSolver_Solve_SampleClues_All(b *testing.B) {
    solver := NewSolver()
    for i := 0; i < b.N; i++ {
        for _, clues := range SampleClues() {
            solver.SolveClues(clues)
        }
    }
}
