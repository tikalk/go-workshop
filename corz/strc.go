// Explains the way go implements inheritance
//
// Author: Dmitri Krasnenko

package corz

type Person struct {
	id int64
	name string
}

func NewPerson(id int64, name string) *Person {
	return &Person{id, name}
}

func (p *Person) Id() int64{
	return p.id
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

type Employee struct {
	// Compiler implicitly adds filed 'Person' of type *Person to the struct.
	*Person
	position string
}

func NewEmployee(id int64, name, position string) *Employee {
	return &Employee{ NewPerson(id, name),
		position,
	}
}

func (e *Employee) Position() string {
	return e.position
}
