package main

type Cloth struct {
	person IPerson
}

func (c *Cloth) ShowClothes() string {
	return c.person.ShowClothes() + " + " + "cloth"
}

type Jacket struct {
	person IPerson
}

func (c *Jacket) ShowClothes() string {
	return c.person.ShowClothes() + " + " + "Jacket"
}
