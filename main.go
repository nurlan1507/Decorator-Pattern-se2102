package main

import "fmt"

func main() {
	person := &Person{
		name: "Nurlan",
	}

	personWithCloth := &Cloth{
		person: person,
	}
	personWithCloth.ShowClothes()
	personWithClothAndJacket := Jacket{
		person: personWithCloth,
	}
	fmt.Println(personWithClothAndJacket.ShowClothes())
}
