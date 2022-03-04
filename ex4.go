package main

import "fmt"

const (
	_ = 1990 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(b, c, d, e)
}

//Utilizando IOTA, crie 4 constantes cujos valores sejam os proximos 4 anos. Demonstre valores.
