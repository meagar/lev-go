package lev

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"testing"
)

type TestCase struct {
	a, b  string
	want  int // expected Levenshtein result
	wantD int // expected Damerau-Levenshtein result
}

var testCases = []TestCase{
	{"foo", "", 3, 3},       // three deletions
	{"", "asdf", 4, 4},      // four additions
	{"foo", "bar", 3, 3},    // three replacements
	{"foo", "foobar", 3, 3}, // three additions ("bar")
	{"foobar", "bar", 3, 3}, // three deletions ("foo")
	{"foo", "foo", 0, 0},    // identical
	{"fast", "past", 1, 1},  // one replacement
	{"boot", "tube", 4, 4},  // four replacements
	{"cabbages", "rabbit", 5, 5},
	{"foot", "poof", 2, 2},
	{"kitten", "sitting", 3, 3},
	{"alphabet", "fgsdgszxc", 9, 9},
	{"aaaaaaaaaa", "bbbbbbbbbb", 10, 10},
	{"ab", "ba", 2, 1},
	{"abab", "baba", 2, 2},
	{"abcd", "badc", 3, 2},
	{"blah", "flha", 3, 2},
}

// Tests a Levenshtein distance function
func testFn(t *testing.T, fn func(a, b string) int) {
	t.Helper()
	fnName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s(%q, %q)", fnName, tc.a, tc.b), func(t *testing.T) {
			got := fn(tc.a, tc.b)

			if got != tc.want {
				t.Errorf("%s(%q, %q): Got %d, want %d", fnName, tc.a, tc.b, got, tc.want)
				t.FailNow()
			}
		})
	}
}

// Tests a Damerau-Levenshtein distance function
func testFnD(t *testing.T, fn func(a, b string) int) {
	t.Helper()
	fnName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s(%q, %q)", fnName, tc.a, tc.b), func(t *testing.T) {
			got := fn(tc.a, tc.b)

			if got != tc.wantD {
				t.Errorf("%s(%q, %q): Got %d, want %d", fnName, tc.a, tc.b, got, tc.wantD)
				t.FailNow()
			}
		})
	}
}
func TestDistance(t *testing.T) {
	testFn(t, Distance)
}
func TestNaiveDistance(t *testing.T) {
	testFn(t, naiveDistance)
}

func TestMatrixDistance(t *testing.T) {
	testFn(t, matrixDistance)
}

func TestMatrixDistanceD(t *testing.T) {
	testFnD(t, matrixDistanceD)
}

func TestDoubleRowDistance(t *testing.T) {
	testFn(t, doubleRowDistance)
}

func TestSingleRowDistance(t *testing.T) {
	testFn(t, singleRowDistance)
}

func BenchmarkNaiveDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		naiveDistance("aaaaaaaaaa", "bbbbbbbbbb")
	}
}

func BenchmarkMatrixDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		matrixDistance("aaaaaaaaaa", "bbbbbbbbbb")
	}
}

func BenchmarkDoubleRowDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		doubleRowDistance("aaaaaaaaaa", "bbbbbbbbbb")
	}
}

func BenchmarkSingleRowDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		singleRowDistance("aaaaaaaaaa", "bbbbbbbbbb")
	}
}

func TestMin(t *testing.T) {
	for n := 0; n < 10000; n++ {
		a := rand.Int()
		b := rand.Int()
		c := rand.Int()

		arr := []int{a, b, c}
		sort.Ints(arr)
		want := arr[0]
		got := min3(a, b, c)

		if got != want {
			t.Errorf("min(%d, %d, %d): Got %d, want %d", a, b, c, got, want)
		}
	}
}
