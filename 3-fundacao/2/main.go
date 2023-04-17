package main

import "fmt"

const a = "Hello, World!"

// Declaração de Escopo Global
// Inicialização das variáveis
var (
	b bool    = true
	c int     = 10
	nome string  = "Palin"
	e float64 = 1.2
)

func main() {
	a := "X" // string
	println(a)
	println(nome)
	fmt.Printf("A variável b é do tipo %T e tem o valor %v", b, b)
}

func x() {
}
