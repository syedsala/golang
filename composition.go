package main
// Composition: Mimic inheritance behavior using composition
import "fmt"

type Animal struct {
	Organs int
}

func (a *Animal) Eat() {
	fmt.Println("Eat functionality ", a.Organs)
}

func (a *Animal) Sleep() {
	fmt.Println("Sleep functionality", a.Organs)
}

type Dog struct {
	Animal
	Breed string
}

func main() {
	puppy := Dog { Animal {3 }, "sheepdog"}
	fmt.Println(puppy)
	puppy.Eat()
	puppy.Sleep()
}