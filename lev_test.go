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
	a, b string
	want int
}

var testCases = []TestCase{
	{"foo", "", 3},
	{"", "asdf", 4},
	{"foo", "bar", 3},
	{"foo", "foobar", 3},
	{"foobar", "bar", 3},
	{"foo", "foo", 0},
	{"fast", "past", 1},
	{"boot", "tube", 4},
	{"cabbages", "rabbit", 5},
	{"foot", "poof", 2},
	{"kitten", "sitting", 3},
	{"alphabet", "fgsdgszxc", 9},
	{"aaaaaaaaaa", "bbbbbbbbbb", 10},
}

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

func TestDistance(t *testing.T) {
	testFn(t, Distance)
}
func TestNaiveDistance(t *testing.T) {
	testFn(t, naiveDistance)
}

func TestMatrixDistance(t *testing.T) {
	testFn(t, matrixDistance)
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
		got := min(a, b, c)

		if got != want {
			t.Errorf("min(%d, %d, %d): Got %d, want %d", a, b, c, got, want)
		}
	}
}
