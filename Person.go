package main

type IPerson interface {
	ShowClothes() string
}

type Person struct {
	name  string
	cloth string
}

func (p *Person) ShowClothes() string {
	return p.name
}

type AddCloth func() string
