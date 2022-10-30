package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAge() int {
	return s.age
}

func (s Student) getAverageGrade() float32 {
	sum := 0

	for _, value := range s.grades {
		sum += value
	}

	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() int {
	max := 0

	for _, grade := range s.grades {
		if grade > max {
			max = grade
		}
	}
	return max
}

func main() {
	s1 := Student{"Barun", []int{70, 90, 80, 85}, 19}
	s2 := Student{"Partho", []int{70, 93, 80, 85, 70}, 20}

	fmt.Println(s1)

	s1.setAge(7)
	fmt.Println(s1)

	average1 := s1.getAverageGrade()
	average2 := s2.getAverageGrade()
	fmt.Println(average1)
	fmt.Println(average2)

	max1 := s1.getMaxGrade()
	max2 := s2.getMaxGrade()
	fmt.Println(max1)
	fmt.Println(max2)
}
