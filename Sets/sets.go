package main

import "fmt"

/* A Set is a linear data structure that has collection of unique values*/

type Set struct {
	integerMap map[int]bool
}

func (s *Set) New() {
	s.integerMap = make(map[int]bool)
}

func (s *Set) ContainsElement(element int) bool {
	_, ok := s.integerMap[element]
	if ok {
		return true
	} else {
		return false
	}
}

func (s *Set) DeleteElement(element int) {
	delete(s.integerMap, element)
}

func (s *Set) AddElement(element int) {
	if !s.ContainsElement(element) {
		s.integerMap[element] = true
	}
}

func main() {
	set := &Set{integerMap: map[int]bool{
		1: true,
	},
	}

	set.AddElement(2)
	set.AddElement(4)
	fmt.Println(set)
	fmt.Println(set.ContainsElement(7))

}
