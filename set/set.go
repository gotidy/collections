package set

// Set.
type Set[M comparable] map[M]struct{}

// New creates a new Set of members.
func New[M comparable](members ...M) Set[M] {
	result := make(Set[M])
	for _, member := range members {
		result[member] = struct{}{}
	}
	return result
}

// NewFromMapKeys creates a new Set of keys of the given map.
func NewFromMapKeys[M comparable, V any](m map[M]V) Set[M] {
	result := make(Set[M])
	for key := range m {
		result[key] = struct{}{}
	}
	return result
}

// NewFromSliceFunc creates a new Set of keys of the given map.
func NewFromSliceFunc[M comparable, V any](s []V, f func(v V) M) Set[M] {
	result := make(Set[M])
	for _, v := range s {
		result[f(v)] = struct{}{}
	}
	return result
}

// Len of set.
func (s Set[M]) Len() int { return len(s) }

// Empty checks that the set is empty.
func (s Set[M]) Empty() bool { return len(s) == 0 }

// Each members.
func (s Set[M]) Each(f func(m M)) {
	for m := range s {
		f(m)
	}
}

// Add members to set.
func (s Set[M]) Add(members ...M) Set[M] {
	for _, member := range members {
		s[member] = struct{}{}
	}
	return s
}

// Delete members from set.
func (s Set[M]) Delete(members ...M) Set[M] {
	for _, member := range members {
		delete(s, member)
	}
	return s
}

// Diff removes members from set.
func (s Set[M]) Diff(set Set[M]) Set[M] {
	for member := range set {
		delete(s, member)
	}
	return s
}

// Union members of sets.
func (s Set[M]) Union(set Set[M]) Set[M] {
	for member := range set {
		s[member] = struct{}{}
	}
	return s
}

// Intersect members of sets.
func (s Set[M]) Intersect(set Set[M]) Set[M] {
	for member := range s {
		if !set.Has(member) {
			delete(s, member)
		}
	}
	return s
}

// SymmetricDiff gets the symmetric difference of two sets and gives a set of elements, which are in either of the sets and not in their intersection.
func (s Set[M]) SymmetricDiff(set Set[M]) Set[M] {
	for member := range set {
		if s.Has(member) {
			delete(s, member)
		} else {
			s[member] = struct{}{}
		}
	}
	return s
}

// Clone set.
func (s Set[M]) Clone() Set[M] {
	result := make(Set[M], len(s))
	for member := range s {
		result[member] = struct{}{}
	}
	return result
}

// Has members of sets.
func (s Set[M]) Has(member M) bool {
	_, exists := s[member]
	return exists
}

// Diff returns s1 - s2.
func Diff[M comparable](s1, s2 Set[M]) Set[M] {
	result := New[M]()
	for member := range s1 {
		if !s2.Has(member) {
			result[member] = struct{}{}
		}
	}
	return result
}

// Union returns s1 + s2.
func Union[M comparable](s1, s2 Set[M]) Set[M] {
	result := New[M]()
	for member := range s1 {
		result[member] = struct{}{}
	}
	for member := range s2 {
		result[member] = struct{}{}
	}
	return result
}

// Intersect returns s1 members that is contained in s2.
func Intersect[M comparable](s1, s2 Set[M]) Set[M] {
	result := New[M]()
	for member := range s1 {
		if s2.Has(member) {
			result[member] = struct{}{}
		}
	}
	return result
}

// SymmetricDiff gets the symmetric difference of two sets and gives a set of elements, which are in either of the sets and not in their intersection.
func SymmetricDiff[M comparable](s1, s2 Set[M]) Set[M] {
	return s1.Clone().SymmetricDiff(s2)
}
