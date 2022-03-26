package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	// Constructores
	// 1
	e := Employee{}
	fmt.Printf("%v\n", e)

	// 2
	e2 := Employee{
		id:       2,
		name:     "John",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)

	// 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 3
	e3.name = "Name"
	fmt.Printf("%v\n", *e3)

	// 4
	e4 := NewEmployee(4, "Name", true)
	fmt.Printf("%v\n", *e4)

	//e.id = 1
	//e.name = "Carlos"
	//fmt.Printf("%v", e)
	//e.SetId(5)
	//e.SetName("Andres")
	//fmt.Printf("%v", e)
	//fmt.Println(e.GetId())
	//fmt.Println(e.GetName())
}
