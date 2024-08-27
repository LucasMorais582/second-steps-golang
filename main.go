// Conceitos básicos da linguagem
package main

import (
	"fmt"
	"curso-go/basicMath"
)

type Client struct {
	name   string
	email  string
	active bool
	Adress
}

type Adress struct {
	street string
	number string
}

// constraints
// type Number interface {
// 	~int | ~float64
// }

func main() {

	// ponteiros
	a := 10
	ponteiro := &a

	// mudando valor de a
	*ponteiro = 20

	// ponteiro2 recebe endereço de a
	ponteiro2 := ponteiro

	// mudando valor de a
	*ponteiro2 = 30

	println("Valor da variável a:", a)

	//closures
	total := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7, 8, 9) * 2
	}()
	fmt.Println(total)

	// closures 2
	// Definindo uma função que retorna uma closure
	increment := createIncrementer()

	// Usando a closure 2
	fmt.Println(increment()) // Saída: 1
	fmt.Println(increment()) // Saída: 2
	fmt.Println(increment()) // Saída: 3

	// struct 2
	client1 := Client{
		name:   "Lucas",
		email:  "email",
		active: true,
	}

	fmt.Println(client1)
	client1.Deactivate()
	fmt.Println(client1)

	// generics
	m := map[string]int {"Lucas": 1000, "Malu": 2000, "Theo": 3000}
	m2 := map[string]float64 {"Lucas": 1000.1, "Malu": 2000.3, "Theo": 3000.6}
	
	fmt.Printf("Soma inteiro: %d\n", sumAll(m))
	fmt.Printf("Soma float: %f\n", sumAll(m2))

	// go mod
	sum := basicMath.Sum(10, 20)
	fmt.Printf("Result:", sum)
}

// struct 2
func (client Client) Deactivate() {
	client.active = false
	fmt.Printf("%s was deactivate\n\n", client.name)
}


// funções variádicas
func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

// closures 2
func createIncrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// generics
func SumAll[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}
