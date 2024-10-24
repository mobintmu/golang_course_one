package model

type Person struct {
	name   string
	family string
}

func New(name, family string) Person {
	return Person{
		name:   name,
		family: family,
	}
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetFamily(family string) {
	p.family = family
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetFamily() string {
	return p.family
}
