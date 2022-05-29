package slice

import (
	"reflect"
	"strconv"
	"testing"
)

func TestIndexExists(t *testing.T) {
	s := []string{"0", "1", "2", "3", "4", "5"}
	expected := 2
	if i := Index(s, "2"); i != expected {
		t.Errorf("expected %d, actual %d", expected, i)
	}
}

func TestIndexNotExists(t *testing.T) {
	s := []string{"0", "1", "2", "3", "4", "5"}
	if i := Index(s, "10"); i != -1 {
		t.Errorf("expected -1, actual %d", i)
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name     string
		s        []string
		pos      int
		v        string
		expected []string
	}{
		{
			name:     "Insert at start with reallocate",
			s:        []string{"0", "1", "2", "3", "4", "5"},
			pos:      0,
			v:        "10",
			expected: []string{"10", "0", "1", "2", "3", "4", "5"},
		},
		{
			name:     "Insert at end with reallocate",
			s:        []string{"0", "1", "2", "3", "4", "5"},
			pos:      6,
			v:        "10",
			expected: []string{"0", "1", "2", "3", "4", "5", "10"},
		},
		{
			name:     "Insert in middle with reallocate",
			s:        []string{"0", "1", "2", "3", "4", "5"},
			pos:      4,
			v:        "10",
			expected: []string{"0", "1", "2", "3", "10", "4", "5"},
		},
		{
			name:     "Insert at start with reallocate",
			s:        []string{"0", "1", "2", "3", "4", "5", ""}[:6],
			pos:      0,
			v:        "10",
			expected: []string{"10", "0", "1", "2", "3", "4", "5"},
		},
		{
			name:     "Insert at end with reallocate",
			s:        []string{"0", "1", "2", "3", "4", "5", ""}[:6],
			pos:      6,
			v:        "10",
			expected: []string{"0", "1", "2", "3", "4", "5", "10"},
		},
		{
			name:     "Insert in middle with reallocate",
			s:        []string{"0", "1", "2", "3", "4", "5", ""}[:6],
			pos:      4,
			v:        "10",
			expected: []string{"0", "1", "2", "3", "10", "4", "5"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := Insert(test.s, test.pos, test.v); !reflect.DeepEqual(res, test.expected) {
				t.Errorf("expected %+v, actual %+v", test.expected, res)
			}
		})
	}
}

func TestInsertPanic(t *testing.T) {
	defer func() { _ = recover() }()
	_ = Insert([]string{"0", "1", "2", "3", "4", "5"}, 7, "10")
	t.Errorf("expected panic")
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		s        []string
		expected []string
	}{
		{
			name:     "even",
			s:        []string{"0", "1", "2", "3", "4", "5"},
			expected: []string{"5", "4", "3", "2", "1", "0"},
		},
		{
			name:     "odd",
			s:        []string{"0", "1", "2", "3", "4", "5", "6"},
			expected: []string{"6", "5", "4", "3", "2", "1", "0"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Reverse(test.s)
			if !reflect.DeepEqual(test.s, test.expected) {
				t.Errorf("expected %+v, actual %+v", test.expected, test.s)
			}
		})
	}
}

func TestMap(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	expected := []string{"0", "1", "2", "3", "4", "5"}

	res := Map(s, func(i int) string {
		return strconv.Itoa(i)
	})
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %+v, actual %+v", expected, res)
	}
}

func TestFilter(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	expected := []int{2, 3}

	res := Filter(s, func(i int) bool {
		return i == 2 || i == 3
	})
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %+v, actual %+v", expected, res)
	}
}

func TestReduce(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	expected := 20

	res := Reduce(s, 5, func(r, i int) int {
		return r + i
	})
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %+v, actual %+v", expected, res)
	}
}
