package school

import (
	"sort"
)

// Grade is a struct with a level and a slice of students.
type Grade struct {
	level    int
	students []string
}

// School is a map with grade as key and slice of students as value.
type School map[int][]string

// New returns an empty school.
func New() *School {
	return &School{}
}

// Add adds a student to appropriate grade.
func (s *School) Add(name string, level int) {
	if stus, prs := (*s)[level]; prs {
		stus = append(stus, name)
		(*s)[level] = stus
	} else {
		(*s)[level] = []string{name}
	}
}

// Grade returns all students enrolled.
func (s *School) Grade(level int) []string {
	if _, prs := (*s)[level]; prs {
		return (*s)[level]
	}
	return []string{}
}

// Enrollment returns a sorted list of all students in all grades.
func (s *School) Enrollment() []Grade {
	ret := []Grade{}

	levels := []int{}
	for k := range *s {
		levels = append(levels, k)
	}
	sort.Ints(levels)

	for _, l := range levels {
		stus := (*s)[l]
		sort.Strings(stus)
		ret = append(ret, Grade{l, stus})
	}

	return ret
}
