package main

func main() {

	//  name <-> grade
	studentGrades := make(map[string]int)

	studentGrades["Alice"] = 85
	studentGrades["Bob"] = 90
	studentGrades["Charlie"] = 78

	delete(studentGrades, "Bob")

	grade, exists := studentGrades["Bob"]
	if exists {
		println("Bob's grade:", grade)
	} else {
		println("Bob's grade not found")
	}

	for name, grade := range studentGrades {
		println(name, ":", grade)
	}

}
