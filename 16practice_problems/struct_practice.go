package main

import "fmt"

type School struct {
	School_Name string
	Address     string
}

type Department struct {
	Dept_Name string
	*School
}

type Teacher struct {
	Name   string
	Gender string
	Age    int
	*Department
	*School
}

type Student struct {
	Name     string
	Gender   string
	Standard int
	Roll     int
	*Department
	*School
}

type Result struct {
	CGPA  float32
	Marks []int
	*Student
	*Department
	*School
}

func (r Result) getAverageResult(st Student) float32 {

}

func main() {
	School1 := School{"Zilla_School", "Faridpur"}
	Dept1 := Department{"Science", &School1}

	Student1 := Student{"Sifat", "Male", 10, 1, &Dept1, &School1}
	Result1 := Result{3.78, []int{80, 70, 85, 95}, &Student1, &Dept1, &School1}
	Teacher1 := Teacher{"Animesh", "Male", 50, &Dept1, &School1}

	fmt.Println(Student1.School_Name)
	fmt.Println(Teacher1.Dept_Name)

	// get average result of student sifat
	Result1.getAverageResult(Student1)
}
