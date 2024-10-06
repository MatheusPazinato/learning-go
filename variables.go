package main

import (
	"fmt"
)

type Person struct {
	name      string
	age       uint8
	document  string
	addresses []Address
}

type Address struct {
	street string
	number uint8
}

// declarando variáveis globais
var (
	age        uint8  = 24
	profession string = "Developer"
)

func main() {
	/* outra forma de criar uma variável, de uma maneira curta
	é utilizando :=, ele permite que você crie uma variável sem
	especificar o tipo, o Go implicitamente define o tipo, baseado
	no valor atribuído.
	*/
	name := "Matheus Pazinato"
	fmt.Println(name, age, profession)

	// structs
	address := Address{
		street: "Rua São Paulo",
		number: 23,
	}

	var addresses []Address
	addresses = append(addresses, address)

	fmt.Println(addresses)

	person := Person{
		name:      "Matheus Pazinato",
		age:       24,
		document:  "123.456.789-00",
		addresses: addresses,
	}

	fmt.Println(person)
}
