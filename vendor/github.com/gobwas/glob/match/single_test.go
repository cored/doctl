package match

import (
	"reflect"
	"testing"
)

func TestSingleIndex(t *testing.T) {
	for id, test := range []struct {
		separators string
		fixture    string
		index      int
		segments   []int
	}{
		{
			".",
			".abc",
			1,
			[]int{1},
		},
		{
			".",
			".",
			-1,
			nil,
		},
	} {
		p := Single{test.separators}
		index, segments := p.Index(test.fixture)
		if index != test.index {
			t.Errorf("#%d unexpected index: exp: %d, act: %d", id, test.index, index)
		}
		if !reflect.DeepEqual(segments, test.segments) {
			t.Errorf("#%d unexpected segments: exp: %v, act: %v", id, test.segments, segments)
		}
	}
}

func BenchmarkIndexSingle(b *testing.B) {
	m := Single{bench_separators}
	for i := 0; i < b.N; i++ {
		m.Index(bench_pattern)
	}
}
